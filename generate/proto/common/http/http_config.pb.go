// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: http_config.proto

package http

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_http_config_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "http.RouteGroup",
		Tag:           "bytes,50000,opt,name=RouteGroup",
		Filename:      "http_config.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "http.Get",
		Tag:           "bytes,50000,opt,name=Get",
		Filename:      "http_config.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "http.Head",
		Tag:           "bytes,50001,opt,name=Head",
		Filename:      "http_config.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50002,
		Name:          "http.Post",
		Tag:           "bytes,50002,opt,name=Post",
		Filename:      "http_config.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50003,
		Name:          "http.Put",
		Tag:           "bytes,50003,opt,name=Put",
		Filename:      "http_config.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50004,
		Name:          "http.Delete",
		Tag:           "bytes,50004,opt,name=Delete",
		Filename:      "http_config.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50005,
		Name:          "http.Options",
		Tag:           "bytes,50005,opt,name=Options",
		Filename:      "http_config.proto",
	},
}

// Extension fields to descriptor.ServiceOptions.
var (
	// 路由分组名字
	//
	// optional string RouteGroup = 50000;
	E_RouteGroup = &file_http_config_proto_extTypes[0]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional string Get = 50000;
	E_Get = &file_http_config_proto_extTypes[1]
	// optional string Head = 50001;
	E_Head = &file_http_config_proto_extTypes[2]
	// optional string Post = 50002;
	E_Post = &file_http_config_proto_extTypes[3]
	// optional string Put = 50003;
	E_Put = &file_http_config_proto_extTypes[4]
	// optional string Delete = 50004;
	E_Delete = &file_http_config_proto_extTypes[5]
	// optional string Options = 50005;
	E_Options = &file_http_config_proto_extTypes[6]
)

var File_http_config_proto protoreflect.FileDescriptor

var file_http_config_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x41, 0x0a, 0x0a, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x32,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47,
	0x65, 0x74, 0x3a, 0x34, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x48, 0x65, 0x61, 0x64, 0x3a, 0x34, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x3a, 0x32,
	0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50,
	0x75, 0x74, 0x3a, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd4, 0x86, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x3a, 0x0a, 0x07,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd5, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x74, 0x66, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_http_config_proto_goTypes = []interface{}{
	(*descriptor.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
	(*descriptor.MethodOptions)(nil),  // 1: google.protobuf.MethodOptions
}
var file_http_config_proto_depIdxs = []int32{
	0, // 0: http.RouteGroup:extendee -> google.protobuf.ServiceOptions
	1, // 1: http.Get:extendee -> google.protobuf.MethodOptions
	1, // 2: http.Head:extendee -> google.protobuf.MethodOptions
	1, // 3: http.Post:extendee -> google.protobuf.MethodOptions
	1, // 4: http.Put:extendee -> google.protobuf.MethodOptions
	1, // 5: http.Delete:extendee -> google.protobuf.MethodOptions
	1, // 6: http.Options:extendee -> google.protobuf.MethodOptions
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	0, // [0:7] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_http_config_proto_init() }
func file_http_config_proto_init() {
	if File_http_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 7,
			NumServices:   0,
		},
		GoTypes:           file_http_config_proto_goTypes,
		DependencyIndexes: file_http_config_proto_depIdxs,
		ExtensionInfos:    file_http_config_proto_extTypes,
	}.Build()
	File_http_config_proto = out.File
	file_http_config_proto_rawDesc = nil
	file_http_config_proto_goTypes = nil
	file_http_config_proto_depIdxs = nil
}
