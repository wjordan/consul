// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/pbproxystate/protocol.proto

package pbproxystate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// +kubebuilder:validation:Enum=PROTOCOL_UNSPECIFIED;PROTOCOL_TCP;PROTOCOL_HTTP;PROTOCOL_HTTP2;PROTOCOL_GRPC
// +kubebuilder:validation:Type=string
type Protocol int32

const (
	Protocol_PROTOCOL_UNSPECIFIED Protocol = 0
	Protocol_PROTOCOL_TCP         Protocol = 1
	Protocol_PROTOCOL_HTTP        Protocol = 2
	Protocol_PROTOCOL_HTTP2       Protocol = 3
	Protocol_PROTOCOL_GRPC        Protocol = 4
	// Protocol Mesh indicates that this port can speak Consul's mTLS based mesh protocol.
	Protocol_PROTOCOL_MESH Protocol = 5
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "PROTOCOL_UNSPECIFIED",
		1: "PROTOCOL_TCP",
		2: "PROTOCOL_HTTP",
		3: "PROTOCOL_HTTP2",
		4: "PROTOCOL_GRPC",
		5: "PROTOCOL_MESH",
	}
	Protocol_value = map[string]int32{
		"PROTOCOL_UNSPECIFIED": 0,
		"PROTOCOL_TCP":         1,
		"PROTOCOL_HTTP":        2,
		"PROTOCOL_HTTP2":       3,
		"PROTOCOL_GRPC":        4,
		"PROTOCOL_MESH":        5,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_pbmesh_v2beta1_pbproxystate_protocol_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_pbmesh_v2beta1_pbproxystate_protocol_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescGZIP(), []int{0}
}

var File_pbmesh_v2beta1_pbproxystate_protocol_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x83, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f,
	0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x54, 0x43, 0x50, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x48, 0x54,
	0x54, 0x50, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c,
	0x5f, 0x48, 0x54, 0x54, 0x50, 0x32, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x54,
	0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x4d, 0x45, 0x53, 0x48, 0x10, 0x05, 0x42, 0xd3,
	0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x62, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0xa2, 0x02, 0x05, 0x48, 0x43, 0x4d, 0x56, 0x50,
	0xaa, 0x02, 0x2a, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0xca, 0x02, 0x2a,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x50, 0x62,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0xe2, 0x02, 0x36, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65,
	0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x50, 0x62, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a,
	0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3a, 0x3a, 0x50, 0x62, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescData = file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDesc
)

func file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDescData
}

var file_pbmesh_v2beta1_pbproxystate_protocol_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pbmesh_v2beta1_pbproxystate_protocol_proto_goTypes = []interface{}{
	(Protocol)(0), // 0: hashicorp.consul.mesh.v2beta1.pbproxystate.Protocol
}
var file_pbmesh_v2beta1_pbproxystate_protocol_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_pbproxystate_protocol_proto_init() }
func file_pbmesh_v2beta1_pbproxystate_protocol_proto_init() {
	if File_pbmesh_v2beta1_pbproxystate_protocol_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_pbproxystate_protocol_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_pbproxystate_protocol_proto_depIdxs,
		EnumInfos:         file_pbmesh_v2beta1_pbproxystate_protocol_proto_enumTypes,
	}.Build()
	File_pbmesh_v2beta1_pbproxystate_protocol_proto = out.File
	file_pbmesh_v2beta1_pbproxystate_protocol_proto_rawDesc = nil
	file_pbmesh_v2beta1_pbproxystate_protocol_proto_goTypes = nil
	file_pbmesh_v2beta1_pbproxystate_protocol_proto_depIdxs = nil
}