// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: services/svcdeviceregistry/deviceregistry.proto

package svcdeviceregistry

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	common "github.com/cybroslabs/hes-2-apis/gen/go/common"
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

var File_services_svcdeviceregistry_deviceregistry_proto protoreflect.FileDescriptor

var file_services_svcdeviceregistry_deviceregistry_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x76, 0x63, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2a, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x76, 0x63, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x63, 0x71, 0x75,
	0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa1, 0x12, 0x0a, 0x15, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x7b,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x42, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x1a, 0x3b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71,
	0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x6f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x65, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x37, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x1a, 0x30, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63,
	0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a,
	0x1b, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x46, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x78, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x3b, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x6f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3c, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x76, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2b, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x5f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x2f, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x69, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x54, 0x6f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61,
	0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x73, 0x0a, 0x16, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x41, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61,
	0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x72, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a,
	0x33, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4d, 0x6f, 0x64, 0x65, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x5b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75,
	0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f,
	0x6c, 0x12, 0x68, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x37, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63,
	0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x37,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x47, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x60, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x12, 0x33, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x12, 0x33, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x76, 0x63, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_services_svcdeviceregistry_deviceregistry_proto_goTypes = []any{
	(*acquisition.SetDriver)(nil),                          // 0: io.clbs.openhes.models.acquisition.SetDriver
	(*acquisition.CreateCommunicationUnitRequest)(nil),     // 1: io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
	(*common.ListSelector)(nil),                            // 2: io.clbs.openhes.models.common.ListSelector
	(*wrapperspb.StringValue)(nil),                         // 3: google.protobuf.StringValue
	(*acquisition.CreateDeviceRequest)(nil),                // 4: io.clbs.openhes.models.acquisition.CreateDeviceRequest
	(*acquisition.SetDeviceCommunicationUnitsRequest)(nil), // 5: io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
	(*acquisition.CreateDeviceGroupRequest)(nil),           // 6: io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
	(*acquisition.AddDevicesToGroupRequest)(nil),           // 7: io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
	(*acquisition.RemoveDevicesFromGroupRequest)(nil),      // 8: io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
	(*acquisition.SetModemPoolRequest)(nil),                // 9: io.clbs.openhes.models.acquisition.SetModemPoolRequest
	(*acquisition.SetModemRequest)(nil),                    // 10: io.clbs.openhes.models.acquisition.SetModemRequest
	(*emptypb.Empty)(nil),                                  // 11: google.protobuf.Empty
	(*acquisition.ListOfCommunicationUnit)(nil),            // 12: io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	(*acquisition.CommunicationUnitSpec)(nil),              // 13: io.clbs.openhes.models.acquisition.CommunicationUnitSpec
	(*acquisition.ListOfDevice)(nil),                       // 14: io.clbs.openhes.models.acquisition.ListOfDevice
	(*acquisition.Device)(nil),                             // 15: io.clbs.openhes.models.acquisition.Device
	(*acquisition.ListOfDeviceGroup)(nil),                  // 16: io.clbs.openhes.models.acquisition.ListOfDeviceGroup
	(*acquisition.DeviceGroup)(nil),                        // 17: io.clbs.openhes.models.acquisition.DeviceGroup
	(*acquisition.ListOfModemPool)(nil),                    // 18: io.clbs.openhes.models.acquisition.ListOfModemPool
	(*acquisition.ModemPool)(nil),                          // 19: io.clbs.openhes.models.acquisition.ModemPool
}
var file_services_svcdeviceregistry_deviceregistry_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDriverTemplates:input_type -> io.clbs.openhes.models.acquisition.SetDriver
	1,  // 1: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateCommunicationUnit:input_type -> io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
	2,  // 2: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListCommunicationUnits:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 3: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetCommunicationUnit:input_type -> google.protobuf.StringValue
	4,  // 4: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDevice:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceRequest
	2,  // 5: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDevices:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 6: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDevice:input_type -> google.protobuf.StringValue
	5,  // 7: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDeviceCommunicationUnits:input_type -> io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
	3,  // 8: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceCommunicationUnits:input_type -> google.protobuf.StringValue
	6,  // 9: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceGroup:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
	2,  // 10: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceGroups:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 11: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceGroup:input_type -> google.protobuf.StringValue
	7,  // 12: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddDevicesToGroup:input_type -> io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
	8,  // 13: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveDevicesFromGroup:input_type -> io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
	2,  // 14: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListModemPools:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 15: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetModemPool:input_type -> google.protobuf.StringValue
	9,  // 16: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModemPool:input_type -> io.clbs.openhes.models.acquisition.SetModemPoolRequest
	9,  // 17: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModemPool:input_type -> io.clbs.openhes.models.acquisition.SetModemPoolRequest
	3,  // 18: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModemPool:input_type -> google.protobuf.StringValue
	10, // 19: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModem:input_type -> io.clbs.openhes.models.acquisition.SetModemRequest
	10, // 20: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModem:input_type -> io.clbs.openhes.models.acquisition.SetModemRequest
	3,  // 21: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModem:input_type -> google.protobuf.StringValue
	11, // 22: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDriverTemplates:output_type -> google.protobuf.Empty
	3,  // 23: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateCommunicationUnit:output_type -> google.protobuf.StringValue
	12, // 24: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListCommunicationUnits:output_type -> io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	13, // 25: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetCommunicationUnit:output_type -> io.clbs.openhes.models.acquisition.CommunicationUnitSpec
	3,  // 26: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDevice:output_type -> google.protobuf.StringValue
	14, // 27: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDevices:output_type -> io.clbs.openhes.models.acquisition.ListOfDevice
	15, // 28: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDevice:output_type -> io.clbs.openhes.models.acquisition.Device
	11, // 29: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDeviceCommunicationUnits:output_type -> google.protobuf.Empty
	12, // 30: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceCommunicationUnits:output_type -> io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	3,  // 31: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceGroup:output_type -> google.protobuf.StringValue
	16, // 32: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceGroups:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceGroup
	17, // 33: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceGroup:output_type -> io.clbs.openhes.models.acquisition.DeviceGroup
	11, // 34: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddDevicesToGroup:output_type -> google.protobuf.Empty
	11, // 35: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveDevicesFromGroup:output_type -> google.protobuf.Empty
	18, // 36: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListModemPools:output_type -> io.clbs.openhes.models.acquisition.ListOfModemPool
	19, // 37: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetModemPool:output_type -> io.clbs.openhes.models.acquisition.ModemPool
	3,  // 38: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModemPool:output_type -> google.protobuf.StringValue
	11, // 39: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModemPool:output_type -> google.protobuf.Empty
	11, // 40: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModemPool:output_type -> google.protobuf.Empty
	3,  // 41: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModem:output_type -> google.protobuf.StringValue
	11, // 42: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModem:output_type -> google.protobuf.Empty
	11, // 43: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModem:output_type -> google.protobuf.Empty
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_svcdeviceregistry_deviceregistry_proto_init() }
func file_services_svcdeviceregistry_deviceregistry_proto_init() {
	if File_services_svcdeviceregistry_deviceregistry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_svcdeviceregistry_deviceregistry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_svcdeviceregistry_deviceregistry_proto_goTypes,
		DependencyIndexes: file_services_svcdeviceregistry_deviceregistry_proto_depIdxs,
	}.Build()
	File_services_svcdeviceregistry_deviceregistry_proto = out.File
	file_services_svcdeviceregistry_deviceregistry_proto_rawDesc = nil
	file_services_svcdeviceregistry_deviceregistry_proto_goTypes = nil
	file_services_svcdeviceregistry_deviceregistry_proto_depIdxs = nil
}
