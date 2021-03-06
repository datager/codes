// Code generated by protoc-gen-language.
// source: proxy.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proxy.proto

It has these top-level messages:
	WeedFileArrayMessage
	WeedFileMapId
	WeedFileMapMessage
*/
package dg_model

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type WeedFileArrayMessage struct {
	ArrayData []string `protobuf:"bytes,1,rep,name=ArrayData,json=arrayData" json:"ArrayData"`
}

func (m *WeedFileArrayMessage) Reset()                    { *m = WeedFileArrayMessage{} }
func (m *WeedFileArrayMessage) String() string            { return proto1.CompactTextString(m) }
func (*WeedFileArrayMessage) ProtoMessage()               {}
func (*WeedFileArrayMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WeedFileArrayMessage) GetArrayData() []string {
	if m != nil {
		return m.ArrayData
	}
	return nil
}

type WeedFileMapId struct {
	MapIds map[string]string `protobuf:"bytes,1,rep,name=MapIds,json=mapIds" json:"MapIds" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WeedFileMapId) Reset()                    { *m = WeedFileMapId{} }
func (m *WeedFileMapId) String() string            { return proto1.CompactTextString(m) }
func (*WeedFileMapId) ProtoMessage()               {}
func (*WeedFileMapId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WeedFileMapId) GetMapIds() map[string]string {
	if m != nil {
		return m.MapIds
	}
	return nil
}

type WeedFileMapMessage struct {
	MapData map[string][]byte `protobuf:"bytes,1,rep,name=MapData,json=mapData" json:"MapData" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *WeedFileMapMessage) Reset()                    { *m = WeedFileMapMessage{} }
func (m *WeedFileMapMessage) String() string            { return proto1.CompactTextString(m) }
func (*WeedFileMapMessage) ProtoMessage()               {}
func (*WeedFileMapMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WeedFileMapMessage) GetMapData() map[string][]byte {
	if m != nil {
		return m.MapData
	}
	return nil
}

