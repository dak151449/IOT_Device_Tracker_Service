// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: api/device_tracker/device_tracker.proto

package dtapi

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

// DeviceTrackerServiceClient is the client API for DeviceTrackerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceTrackerServiceClient interface {
	GetDeviceGroups(ctx context.Context, in *GetDeviceGroupsRequest, opts ...grpc.CallOption) (*GetDeviceGroupsResponse, error)
	GetDevicesFromGroup(ctx context.Context, in *GetDevicesFromGroupRequest, opts ...grpc.CallOption) (*GetDevicesFromGroupResponse, error)
	CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*CreateDeviceGroupResponse, error)
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error)
}

type deviceTrackerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceTrackerServiceClient(cc grpc.ClientConnInterface) DeviceTrackerServiceClient {
	return &deviceTrackerServiceClient{cc}
}

func (c *deviceTrackerServiceClient) GetDeviceGroups(ctx context.Context, in *GetDeviceGroupsRequest, opts ...grpc.CallOption) (*GetDeviceGroupsResponse, error) {
	out := new(GetDeviceGroupsResponse)
	err := c.cc.Invoke(ctx, "/device_tracker.DeviceTrackerService/GetDeviceGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceTrackerServiceClient) GetDevicesFromGroup(ctx context.Context, in *GetDevicesFromGroupRequest, opts ...grpc.CallOption) (*GetDevicesFromGroupResponse, error) {
	out := new(GetDevicesFromGroupResponse)
	err := c.cc.Invoke(ctx, "/device_tracker.DeviceTrackerService/GetDevicesFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceTrackerServiceClient) CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*CreateDeviceGroupResponse, error) {
	out := new(CreateDeviceGroupResponse)
	err := c.cc.Invoke(ctx, "/device_tracker.DeviceTrackerService/CreateDeviceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceTrackerServiceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*CreateDeviceResponse, error) {
	out := new(CreateDeviceResponse)
	err := c.cc.Invoke(ctx, "/device_tracker.DeviceTrackerService/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceTrackerServiceServer is the server API for DeviceTrackerService service.
// All implementations must embed UnimplementedDeviceTrackerServiceServer
// for forward compatibility
type DeviceTrackerServiceServer interface {
	GetDeviceGroups(context.Context, *GetDeviceGroupsRequest) (*GetDeviceGroupsResponse, error)
	GetDevicesFromGroup(context.Context, *GetDevicesFromGroupRequest) (*GetDevicesFromGroupResponse, error)
	CreateDeviceGroup(context.Context, *CreateDeviceGroupRequest) (*CreateDeviceGroupResponse, error)
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
	mustEmbedUnimplementedDeviceTrackerServiceServer()
}

// UnimplementedDeviceTrackerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceTrackerServiceServer struct {
}

func (UnimplementedDeviceTrackerServiceServer) GetDeviceGroups(context.Context, *GetDeviceGroupsRequest) (*GetDeviceGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceGroups not implemented")
}
func (UnimplementedDeviceTrackerServiceServer) GetDevicesFromGroup(context.Context, *GetDevicesFromGroupRequest) (*GetDevicesFromGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevicesFromGroup not implemented")
}
func (UnimplementedDeviceTrackerServiceServer) CreateDeviceGroup(context.Context, *CreateDeviceGroupRequest) (*CreateDeviceGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceGroup not implemented")
}
func (UnimplementedDeviceTrackerServiceServer) CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedDeviceTrackerServiceServer) mustEmbedUnimplementedDeviceTrackerServiceServer() {}

// UnsafeDeviceTrackerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceTrackerServiceServer will
// result in compilation errors.
type UnsafeDeviceTrackerServiceServer interface {
	mustEmbedUnimplementedDeviceTrackerServiceServer()
}

func RegisterDeviceTrackerServiceServer(s grpc.ServiceRegistrar, srv DeviceTrackerServiceServer) {
	s.RegisterService(&DeviceTrackerService_ServiceDesc, srv)
}

func _DeviceTrackerService_GetDeviceGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceTrackerServiceServer).GetDeviceGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device_tracker.DeviceTrackerService/GetDeviceGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceTrackerServiceServer).GetDeviceGroups(ctx, req.(*GetDeviceGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceTrackerService_GetDevicesFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDevicesFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceTrackerServiceServer).GetDevicesFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device_tracker.DeviceTrackerService/GetDevicesFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceTrackerServiceServer).GetDevicesFromGroup(ctx, req.(*GetDevicesFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceTrackerService_CreateDeviceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceTrackerServiceServer).CreateDeviceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device_tracker.DeviceTrackerService/CreateDeviceGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceTrackerServiceServer).CreateDeviceGroup(ctx, req.(*CreateDeviceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceTrackerService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceTrackerServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/device_tracker.DeviceTrackerService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceTrackerServiceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceTrackerService_ServiceDesc is the grpc.ServiceDesc for DeviceTrackerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceTrackerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "device_tracker.DeviceTrackerService",
	HandlerType: (*DeviceTrackerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceGroups",
			Handler:    _DeviceTrackerService_GetDeviceGroups_Handler,
		},
		{
			MethodName: "GetDevicesFromGroup",
			Handler:    _DeviceTrackerService_GetDevicesFromGroup_Handler,
		},
		{
			MethodName: "CreateDeviceGroup",
			Handler:    _DeviceTrackerService_CreateDeviceGroup_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _DeviceTrackerService_CreateDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/device_tracker/device_tracker.proto",
}
