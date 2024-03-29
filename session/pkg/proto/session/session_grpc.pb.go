// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package session

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

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	Create(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Session, error)
	Get(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error)
	Delete(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Nothing, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Create(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/session.SessionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) Get(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/session.SessionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) Delete(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.SessionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	Create(context.Context, *UserID) (*Session, error)
	Get(context.Context, *SessionID) (*Session, error)
	Delete(context.Context, *SessionID) (*Nothing, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) Create(context.Context, *UserID) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSessionServiceServer) Get(context.Context, *SessionID) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSessionServiceServer) Delete(context.Context, *SessionID) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Create(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Get(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Delete(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SessionService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SessionService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SessionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
