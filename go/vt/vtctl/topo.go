/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vtctl

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	"vitess.io/vitess/go/vt/topo"
	"vitess.io/vitess/go/vt/wrangler"

	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	vschemapb "vitess.io/vitess/go/vt/proto/vschema"
)

// This file contains the topo command group for vtctl.

const topoGroupName = "Topo"

func init() {
	addCommandGroup(topoGroupName)

	addCommand(topoGroupName, command{
		name:   "TopoCat",
		method: commandTopoCat,
		params: "[-cell <cell>] [-decode_proto] [-decode_proto_json] [-long] <path> [<path>...]",
		help:   "Retrieves the file(s) at <path> from the topo service, and displays it. It can resolve wildcards, and decode the proto-encoded data.",
	})

	addCommand(topoGroupName, command{
		name:   "TopoCp",
		method: commandTopoCp,
		params: "[-cell <cell>] [-to_topo] <src> <dst>",
		help:   "Copies a file from topo to local file structure, or the other way around",
	})
}

// DecodeContent uses the filename to imply a type, and proto-decodes
// the right object, then echoes it as a string.
func DecodeContent(filename string, data []byte, json bool) (string, error) {
	name := path.Base(filename)
	dir := path.Dir(filename)
	var p proto.Message
	switch name {
	case topo.CellInfoFile:
		p = new(topodatapb.CellInfo)
	case topo.KeyspaceFile:
		p = new(topodatapb.Keyspace)
	case topo.ShardFile:
		p = new(topodatapb.Shard)
	case topo.VSchemaFile:
		p = new(vschemapb.Keyspace)
	case topo.ShardReplicationFile:
		p = new(topodatapb.ShardReplication)
	case topo.TabletFile:
		p = new(topodatapb.Tablet)
	case topo.SrvVSchemaFile:
		p = new(vschemapb.SrvVSchema)
	case topo.SrvKeyspaceFile:
		p = new(topodatapb.SrvKeyspace)
	case topo.RoutingRulesFile:
		p = new(vschemapb.RoutingRules)
	default:
		switch dir {
		case "/" + topo.GetExternalVitessClusterDir():
			p = new(topodatapb.ExternalVitessCluster)
		default:
		}
		if p == nil {
			if json {
				return "", fmt.Errorf("unknown topo protobuf type for %v", name)
			}
			return string(data), nil
		}
	}

	if err := proto.Unmarshal(data, p); err != nil {
		return string(data), err
	}

	var marshalled []byte
	var err error
	if json {
		marshalled, err = protojson.Marshal(p)
	} else {
		marshalled, err = prototext.Marshal(p)
	}
	return string(marshalled), err
}

func commandTopoCat(ctx context.Context, wr *wrangler.Wrangler, subFlags *flag.FlagSet, args []string) error {
	cell := subFlags.String("cell", topo.GlobalCell, "topology cell to cat the file from. Defaults to global cell.")
	long := subFlags.Bool("long", false, "long listing.")
	decodeProtoJSON := subFlags.Bool("decode_proto_json", false, "decode proto files and display them as json")
	decodeProto := subFlags.Bool("decode_proto", false, "decode proto files and display them as text")
	subFlags.Parse(args)
	if subFlags.NArg() == 0 {
		return fmt.Errorf("TopoCat: no path specified")
	}
	resolved, err := wr.TopoServer().ResolveWildcards(ctx, *cell, subFlags.Args())
	if err != nil {
		return fmt.Errorf("TopoCat: invalid wildcards: %v", err)
	}
	if len(resolved) == 0 {
		// The wildcards didn't result in anything, we're done.
		return nil
	}

	conn, err := wr.TopoServer().ConnForCell(ctx, *cell)
	if err != nil {
		return err
	}

	var topologyDecoder TopologyDecoder
	switch {
	case *decodeProtoJSON:
		topologyDecoder = JSONTopologyDecoder{}
	case *decodeProto:
		topologyDecoder = ProtoTopologyDecoder{}
	default:
		topologyDecoder = PlainTopologyDecoder{}
	}

	return topologyDecoder.decode(ctx, resolved, conn, wr, *long)
}

func commandTopoCp(ctx context.Context, wr *wrangler.Wrangler, subFlags *flag.FlagSet, args []string) error {
	cell := subFlags.String("cell", topo.GlobalCell, "topology cell to use for the copy. Defaults to global cell.")
	toTopo := subFlags.Bool("to_topo", false, "copies from local server to topo instead (reverse direction).")
	subFlags.Parse(args)
	if subFlags.NArg() != 2 {
		return fmt.Errorf("TopoCp: need source and destination")
	}
	from := subFlags.Arg(0)
	to := subFlags.Arg(1)
	if *toTopo {
		return copyFileToTopo(ctx, wr.TopoServer(), *cell, from, to)
	}
	return copyFileFromTopo(ctx, wr.TopoServer(), *cell, from, to)
}

