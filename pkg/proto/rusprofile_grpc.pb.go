// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/rusprofile.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RusProfileClient is the client API for RusProfile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RusProfileClient interface {
	GetRusProfileData(ctx context.Context, in *RpRequest, opts ...grpc.CallOption) (*RpResponse, error)
}

type rusProfileClient struct {
	cc grpc.ClientConnInterface
}

func NewRusProfileClient(cc grpc.ClientConnInterface) RusProfileClient {
	return &rusProfileClient{cc}
}

func (c *rusProfileClient) GetRusProfileData(ctx context.Context, in *RpRequest, opts ...grpc.CallOption) (*RpResponse, error) {
	out := new(RpResponse)
	err := c.cc.Invoke(ctx, "/api.RusProfile/GetRusProfileData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RusProfileServer is the server API for RusProfile service.
// All implementations must embed UnimplementedRusProfileServer
// for forward compatibility
type RusProfileServer interface {
	GetRusProfileData(context.Context, *RpRequest) (*RpResponse, error)
	mustEmbedUnimplementedRusProfileServer()
}

// UnimplementedRusProfileServer must be embedded to have forward compatible implementations.
type UnimplementedRusProfileServer struct {
}

func (UnimplementedRusProfileServer) GetRusProfileData(context.Context, *RpRequest) (*RpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRusProfileData not implemented")
}
func (UnimplementedRusProfileServer) mustEmbedUnimplementedRusProfileServer() {}

// UnsafeRusProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RusProfileServer will
// result in compilation errors.
type UnsafeRusProfileServer interface {
	mustEmbedUnimplementedRusProfileServer()
}

func RegisterRusProfileServer(s grpc.ServiceRegistrar, srv RusProfileServer) {
	s.RegisterService(&RusProfile_ServiceDesc, srv)
}

func _RusProfile_GetRusProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RusProfileServer).GetRusProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RusProfile/GetRusProfileData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RusProfileServer).GetRusProfileData(ctx, req.(*RpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RusProfile_ServiceDesc is the grpc.ServiceDesc for RusProfile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RusProfile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RusProfile",
	HandlerType: (*RusProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRusProfileData",
			Handler:    _RusProfile_GetRusProfileData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rusprofile.proto",
}
