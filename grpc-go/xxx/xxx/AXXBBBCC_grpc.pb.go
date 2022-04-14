// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xxx

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

// AXXBBBCCClient is the client API for AXXBBBCC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AXXBBBCCClient interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type aXXBBBCCClient struct {
	cc grpc.ClientConnInterface
}

func NewAXXBBBCCClient(cc grpc.ClientConnInterface) AXXBBBCCClient {
	return &aXXBBBCCClient{cc}
}

func (c *aXXBBBCCClient) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/xxx.xxx.AXXBBBCC/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aXXBBBCCClient) Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxx.xxx.AXXBBBCC/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AXXBBBCCServer is the server API for AXXBBBCC service.
// All implementations must embed UnimplementedAXXBBBCCServer
// for forward compatibility
type AXXBBBCCServer interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedAXXBBBCCServer()
}

// UnimplementedAXXBBBCCServer must be embedded to have forward compatible implementations.
type UnimplementedAXXBBBCCServer struct {
}

func (UnimplementedAXXBBBCCServer) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedAXXBBBCCServer) Helloworld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedAXXBBBCCServer) mustEmbedUnimplementedAXXBBBCCServer() {}

// UnsafeAXXBBBCCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AXXBBBCCServer will
// result in compilation errors.
type UnsafeAXXBBBCCServer interface {
	mustEmbedUnimplementedAXXBBBCCServer()
}

func RegisterAXXBBBCCServer(s grpc.ServiceRegistrar, srv AXXBBBCCServer) {
	s.RegisterService(&AXXBBBCC_ServiceDesc, srv)
}

func _AXXBBBCC_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AXXBBBCCServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxx.xxx.AXXBBBCC/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AXXBBBCCServer).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AXXBBBCC_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AXXBBBCCServer).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxx.xxx.AXXBBBCC/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AXXBBBCCServer).Helloworld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AXXBBBCC_ServiceDesc is the grpc.ServiceDesc for AXXBBBCC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AXXBBBCC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxx.xxx.AXXBBBCC",
	HandlerType: (*AXXBBBCCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _AXXBBBCC_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _AXXBBBCC_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AXXBBBCC.proto",
}
