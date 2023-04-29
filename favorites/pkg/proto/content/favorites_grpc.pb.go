// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package content

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

// FavoritesContentServiceClient is the client API for FavoritesContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoritesContentServiceClient interface {
	DeleteContent(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*Nothing, error)
	AddContent(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*Nothing, error)
	GetContent(ctx context.Context, in *FavoritesOptions, opts ...grpc.CallOption) (*Favorites, error)
	HasFavContent(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*HasFav, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type favoritesContentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoritesContentServiceClient(cc grpc.ClientConnInterface) FavoritesContentServiceClient {
	return &favoritesContentServiceClient{cc}
}

func (c *favoritesContentServiceClient) DeleteContent(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/content.FavoritesContentService/DeleteContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoritesContentServiceClient) AddContent(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/content.FavoritesContentService/AddContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoritesContentServiceClient) GetContent(ctx context.Context, in *FavoritesOptions, opts ...grpc.CallOption) (*Favorites, error) {
	out := new(Favorites)
	err := c.cc.Invoke(ctx, "/content.FavoritesContentService/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoritesContentServiceClient) HasFavContent(ctx context.Context, in *Favorite, opts ...grpc.CallOption) (*HasFav, error) {
	out := new(HasFav)
	err := c.cc.Invoke(ctx, "/content.FavoritesContentService/HasFavContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoritesContentServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/content.FavoritesContentService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoritesContentServiceServer is the server API for FavoritesContentService service.
// All implementations must embed UnimplementedFavoritesContentServiceServer
// for forward compatibility
type FavoritesContentServiceServer interface {
	DeleteContent(context.Context, *Favorite) (*Nothing, error)
	AddContent(context.Context, *Favorite) (*Nothing, error)
	GetContent(context.Context, *FavoritesOptions) (*Favorites, error)
	HasFavContent(context.Context, *Favorite) (*HasFav, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedFavoritesContentServiceServer()
}

// UnimplementedFavoritesContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFavoritesContentServiceServer struct {
}

func (UnimplementedFavoritesContentServiceServer) DeleteContent(context.Context, *Favorite) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}
func (UnimplementedFavoritesContentServiceServer) AddContent(context.Context, *Favorite) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContent not implemented")
}
func (UnimplementedFavoritesContentServiceServer) GetContent(context.Context, *FavoritesOptions) (*Favorites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedFavoritesContentServiceServer) HasFavContent(context.Context, *Favorite) (*HasFav, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasFavContent not implemented")
}
func (UnimplementedFavoritesContentServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedFavoritesContentServiceServer) mustEmbedUnimplementedFavoritesContentServiceServer() {
}

// UnsafeFavoritesContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoritesContentServiceServer will
// result in compilation errors.
type UnsafeFavoritesContentServiceServer interface {
	mustEmbedUnimplementedFavoritesContentServiceServer()
}

func RegisterFavoritesContentServiceServer(s grpc.ServiceRegistrar, srv FavoritesContentServiceServer) {
	s.RegisterService(&FavoritesContentService_ServiceDesc, srv)
}

func _FavoritesContentService_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoritesContentServiceServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FavoritesContentService/DeleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoritesContentServiceServer).DeleteContent(ctx, req.(*Favorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoritesContentService_AddContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoritesContentServiceServer).AddContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FavoritesContentService/AddContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoritesContentServiceServer).AddContent(ctx, req.(*Favorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoritesContentService_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoritesOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoritesContentServiceServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FavoritesContentService/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoritesContentServiceServer).GetContent(ctx, req.(*FavoritesOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoritesContentService_HasFavContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoritesContentServiceServer).HasFavContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FavoritesContentService/HasFavContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoritesContentServiceServer).HasFavContent(ctx, req.(*Favorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoritesContentService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoritesContentServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.FavoritesContentService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoritesContentServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoritesContentService_ServiceDesc is the grpc.ServiceDesc for FavoritesContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoritesContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.FavoritesContentService",
	HandlerType: (*FavoritesContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteContent",
			Handler:    _FavoritesContentService_DeleteContent_Handler,
		},
		{
			MethodName: "AddContent",
			Handler:    _FavoritesContentService_AddContent_Handler,
		},
		{
			MethodName: "GetContent",
			Handler:    _FavoritesContentService_GetContent_Handler,
		},
		{
			MethodName: "HasFavContent",
			Handler:    _FavoritesContentService_HasFavContent_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _FavoritesContentService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorites.proto",
}