func copyFileFromTopo(ctx context.Context, ts *topo.Server, cell, from, to string) error {
	conn, err := ts.ConnForCell(ctx, cell)
	if err != nil {
		return err
	}
	data, _, err := conn.Get(ctx, from)
	if err != nil {
		return err
	}
	return os.WriteFile(to, data, 0644)
}

func copyFileToTopo(ctx context.Context, ts *topo.Server, cell, from, to string) error {
	conn, err := ts.ConnForCell(ctx, cell)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}
	_, err = conn.Update(ctx, to, data, nil)
	return err
}

// TopologyDecoder interface for exporting out a leaf node in a readable form
type TopologyDecoder interface {
	decode(context.Context, []string, topo.Conn, *wrangler.Wrangler, bool) error
}

// ProtoTopologyDecoder exports topo node as a proto
type ProtoTopologyDecoder struct{}

// PlainTopologyDecoder exports topo node as plain text
type PlainTopologyDecoder struct{}

// JSONTopologyDecoder exports topo node as JSON
type JSONTopologyDecoder struct{}

func (d ProtoTopologyDecoder) decode(ctx context.Context, topoPaths []string, conn topo.Conn, wr *wrangler.Wrangler, long bool) error {
	hasError := false
	for _, topoPath := range topoPaths {
		data, version, err := conn.Get(ctx, topoPath)
		if err != nil {
			hasError = true
			wr.Logger().Printf("TopoCat: Get(%v) failed: %v\n", topoPath, err)
			continue
		}

		if long {
			wr.Logger().Printf("path=%v version=%v\n", topoPath, version)
		}

		decoded, err := DecodeContent(topoPath, data, false)
		if err != nil {
			wr.Logger().Warningf("TopoCat: cannot proto decode %v: %v", topoPath, err)
			decoded = string(data)
		}

		wr.Logger().Printf(decoded)
		if len(decoded) > 0 && decoded[len(decoded)-1] != '\n' && long {
			wr.Logger().Printf("\n")
		}
	}

	if hasError {
		return fmt.Errorf("TopoCat: some paths had errors")
	}
	return nil
}

func (d PlainTopologyDecoder) decode(ctx context.Context, topoPaths []string, conn topo.Conn, wr *wrangler.Wrangler, long bool) error {
	hasError := false
	for _, topoPath := range topoPaths {
		data, version, err := conn.Get(ctx, topoPath)
		if err != nil {
			hasError = true
			wr.Logger().Printf("TopoCat: Get(%v) failed: %v\n", topoPath, err)
			continue
		}

		if long {
			wr.Logger().Printf("path=%v version=%v\n", topoPath, version)
		}
		decoded := string(data)
		wr.Logger().Printf(decoded)
		if len(decoded) > 0 && decoded[len(decoded)-1] != '\n' && long {
			wr.Logger().Printf("\n")
		}
	}

	if hasError {
		return fmt.Errorf("TopoCat: some paths had errors")
	}
	return nil
}

func (d JSONTopologyDecoder) decode(ctx context.Context, topoPaths []string, conn topo.Conn, wr *wrangler.Wrangler, long bool) error {
	hasError := false
	var jsonData []any
	for _, topoPath := range topoPaths {
		data, version, err := conn.Get(ctx, topoPath)
		if err != nil {
			hasError = true
			wr.Logger().Printf("TopoCat: Get(%v) failed: %v\n", topoPath, err)
			continue
		}

		decoded, err := DecodeContent(topoPath, data, true)
		if err != nil {
			hasError = true
			wr.Logger().Printf("TopoCat: cannot proto decode %v: %v", topoPath, err)
			continue
		}

		var jsonDatum map[string]any
		if err = json.Unmarshal([]byte(decoded), &jsonDatum); err != nil {
			hasError = true
			wr.Logger().Printf("TopoCat: cannot json Unmarshal %v: %v", topoPath, err)
			continue
		}

		if long {
			jsonDatum["__path"] = topoPath
			jsonDatum["__version"] = version.String()
		}
		jsonData = append(jsonData, jsonDatum)
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		hasError = true
		wr.Logger().Printf("TopoCat: cannot json Marshal: %v", err)
	} else {
		wr.Logger().Printf(string(jsonBytes) + "\n")
	}

	if hasError {
		return fmt.Errorf("TopoCat: some paths had errors")
	}
	return nil
}
