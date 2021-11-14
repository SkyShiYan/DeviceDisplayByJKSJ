// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v4

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

// LayoutClient is the client API for Layout service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LayoutClient interface {
	// 获取获取播放资源列表
	GetDeviceResourceList(ctx context.Context, in *LayoutResourceRequest, opts ...grpc.CallOption) (*LayoutResourceReply, error)
}

type layoutClient struct {
	cc grpc.ClientConnInterface
}

func NewLayoutClient(cc grpc.ClientConnInterface) LayoutClient {
	return &layoutClient{cc}
}

func (c *layoutClient) GetDeviceResourceList(ctx context.Context, in *LayoutResourceRequest, opts ...grpc.CallOption) (*LayoutResourceReply, error) {
	out := new(LayoutResourceReply)
	err := c.cc.Invoke(ctx, "/layout.v4.Layout/getDeviceResourceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LayoutServer is the server API for Layout service.
// All implementations must embed UnimplementedLayoutServer
// for forward compatibility
type LayoutServer interface {
	// 获取获取播放资源列表
	GetDeviceResourceList(context.Context, *LayoutResourceRequest) (*LayoutResourceReply, error)
	mustEmbedUnimplementedLayoutServer()
}

// UnimplementedLayoutServer must be embedded to have forward compatible implementations.
type UnimplementedLayoutServer struct {
}

func (UnimplementedLayoutServer) GetDeviceResourceList(context.Context, *LayoutResourceRequest) (*LayoutResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceResourceList not implemented")
}
func (UnimplementedLayoutServer) mustEmbedUnimplementedLayoutServer() {}

// UnsafeLayoutServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayoutServer will
// result in compilation errors.
type UnsafeLayoutServer interface {
	mustEmbedUnimplementedLayoutServer()
}

func RegisterLayoutServer(s grpc.ServiceRegistrar, srv LayoutServer) {
	s.RegisterService(&Layout_ServiceDesc, srv)
}

func _Layout_GetDeviceResourceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LayoutResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayoutServer).GetDeviceResourceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/layout.v4.Layout/getDeviceResourceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayoutServer).GetDeviceResourceList(ctx, req.(*LayoutResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Layout_ServiceDesc is the grpc.ServiceDesc for Layout service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Layout_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "layout.v4.Layout",
	HandlerType: (*LayoutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getDeviceResourceList",
			Handler:    _Layout_GetDeviceResourceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/layout/v4/layout.proto",
}
