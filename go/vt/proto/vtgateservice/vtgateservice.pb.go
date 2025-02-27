//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Service definition for vtgateservice.
// This is the main entry point to Vitess.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: vtgateservice.proto

package vtgateservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	vtgate "vitess.io/vitess/go/vt/proto/vtgate"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_vtgateservice_proto protoreflect.FileDescriptor

var file_vtgateservice_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x8f, 0x04, 0x0a, 0x06, 0x56, 0x69, 0x74, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a,
	0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x76, 0x74,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x76, 0x74, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5d, 0x0a, 0x12, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x56, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76,
	0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x07, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76,
	0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x0a, 0x14, 0x69, 0x6f, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x2a, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x76, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_vtgateservice_proto_goTypes = []interface{}{
	(*vtgate.ExecuteRequest)(nil),             // 0: vtgate.ExecuteRequest
	(*vtgate.ExecuteBatchRequest)(nil),        // 1: vtgate.ExecuteBatchRequest
	(*vtgate.StreamExecuteRequest)(nil),       // 2: vtgate.StreamExecuteRequest
	(*vtgate.ResolveTransactionRequest)(nil),  // 3: vtgate.ResolveTransactionRequest
	(*vtgate.VStreamRequest)(nil),             // 4: vtgate.VStreamRequest
	(*vtgate.PrepareRequest)(nil),             // 5: vtgate.PrepareRequest
	(*vtgate.CloseSessionRequest)(nil),        // 6: vtgate.CloseSessionRequest
	(*vtgate.ExecuteResponse)(nil),            // 7: vtgate.ExecuteResponse
	(*vtgate.ExecuteBatchResponse)(nil),       // 8: vtgate.ExecuteBatchResponse
	(*vtgate.StreamExecuteResponse)(nil),      // 9: vtgate.StreamExecuteResponse
	(*vtgate.ResolveTransactionResponse)(nil), // 10: vtgate.ResolveTransactionResponse
	(*vtgate.VStreamResponse)(nil),            // 11: vtgate.VStreamResponse
	(*vtgate.PrepareResponse)(nil),            // 12: vtgate.PrepareResponse
	(*vtgate.CloseSessionResponse)(nil),       // 13: vtgate.CloseSessionResponse
}
var file_vtgateservice_proto_depIdxs = []int32{
	0,  // 0: vtgateservice.Vitess.Execute:input_type -> vtgate.ExecuteRequest
	1,  // 1: vtgateservice.Vitess.ExecuteBatch:input_type -> vtgate.ExecuteBatchRequest
	2,  // 2: vtgateservice.Vitess.StreamExecute:input_type -> vtgate.StreamExecuteRequest
	3,  // 3: vtgateservice.Vitess.ResolveTransaction:input_type -> vtgate.ResolveTransactionRequest
	4,  // 4: vtgateservice.Vitess.VStream:input_type -> vtgate.VStreamRequest
	5,  // 5: vtgateservice.Vitess.Prepare:input_type -> vtgate.PrepareRequest
	6,  // 6: vtgateservice.Vitess.CloseSession:input_type -> vtgate.CloseSessionRequest
	7,  // 7: vtgateservice.Vitess.Execute:output_type -> vtgate.ExecuteResponse
	8,  // 8: vtgateservice.Vitess.ExecuteBatch:output_type -> vtgate.ExecuteBatchResponse
	9,  // 9: vtgateservice.Vitess.StreamExecute:output_type -> vtgate.StreamExecuteResponse
	10, // 10: vtgateservice.Vitess.ResolveTransaction:output_type -> vtgate.ResolveTransactionResponse
	11, // 11: vtgateservice.Vitess.VStream:output_type -> vtgate.VStreamResponse
	12, // 12: vtgateservice.Vitess.Prepare:output_type -> vtgate.PrepareResponse
	13, // 13: vtgateservice.Vitess.CloseSession:output_type -> vtgate.CloseSessionResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_vtgateservice_proto_init() }
func file_vtgateservice_proto_init() {
	if File_vtgateservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vtgateservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vtgateservice_proto_goTypes,
		DependencyIndexes: file_vtgateservice_proto_depIdxs,
	}.Build()
	File_vtgateservice_proto = out.File
	file_vtgateservice_proto_rawDesc = nil
	file_vtgateservice_proto_goTypes = nil
	file_vtgateservice_proto_depIdxs = nil
}
