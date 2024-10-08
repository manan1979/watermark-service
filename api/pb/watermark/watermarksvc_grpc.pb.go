// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: watermarksvc.proto

package watermark

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Watermark_Get_FullMethodName           = "/pb.Watermark/Get"
	Watermark_Watermark_FullMethodName     = "/pb.Watermark/Watermark"
	Watermark_Status_FullMethodName        = "/pb.Watermark/Status"
	Watermark_AddDocument_FullMethodName   = "/pb.Watermark/AddDocument"
	Watermark_ServiceStatus_FullMethodName = "/pb.Watermark/ServiceStatus"
)

// WatermarkClient is the client API for Watermark service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatermarkClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDocumentReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type watermarkClient struct {
	cc grpc.ClientConnInterface
}

func NewWatermarkClient(cc grpc.ClientConnInterface) WatermarkClient {
	return &watermarkClient{cc}
}

func (c *watermarkClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReply)
	err := c.cc.Invoke(ctx, Watermark_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WatermarkReply)
	err := c.cc.Invoke(ctx, Watermark_Watermark_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, Watermark_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDocumentReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDocumentReply)
	err := c.cc.Invoke(ctx, Watermark_AddDocument_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceStatusReply)
	err := c.cc.Invoke(ctx, Watermark_ServiceStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatermarkServer is the server API for Watermark service.
// All implementations must embed UnimplementedWatermarkServer
// for forward compatibility.
type WatermarkServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Watermark(context.Context, *WatermarkRequest) (*WatermarkReply, error)
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	AddDocument(context.Context, *AddDocumentRequest) (*AddDocumentReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
	mustEmbedUnimplementedWatermarkServer()
}

// UnimplementedWatermarkServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWatermarkServer struct{}

func (UnimplementedWatermarkServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWatermarkServer) Watermark(context.Context, *WatermarkRequest) (*WatermarkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watermark not implemented")
}
func (UnimplementedWatermarkServer) Status(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedWatermarkServer) AddDocument(context.Context, *AddDocumentRequest) (*AddDocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDocument not implemented")
}
func (UnimplementedWatermarkServer) ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}
func (UnimplementedWatermarkServer) mustEmbedUnimplementedWatermarkServer() {}
func (UnimplementedWatermarkServer) testEmbeddedByValue()                   {}

// UnsafeWatermarkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatermarkServer will
// result in compilation errors.
type UnsafeWatermarkServer interface {
	mustEmbedUnimplementedWatermarkServer()
}

func RegisterWatermarkServer(s grpc.ServiceRegistrar, srv WatermarkServer) {
	// If the following call pancis, it indicates UnimplementedWatermarkServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Watermark_ServiceDesc, srv)
}

func _Watermark_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watermark_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_Watermark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatermarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Watermark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watermark_Watermark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Watermark(ctx, req.(*WatermarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watermark_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_AddDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).AddDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watermark_AddDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).AddDocument(ctx, req.(*AddDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Watermark_ServiceStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Watermark_ServiceDesc is the grpc.ServiceDesc for Watermark service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watermark_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Watermark",
	HandlerType: (*WatermarkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Watermark_Get_Handler,
		},
		{
			MethodName: "Watermark",
			Handler:    _Watermark_Watermark_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Watermark_Status_Handler,
		},
		{
			MethodName: "AddDocument",
			Handler:    _Watermark_AddDocument_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Watermark_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watermarksvc.proto",
}
