// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: all_messages.proto

package proto

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
	DispatcherService_AcceptRequest_FullMethodName = "/DispatcherService/AcceptRequest"
)

// DispatcherServiceClient is the client API for DispatcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherServiceClient interface {
	AcceptRequest(ctx context.Context, in *DispatcherRequest, opts ...grpc.CallOption) (*DispatcherResponse, error)
}

type dispatcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherServiceClient(cc grpc.ClientConnInterface) DispatcherServiceClient {
	return &dispatcherServiceClient{cc}
}

func (c *dispatcherServiceClient) AcceptRequest(ctx context.Context, in *DispatcherRequest, opts ...grpc.CallOption) (*DispatcherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DispatcherResponse)
	err := c.cc.Invoke(ctx, DispatcherService_AcceptRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServiceServer is the server API for DispatcherService service.
// All implementations must embed UnimplementedDispatcherServiceServer
// for forward compatibility.
type DispatcherServiceServer interface {
	AcceptRequest(context.Context, *DispatcherRequest) (*DispatcherResponse, error)
	mustEmbedUnimplementedDispatcherServiceServer()
}

// UnimplementedDispatcherServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDispatcherServiceServer struct{}

func (UnimplementedDispatcherServiceServer) AcceptRequest(context.Context, *DispatcherRequest) (*DispatcherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRequest not implemented")
}
func (UnimplementedDispatcherServiceServer) mustEmbedUnimplementedDispatcherServiceServer() {}
func (UnimplementedDispatcherServiceServer) testEmbeddedByValue()                           {}

// UnsafeDispatcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherServiceServer will
// result in compilation errors.
type UnsafeDispatcherServiceServer interface {
	mustEmbedUnimplementedDispatcherServiceServer()
}

func RegisterDispatcherServiceServer(s grpc.ServiceRegistrar, srv DispatcherServiceServer) {
	// If the following call pancis, it indicates UnimplementedDispatcherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DispatcherService_ServiceDesc, srv)
}

func _DispatcherService_AcceptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServiceServer).AcceptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherService_AcceptRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServiceServer).AcceptRequest(ctx, req.(*DispatcherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherService_ServiceDesc is the grpc.ServiceDesc for DispatcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DispatcherService",
	HandlerType: (*DispatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptRequest",
			Handler:    _DispatcherService_AcceptRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "all_messages.proto",
}

const (
	FilesystemService_AcceptRequest_FullMethodName         = "/FilesystemService/AcceptRequest"
	FilesystemService_AcceptMetadataRequest_FullMethodName = "/FilesystemService/AcceptMetadataRequest"
)

// FilesystemServiceClient is the client API for FilesystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesystemServiceClient interface {
	AcceptRequest(ctx context.Context, in *FilesystemRequest, opts ...grpc.CallOption) (*FilesystemResponse, error)
	AcceptMetadataRequest(ctx context.Context, in *FilesystemMetadataRequest, opts ...grpc.CallOption) (*FilesystemMetadataResponse, error)
}

type filesystemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesystemServiceClient(cc grpc.ClientConnInterface) FilesystemServiceClient {
	return &filesystemServiceClient{cc}
}

func (c *filesystemServiceClient) AcceptRequest(ctx context.Context, in *FilesystemRequest, opts ...grpc.CallOption) (*FilesystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FilesystemResponse)
	err := c.cc.Invoke(ctx, FilesystemService_AcceptRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) AcceptMetadataRequest(ctx context.Context, in *FilesystemMetadataRequest, opts ...grpc.CallOption) (*FilesystemMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FilesystemMetadataResponse)
	err := c.cc.Invoke(ctx, FilesystemService_AcceptMetadataRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesystemServiceServer is the server API for FilesystemService service.
// All implementations must embed UnimplementedFilesystemServiceServer
// for forward compatibility.
type FilesystemServiceServer interface {
	AcceptRequest(context.Context, *FilesystemRequest) (*FilesystemResponse, error)
	AcceptMetadataRequest(context.Context, *FilesystemMetadataRequest) (*FilesystemMetadataResponse, error)
	mustEmbedUnimplementedFilesystemServiceServer()
}

// UnimplementedFilesystemServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFilesystemServiceServer struct{}

func (UnimplementedFilesystemServiceServer) AcceptRequest(context.Context, *FilesystemRequest) (*FilesystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRequest not implemented")
}
func (UnimplementedFilesystemServiceServer) AcceptMetadataRequest(context.Context, *FilesystemMetadataRequest) (*FilesystemMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptMetadataRequest not implemented")
}
func (UnimplementedFilesystemServiceServer) mustEmbedUnimplementedFilesystemServiceServer() {}
func (UnimplementedFilesystemServiceServer) testEmbeddedByValue()                           {}

// UnsafeFilesystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesystemServiceServer will
// result in compilation errors.
type UnsafeFilesystemServiceServer interface {
	mustEmbedUnimplementedFilesystemServiceServer()
}

func RegisterFilesystemServiceServer(s grpc.ServiceRegistrar, srv FilesystemServiceServer) {
	// If the following call pancis, it indicates UnimplementedFilesystemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FilesystemService_ServiceDesc, srv)
}

func _FilesystemService_AcceptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).AcceptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_AcceptRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).AcceptRequest(ctx, req.(*FilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_AcceptMetadataRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesystemMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).AcceptMetadataRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_AcceptMetadataRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).AcceptMetadataRequest(ctx, req.(*FilesystemMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FilesystemService_ServiceDesc is the grpc.ServiceDesc for FilesystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilesystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FilesystemService",
	HandlerType: (*FilesystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptRequest",
			Handler:    _FilesystemService_AcceptRequest_Handler,
		},
		{
			MethodName: "AcceptMetadataRequest",
			Handler:    _FilesystemService_AcceptMetadataRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "all_messages.proto",
}

const (
	ConsolidatorService_AcceptRequest_FullMethodName = "/ConsolidatorService/AcceptRequest"
)

// ConsolidatorServiceClient is the client API for ConsolidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsolidatorServiceClient interface {
	AcceptRequest(ctx context.Context, in *ConsolidatorRequest, opts ...grpc.CallOption) (*ConsolidatorResponse, error)
}

type consolidatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsolidatorServiceClient(cc grpc.ClientConnInterface) ConsolidatorServiceClient {
	return &consolidatorServiceClient{cc}
}

func (c *consolidatorServiceClient) AcceptRequest(ctx context.Context, in *ConsolidatorRequest, opts ...grpc.CallOption) (*ConsolidatorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConsolidatorResponse)
	err := c.cc.Invoke(ctx, ConsolidatorService_AcceptRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsolidatorServiceServer is the server API for ConsolidatorService service.
// All implementations must embed UnimplementedConsolidatorServiceServer
// for forward compatibility.
type ConsolidatorServiceServer interface {
	AcceptRequest(context.Context, *ConsolidatorRequest) (*ConsolidatorResponse, error)
	mustEmbedUnimplementedConsolidatorServiceServer()
}

// UnimplementedConsolidatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConsolidatorServiceServer struct{}

func (UnimplementedConsolidatorServiceServer) AcceptRequest(context.Context, *ConsolidatorRequest) (*ConsolidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRequest not implemented")
}
func (UnimplementedConsolidatorServiceServer) mustEmbedUnimplementedConsolidatorServiceServer() {}
func (UnimplementedConsolidatorServiceServer) testEmbeddedByValue()                             {}

// UnsafeConsolidatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsolidatorServiceServer will
// result in compilation errors.
type UnsafeConsolidatorServiceServer interface {
	mustEmbedUnimplementedConsolidatorServiceServer()
}

func RegisterConsolidatorServiceServer(s grpc.ServiceRegistrar, srv ConsolidatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedConsolidatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConsolidatorService_ServiceDesc, srv)
}

func _ConsolidatorService_AcceptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsolidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsolidatorServiceServer).AcceptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsolidatorService_AcceptRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsolidatorServiceServer).AcceptRequest(ctx, req.(*ConsolidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsolidatorService_ServiceDesc is the grpc.ServiceDesc for ConsolidatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsolidatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ConsolidatorService",
	HandlerType: (*ConsolidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcceptRequest",
			Handler:    _ConsolidatorService_AcceptRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "all_messages.proto",
}
