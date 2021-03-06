// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: example_api_server.proto

package example_api_server_rpc_api_bindings

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExampleAPIServerServiceClient is the client API for ExampleAPIServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleAPIServerServiceClient interface {
	//Used to check service availability
	IsAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//Add a person ID into the book reads store
	AddPerson(ctx context.Context, in *AddPersonArgs, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//Get a person books read value by person ID
	GetPerson(ctx context.Context, in *GetPersonArgs, opts ...grpc.CallOption) (*GetPersonResponse, error)
	//Increment in 1 the amount of person books read by person ID
	IncrementBooksRead(ctx context.Context, in *IncrementBooksReadArgs, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type exampleAPIServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleAPIServerServiceClient(cc grpc.ClientConnInterface) ExampleAPIServerServiceClient {
	return &exampleAPIServerServiceClient{cc}
}

func (c *exampleAPIServerServiceClient) IsAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/example_api_server_api.ExampleAPIServerService/IsAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAPIServerServiceClient) AddPerson(ctx context.Context, in *AddPersonArgs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/example_api_server_api.ExampleAPIServerService/AddPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAPIServerServiceClient) GetPerson(ctx context.Context, in *GetPersonArgs, opts ...grpc.CallOption) (*GetPersonResponse, error) {
	out := new(GetPersonResponse)
	err := c.cc.Invoke(ctx, "/example_api_server_api.ExampleAPIServerService/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleAPIServerServiceClient) IncrementBooksRead(ctx context.Context, in *IncrementBooksReadArgs, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/example_api_server_api.ExampleAPIServerService/IncrementBooksRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleAPIServerServiceServer is the server API for ExampleAPIServerService service.
// All implementations must embed UnimplementedExampleAPIServerServiceServer
// for forward compatibility
type ExampleAPIServerServiceServer interface {
	//Used to check service availability
	IsAvailable(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	//Add a person ID into the book reads store
	AddPerson(context.Context, *AddPersonArgs) (*emptypb.Empty, error)
	//Get a person books read value by person ID
	GetPerson(context.Context, *GetPersonArgs) (*GetPersonResponse, error)
	//Increment in 1 the amount of person books read by person ID
	IncrementBooksRead(context.Context, *IncrementBooksReadArgs) (*emptypb.Empty, error)
	mustEmbedUnimplementedExampleAPIServerServiceServer()
}

// UnimplementedExampleAPIServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleAPIServerServiceServer struct {
}

func (UnimplementedExampleAPIServerServiceServer) IsAvailable(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvailable not implemented")
}
func (UnimplementedExampleAPIServerServiceServer) AddPerson(context.Context, *AddPersonArgs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (UnimplementedExampleAPIServerServiceServer) GetPerson(context.Context, *GetPersonArgs) (*GetPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedExampleAPIServerServiceServer) IncrementBooksRead(context.Context, *IncrementBooksReadArgs) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementBooksRead not implemented")
}
func (UnimplementedExampleAPIServerServiceServer) mustEmbedUnimplementedExampleAPIServerServiceServer() {
}

// UnsafeExampleAPIServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleAPIServerServiceServer will
// result in compilation errors.
type UnsafeExampleAPIServerServiceServer interface {
	mustEmbedUnimplementedExampleAPIServerServiceServer()
}

func RegisterExampleAPIServerServiceServer(s grpc.ServiceRegistrar, srv ExampleAPIServerServiceServer) {
	s.RegisterService(&ExampleAPIServerService_ServiceDesc, srv)
}

func _ExampleAPIServerService_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServerServiceServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example_api_server_api.ExampleAPIServerService/IsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServerServiceServer).IsAvailable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleAPIServerService_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPersonArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServerServiceServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example_api_server_api.ExampleAPIServerService/AddPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServerServiceServer).AddPerson(ctx, req.(*AddPersonArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleAPIServerService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServerServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example_api_server_api.ExampleAPIServerService/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServerServiceServer).GetPerson(ctx, req.(*GetPersonArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleAPIServerService_IncrementBooksRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementBooksReadArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleAPIServerServiceServer).IncrementBooksRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example_api_server_api.ExampleAPIServerService/IncrementBooksRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleAPIServerServiceServer).IncrementBooksRead(ctx, req.(*IncrementBooksReadArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleAPIServerService_ServiceDesc is the grpc.ServiceDesc for ExampleAPIServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleAPIServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example_api_server_api.ExampleAPIServerService",
	HandlerType: (*ExampleAPIServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvailable",
			Handler:    _ExampleAPIServerService_IsAvailable_Handler,
		},
		{
			MethodName: "AddPerson",
			Handler:    _ExampleAPIServerService_AddPerson_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _ExampleAPIServerService_GetPerson_Handler,
		},
		{
			MethodName: "IncrementBooksRead",
			Handler:    _ExampleAPIServerService_IncrementBooksRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example_api_server.proto",
}
