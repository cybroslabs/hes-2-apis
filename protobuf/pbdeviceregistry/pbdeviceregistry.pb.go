// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: pbdeviceregistry.proto

package pbdeviceregistry

import (
	pbdriver "github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pbdeviceregistry_proto protoreflect.FileDescriptor

var file_pbdeviceregistry_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x10,
	0x0a, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x73, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x40, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70,
	0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x98, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x12, 0x3e, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3f, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x33,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7b, 0x0a, 0x1b, 0x53, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x44, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0xad, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x45, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3a, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x3a, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x54, 0x6f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x71, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3f,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46,
	0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x37, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x36, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x60, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x12, 0x31, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x12, 0x31, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68,
	0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x62, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pbdeviceregistry_proto_goTypes = []any{
	(*pbdriver.NegotiateRequest)(nil),            // 0: io.clbs.openhes.pbdriver.NegotiateRequest
	(*CreateCommunicationUnitRequest)(nil),       // 1: io.clbs.openhes.pbdeviceregistry.CreateCommunicationUnitRequest
	(*GetCommunicationUnitsRequest)(nil),         // 2: io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsRequest
	(*CreateDeviceRequest)(nil),                  // 3: io.clbs.openhes.pbdeviceregistry.CreateDeviceRequest
	(*GetDevicesRequest)(nil),                    // 4: io.clbs.openhes.pbdeviceregistry.GetDevicesRequest
	(*SetDeviceCommunicationUnitsRequest)(nil),   // 5: io.clbs.openhes.pbdeviceregistry.SetDeviceCommunicationUnitsRequest
	(*GetDevicesCommunicationUnitsRequest)(nil),  // 6: io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsRequest
	(*CreateDeviceGroupRequest)(nil),             // 7: io.clbs.openhes.pbdeviceregistry.CreateDeviceGroupRequest
	(*wrapperspb.StringValue)(nil),               // 8: google.protobuf.StringValue
	(*AddDevicesToGroupRequest)(nil),             // 9: io.clbs.openhes.pbdeviceregistry.AddDevicesToGroupRequest
	(*RemoveDevicesFromGroupRequest)(nil),        // 10: io.clbs.openhes.pbdeviceregistry.RemoveDevicesFromGroupRequest
	(*emptypb.Empty)(nil),                        // 11: google.protobuf.Empty
	(*SetModemPoolRequest)(nil),                  // 12: io.clbs.openhes.pbdeviceregistry.SetModemPoolRequest
	(*SetModemRequest)(nil),                      // 13: io.clbs.openhes.pbdeviceregistry.SetModemRequest
	(*GetCommunicationUnitsResponse)(nil),        // 14: io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsResponse
	(*GetDevicesResponse)(nil),                   // 15: io.clbs.openhes.pbdeviceregistry.GetDevicesResponse
	(*GetDevicesCommunicationUnitsResponse)(nil), // 16: io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsResponse
	(*GetDeviceGroupsResponse)(nil),              // 17: io.clbs.openhes.pbdeviceregistry.GetDeviceGroupsResponse
	(*GetModemPoolsResponse)(nil),                // 18: io.clbs.openhes.pbdeviceregistry.GetModemPoolsResponse
	(*GetModemPoolResponse)(nil),                 // 19: io.clbs.openhes.pbdeviceregistry.GetModemPoolResponse
}
var file_pbdeviceregistry_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.SetDriverTemplates:input_type -> io.clbs.openhes.pbdriver.NegotiateRequest
	1,  // 1: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateCommunicationUnit:input_type -> io.clbs.openhes.pbdeviceregistry.CreateCommunicationUnitRequest
	2,  // 2: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetCommunicationUnits:input_type -> io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsRequest
	3,  // 3: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateDevice:input_type -> io.clbs.openhes.pbdeviceregistry.CreateDeviceRequest
	4,  // 4: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetDevices:input_type -> io.clbs.openhes.pbdeviceregistry.GetDevicesRequest
	5,  // 5: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.SetDeviceCommunicationUnits:input_type -> io.clbs.openhes.pbdeviceregistry.SetDeviceCommunicationUnitsRequest
	6,  // 6: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetDevicesCommunicationUnits:input_type -> io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsRequest
	7,  // 7: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateDeviceGroup:input_type -> io.clbs.openhes.pbdeviceregistry.CreateDeviceGroupRequest
	8,  // 8: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetDeviceGroups:input_type -> google.protobuf.StringValue
	9,  // 9: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.AddDevicesToGroup:input_type -> io.clbs.openhes.pbdeviceregistry.AddDevicesToGroupRequest
	10, // 10: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.RemoveDevicesFromGroup:input_type -> io.clbs.openhes.pbdeviceregistry.RemoveDevicesFromGroupRequest
	11, // 11: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetModemPools:input_type -> google.protobuf.Empty
	8,  // 12: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetModemPool:input_type -> google.protobuf.StringValue
	12, // 13: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateModemPool:input_type -> io.clbs.openhes.pbdeviceregistry.SetModemPoolRequest
	12, // 14: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.UpdateModemPool:input_type -> io.clbs.openhes.pbdeviceregistry.SetModemPoolRequest
	8,  // 15: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.DeleteModemPool:input_type -> google.protobuf.StringValue
	13, // 16: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateModem:input_type -> io.clbs.openhes.pbdeviceregistry.SetModemRequest
	13, // 17: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.UpdateModem:input_type -> io.clbs.openhes.pbdeviceregistry.SetModemRequest
	8,  // 18: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.DeleteModem:input_type -> google.protobuf.StringValue
	11, // 19: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.SetDriverTemplates:output_type -> google.protobuf.Empty
	11, // 20: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateCommunicationUnit:output_type -> google.protobuf.Empty
	14, // 21: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetCommunicationUnits:output_type -> io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsResponse
	11, // 22: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateDevice:output_type -> google.protobuf.Empty
	15, // 23: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetDevices:output_type -> io.clbs.openhes.pbdeviceregistry.GetDevicesResponse
	11, // 24: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.SetDeviceCommunicationUnits:output_type -> google.protobuf.Empty
	16, // 25: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetDevicesCommunicationUnits:output_type -> io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsResponse
	11, // 26: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateDeviceGroup:output_type -> google.protobuf.Empty
	17, // 27: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetDeviceGroups:output_type -> io.clbs.openhes.pbdeviceregistry.GetDeviceGroupsResponse
	11, // 28: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.AddDevicesToGroup:output_type -> google.protobuf.Empty
	11, // 29: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.RemoveDevicesFromGroup:output_type -> google.protobuf.Empty
	18, // 30: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetModemPools:output_type -> io.clbs.openhes.pbdeviceregistry.GetModemPoolsResponse
	19, // 31: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.GetModemPool:output_type -> io.clbs.openhes.pbdeviceregistry.GetModemPoolResponse
	11, // 32: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateModemPool:output_type -> google.protobuf.Empty
	11, // 33: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.UpdateModemPool:output_type -> google.protobuf.Empty
	11, // 34: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.DeleteModemPool:output_type -> google.protobuf.Empty
	11, // 35: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.CreateModem:output_type -> google.protobuf.Empty
	11, // 36: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.UpdateModem:output_type -> google.protobuf.Empty
	11, // 37: io.clbs.openhes.pbdeviceregistry.DeviceRegistryService.DeleteModem:output_type -> google.protobuf.Empty
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pbdeviceregistry_proto_init() }
func file_pbdeviceregistry_proto_init() {
	if File_pbdeviceregistry_proto != nil {
		return
	}
	file_pbdeviceregistry_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbdeviceregistry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbdeviceregistry_proto_goTypes,
		DependencyIndexes: file_pbdeviceregistry_proto_depIdxs,
	}.Build()
	File_pbdeviceregistry_proto = out.File
	file_pbdeviceregistry_proto_rawDesc = nil
	file_pbdeviceregistry_proto_goTypes = nil
	file_pbdeviceregistry_proto_depIdxs = nil
}
