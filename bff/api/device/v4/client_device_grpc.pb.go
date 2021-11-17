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

// DeviceClient is the client API for Device service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceClient interface {
	// 获取设备
	GetDeviceByKey(ctx context.Context, in *GetDeviceByKeyRequest, opts ...grpc.CallOption) (*GetDeviceByKeyReply, error)
	// 新增设备信息
	AddDeviceByKey(ctx context.Context, in *AddDeviceByKeyRequest, opts ...grpc.CallOption) (*AddDeviceByKeyReply, error)
	// 更新设备信息
	UpdateDeviceByKey(ctx context.Context, in *UpdateDeviceByKeyRequest, opts ...grpc.CallOption) (*UpdateDeviceByKeyReply, error)
}

type deviceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceClient(cc grpc.ClientConnInterface) DeviceClient {
	return &deviceClient{cc}
}

func (c *deviceClient) GetDeviceByKey(ctx context.Context, in *GetDeviceByKeyRequest, opts ...grpc.CallOption) (*GetDeviceByKeyReply, error) {
	out := new(GetDeviceByKeyReply)
	err := c.cc.Invoke(ctx, "/device.v4.Device/getDeviceByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) AddDeviceByKey(ctx context.Context, in *AddDeviceByKeyRequest, opts ...grpc.CallOption) (*AddDeviceByKeyReply, error) {
	out := new(AddDeviceByKeyReply)
	err := c.cc.Invoke(ctx, "/device.v4.Device/addDeviceByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceClient) UpdateDeviceByKey(ctx context.Context, in *UpdateDeviceByKeyRequest, opts ...grpc.CallOption) (*UpdateDeviceByKeyReply, error) {
	out := new(UpdateDeviceByKeyReply)
	err := c.cc.Invoke(ctx, "/device.v4.Device/updateDeviceByKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServer is the server API for Device service.
// All implementations must embed UnimplementedDeviceServer
// for forward compatibility
type DeviceServer interface {
	// 获取设备
	GetDeviceByKey(context.Context, *GetDeviceByKeyRequest) (*GetDeviceByKeyReply, error)
	// 新增设备信息
	AddDeviceByKey(context.Context, *AddDeviceByKeyRequest) (*AddDeviceByKeyReply, error)
	// 更新设备信息
	UpdateDeviceByKey(context.Context, *UpdateDeviceByKeyRequest) (*UpdateDeviceByKeyReply, error)
	mustEmbedUnimplementedDeviceServer()
}

// UnimplementedDeviceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceServer struct {
}

func (UnimplementedDeviceServer) GetDeviceByKey(context.Context, *GetDeviceByKeyRequest) (*GetDeviceByKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceByKey not implemented")
}
func (UnimplementedDeviceServer) AddDeviceByKey(context.Context, *AddDeviceByKeyRequest) (*AddDeviceByKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeviceByKey not implemented")
}
func (UnimplementedDeviceServer) UpdateDeviceByKey(context.Context, *UpdateDeviceByKeyRequest) (*UpdateDeviceByKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceByKey not implemented")
}
func (UnimplementedDeviceServer) mustEmbedUnimplementedDeviceServer() {}

// UnsafeDeviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServer will
// result in compilation errors.
type UnsafeDeviceServer interface {
	mustEmbedUnimplementedDeviceServer()
}

func RegisterDeviceServer(s grpc.ServiceRegistrar, srv DeviceServer) {
	s.RegisterService(&Device_ServiceDesc, srv)
}

func _Device_GetDeviceByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).GetDeviceByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.v4.Device/getDeviceByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).GetDeviceByKey(ctx, req.(*GetDeviceByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_AddDeviceByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).AddDeviceByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.v4.Device/addDeviceByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).AddDeviceByKey(ctx, req.(*AddDeviceByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Device_UpdateDeviceByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceByKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServer).UpdateDeviceByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device.v4.Device/updateDeviceByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServer).UpdateDeviceByKey(ctx, req.(*UpdateDeviceByKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Device_ServiceDesc is the grpc.ServiceDesc for Device service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Device_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device.v4.Device",
	HandlerType: (*DeviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getDeviceByKey",
			Handler:    _Device_GetDeviceByKey_Handler,
		},
		{
			MethodName: "addDeviceByKey",
			Handler:    _Device_AddDeviceByKey_Handler,
		},
		{
			MethodName: "updateDeviceByKey",
			Handler:    _Device_UpdateDeviceByKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/device/v4/client_device.proto",
}