func init() {
	proto1.RegisterType((*WeedFileArrayMessage)(nil), "proto.WeedFileArrayMessage")
	proto1.RegisterType((*WeedFileMapId)(nil), "proto.WeedFileMapId")
	proto1.RegisterType((*WeedFileMapMessage)(nil), "proto.WeedFileMapMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GrpcService service

type GrpcServiceClient interface {
	// ??????Fid??????????????????
	GetFileContent(ctx context.Context, in *WeedFileArrayMessage, opts ...grpc.CallOption) (*WeedFileMapMessage, error)
	// ???????????????????????????Fid
	PostFileContent(ctx context.Context, in *WeedFileMapMessage, opts ...grpc.CallOption) (*WeedFileMapId, error)
	// ????????????Fid???????????????????????????
	DeleteFile(ctx context.Context, in *WeedFileArrayMessage, opts ...grpc.CallOption) (*WeedFileArrayMessage, error)
}

type grpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewGrpcServiceClient(cc *grpc.ClientConn) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) GetFileContent(ctx context.Context, in *WeedFileArrayMessage, opts ...grpc.CallOption) (*WeedFileMapMessage, error) {
	out := new(WeedFileMapMessage)
	err := grpc.Invoke(ctx, "/proto.GrpcService/GetFileContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServiceClient) PostFileContent(ctx context.Context, in *WeedFileMapMessage, opts ...grpc.CallOption) (*WeedFileMapId, error) {
	out := new(WeedFileMapId)
	err := grpc.Invoke(ctx, "/proto.GrpcService/PostFileContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServiceClient) DeleteFile(ctx context.Context, in *WeedFileArrayMessage, opts ...grpc.CallOption) (*WeedFileArrayMessage, error) {
	out := new(WeedFileArrayMessage)
	err := grpc.Invoke(ctx, "/proto.GrpcService/DeleteFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GrpcService service

type GrpcServiceServer interface {
	// ??????Fid??????????????????
	GetFileContent(context.Context, *WeedFileArrayMessage) (*WeedFileMapMessage, error)
	// ???????????????????????????Fid
	PostFileContent(context.Context, *WeedFileMapMessage) (*WeedFileMapId, error)
	// ????????????Fid???????????????????????????
	DeleteFile(context.Context, *WeedFileArrayMessage) (*WeedFileArrayMessage, error)
}

func RegisterGrpcServiceServer(s *grpc.Server, srv GrpcServiceServer) {
	s.RegisterService(&_GrpcService_serviceDesc, srv)
}

func _GrpcService_GetFileContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeedFileArrayMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GetFileContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GrpcService/GetFileContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GetFileContent(ctx, req.(*WeedFileArrayMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcService_PostFileContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeedFileMapMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).PostFileContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GrpcService/PostFileContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).PostFileContent(ctx, req.(*WeedFileMapMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeedFileArrayMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GrpcService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).DeleteFile(ctx, req.(*WeedFileArrayMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GrpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileContent",
			Handler:    _GrpcService_GetFileContent_Handler,
		},
		{
			MethodName: "PostFileContent",
			Handler:    _GrpcService_PostFileContent_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _GrpcService_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy.proto",
}

func init() { proto1.RegisterFile("proxy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0xaf,
	0xa8, 0xd4, 0x03, 0x92, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x4a, 0xc9, 0x84, 0x4b, 0x24, 0x3c, 0x35,
	0x35, 0xc5, 0x2d, 0x33, 0x27, 0xd5, 0xb1, 0xa8, 0x28, 0xb1, 0xd2, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x3d, 0x55, 0x48, 0x86, 0x8b, 0x13, 0xcc, 0x77, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x0c, 0xe2, 0x4c, 0x84, 0x09, 0x28, 0xb5, 0x30, 0x72, 0xf1, 0xc2, 0xb4, 0xf9, 0x26, 0x16,
	0x78, 0xa6, 0x08, 0x59, 0x70, 0xb1, 0x81, 0x19, 0xc5, 0x60, 0xc5, 0xdc, 0x46, 0x0a, 0x10, 0x6b,
	0xf4, 0x50, 0x54, 0xe9, 0x41, 0x94, 0xb8, 0xe6, 0x95, 0x14, 0x55, 0x06, 0xb1, 0xe5, 0x82, 0x39,
	0x52, 0x96, 0x5c, 0xdc, 0x48, 0xc2, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x40, 0x53, 0x18,
	0x81, 0x56, 0x82, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60,
	0x31, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0x51, 0x69, 0x12, 0x23, 0x97, 0x10, 0x92, 0x05, 0x30, 0xb7,
	0x3b, 0x70, 0xb1, 0x03, 0x79, 0x70, 0x97, 0x73, 0x1b, 0xa9, 0x61, 0x3a, 0x06, 0xaa, 0x56, 0x0f,
	0xaa, 0x10, 0xe2, 0x24, 0xf6, 0x5c, 0x08, 0x4f, 0xca, 0x8a, 0x8b, 0x07, 0x59, 0x82, 0x90, 0xa3,
	0x78, 0x90, 0x1c, 0x65, 0xf4, 0x92, 0x91, 0x8b, 0xdb, 0xbd, 0xa8, 0x20, 0x39, 0x38, 0xb5, 0xa8,
	0x2c, 0x33, 0x39, 0x55, 0xc8, 0x8b, 0x8b, 0xcf, 0x3d, 0xb5, 0x04, 0x64, 0xad, 0x73, 0x7e, 0x5e,
	0x49, 0x6a, 0x5e, 0x89, 0x90, 0x34, 0x9a, 0x73, 0x90, 0x03, 0x5e, 0x4a, 0x12, 0xa7, 0x5b, 0x95,
	0x18, 0x84, 0x5c, 0xb8, 0xf8, 0x03, 0xf2, 0x8b, 0x51, 0x0c, 0xc3, 0xad, 0x5e, 0x4a, 0x04, 0x5b,
	0x1c, 0x00, 0x4d, 0xf1, 0xe0, 0xe2, 0x72, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x05, 0x09, 0xe2, 0x77,
	0x0d, 0x3e, 0x49, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0xf6, 0x69, 0xde, 0x5a, 0x02, 0x00, 0x00,
}
