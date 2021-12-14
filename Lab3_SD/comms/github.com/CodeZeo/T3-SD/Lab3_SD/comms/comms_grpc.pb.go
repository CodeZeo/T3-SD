// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package comms

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

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	GetIP(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Conn, error)
	GetNumberRebelds(ctx context.Context, in *LocateCity, opts ...grpc.CallOption) (*NumberRebelds, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) GetIP(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Conn, error) {
	out := new(Conn)
	err := c.cc.Invoke(ctx, "/comms.Broker/getIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) GetNumberRebelds(ctx context.Context, in *LocateCity, opts ...grpc.CallOption) (*NumberRebelds, error) {
	out := new(NumberRebelds)
	err := c.cc.Invoke(ctx, "/comms.Broker/getNumberRebelds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility
type BrokerServer interface {
	GetIP(context.Context, *Command) (*Conn, error)
	GetNumberRebelds(context.Context, *LocateCity) (*NumberRebelds, error)
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (UnimplementedBrokerServer) GetIP(context.Context, *Command) (*Conn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIP not implemented")
}
func (UnimplementedBrokerServer) GetNumberRebelds(context.Context, *LocateCity) (*NumberRebelds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebelds not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_GetIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).GetIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Broker/getIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).GetIP(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_GetNumberRebelds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).GetNumberRebelds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Broker/getNumberRebelds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).GetNumberRebelds(ctx, req.(*LocateCity))
	}
	return interceptor(ctx, in, info, handler)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comms.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getIP",
			Handler:    _Broker_GetIP_Handler,
		},
		{
			MethodName: "getNumberRebelds",
			Handler:    _Broker_GetNumberRebelds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comms.proto",
}

// FulcrumClient is the client API for Fulcrum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulcrumClient interface {
	ReturnNumberRebelds(ctx context.Context, in *LocateCity, opts ...grpc.CallOption) (*NumberRebelds, error)
	AddCity(ctx context.Context, in *DataCity, opts ...grpc.CallOption) (*Clock, error)
	UpdateName(ctx context.Context, in *ChangeNameCity, opts ...grpc.CallOption) (*Clock, error)
	UpdateNumber(ctx context.Context, in *DataCity, opts ...grpc.CallOption) (*Clock, error)
	DeleteCity(ctx context.Context, in *LocateCity, opts ...grpc.CallOption) (*Clock, error)
	GetClock(ctx context.Context, in *Planet, opts ...grpc.CallOption) (*Clock, error)
}

type fulcrumClient struct {
	cc grpc.ClientConnInterface
}

func NewFulcrumClient(cc grpc.ClientConnInterface) FulcrumClient {
	return &fulcrumClient{cc}
}

func (c *fulcrumClient) ReturnNumberRebelds(ctx context.Context, in *LocateCity, opts ...grpc.CallOption) (*NumberRebelds, error) {
	out := new(NumberRebelds)
	err := c.cc.Invoke(ctx, "/comms.Fulcrum/returnNumberRebelds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) AddCity(ctx context.Context, in *DataCity, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/comms.Fulcrum/addCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) UpdateName(ctx context.Context, in *ChangeNameCity, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/comms.Fulcrum/updateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) UpdateNumber(ctx context.Context, in *DataCity, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/comms.Fulcrum/updateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) DeleteCity(ctx context.Context, in *LocateCity, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/comms.Fulcrum/deleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumClient) GetClock(ctx context.Context, in *Planet, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/comms.Fulcrum/getClock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulcrumServer is the server API for Fulcrum service.
// All implementations must embed UnimplementedFulcrumServer
// for forward compatibility
type FulcrumServer interface {
	ReturnNumberRebelds(context.Context, *LocateCity) (*NumberRebelds, error)
	AddCity(context.Context, *DataCity) (*Clock, error)
	UpdateName(context.Context, *ChangeNameCity) (*Clock, error)
	UpdateNumber(context.Context, *DataCity) (*Clock, error)
	DeleteCity(context.Context, *LocateCity) (*Clock, error)
	GetClock(context.Context, *Planet) (*Clock, error)
	mustEmbedUnimplementedFulcrumServer()
}

// UnimplementedFulcrumServer must be embedded to have forward compatible implementations.
type UnimplementedFulcrumServer struct {
}

func (UnimplementedFulcrumServer) ReturnNumberRebelds(context.Context, *LocateCity) (*NumberRebelds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnNumberRebelds not implemented")
}
func (UnimplementedFulcrumServer) AddCity(context.Context, *DataCity) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedFulcrumServer) UpdateName(context.Context, *ChangeNameCity) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedFulcrumServer) UpdateNumber(context.Context, *DataCity) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedFulcrumServer) DeleteCity(context.Context, *LocateCity) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedFulcrumServer) GetClock(context.Context, *Planet) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClock not implemented")
}
func (UnimplementedFulcrumServer) mustEmbedUnimplementedFulcrumServer() {}

// UnsafeFulcrumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulcrumServer will
// result in compilation errors.
type UnsafeFulcrumServer interface {
	mustEmbedUnimplementedFulcrumServer()
}

func RegisterFulcrumServer(s grpc.ServiceRegistrar, srv FulcrumServer) {
	s.RegisterService(&Fulcrum_ServiceDesc, srv)
}

func _Fulcrum_ReturnNumberRebelds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).ReturnNumberRebelds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Fulcrum/returnNumberRebelds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).ReturnNumberRebelds(ctx, req.(*LocateCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Fulcrum/addCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).AddCity(ctx, req.(*DataCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNameCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Fulcrum/updateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).UpdateName(ctx, req.(*ChangeNameCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Fulcrum/updateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).UpdateNumber(ctx, req.(*DataCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Fulcrum/deleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).DeleteCity(ctx, req.(*LocateCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fulcrum_GetClock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Planet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServer).GetClock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms.Fulcrum/getClock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServer).GetClock(ctx, req.(*Planet))
	}
	return interceptor(ctx, in, info, handler)
}

// Fulcrum_ServiceDesc is the grpc.ServiceDesc for Fulcrum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fulcrum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comms.Fulcrum",
	HandlerType: (*FulcrumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "returnNumberRebelds",
			Handler:    _Fulcrum_ReturnNumberRebelds_Handler,
		},
		{
			MethodName: "addCity",
			Handler:    _Fulcrum_AddCity_Handler,
		},
		{
			MethodName: "updateName",
			Handler:    _Fulcrum_UpdateName_Handler,
		},
		{
			MethodName: "updateNumber",
			Handler:    _Fulcrum_UpdateNumber_Handler,
		},
		{
			MethodName: "deleteCity",
			Handler:    _Fulcrum_DeleteCity_Handler,
		},
		{
			MethodName: "getClock",
			Handler:    _Fulcrum_GetClock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comms.proto",
}
