// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: server/services/liefer.proto

package services

import (
	types "github.com/lxgr-linux/liefer/server/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_server_services_liefer_proto protoreflect.FileDescriptor

var file_server_services_liefer_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x6c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x54, 0x0a, 0x06, 0x4c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x07, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x00, 0x30, 0x01, 0x42, 0xd1, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x69,
	0x65, 0x66, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x42, 0x0b, 0x4c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x78, 0x67, 0x72, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x6c, 0x69, 0x65, 0x66, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x4c, 0x53, 0x53, 0xaa, 0x02, 0x16, 0x4c, 0x69, 0x65, 0x66, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xca, 0x02, 0x16, 0x4c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x22, 0x4c, 0x69, 0x65, 0x66,
	0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x18, 0x4c, 0x69, 0x65, 0x66, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var file_server_services_liefer_proto_goTypes = []any{
	(*types.Payload)(nil),  // 0: liefer.server.types.Payload
	(*types.Progress)(nil), // 1: liefer.server.types.Progress
}
var file_server_services_liefer_proto_depIdxs = []int32{
	0, // 0: liefer.server.services.Liefer.Deliver:input_type -> liefer.server.types.Payload
	1, // 1: liefer.server.services.Liefer.Deliver:output_type -> liefer.server.types.Progress
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_services_liefer_proto_init() }
func file_server_services_liefer_proto_init() {
	if File_server_services_liefer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_server_services_liefer_proto_rawDesc), len(file_server_services_liefer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_services_liefer_proto_goTypes,
		DependencyIndexes: file_server_services_liefer_proto_depIdxs,
	}.Build()
	File_server_services_liefer_proto = out.File
	file_server_services_liefer_proto_goTypes = nil
	file_server_services_liefer_proto_depIdxs = nil
}
