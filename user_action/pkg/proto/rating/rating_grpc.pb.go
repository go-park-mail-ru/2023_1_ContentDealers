// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rating

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	DeleteRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error)
	AddRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error)
	GetRatingByUser(ctx context.Context, in *RatingsOptions, opts ...grpc.CallOption) (*Ratings, error)
	GetRatingByContent(ctx context.Context, in *RatingsOptions, opts ...grpc.CallOption) (*Ratings, error)
	HasRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*HasRate, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) DeleteRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/rating.RatingService/DeleteRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) AddRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/rating.RatingService/AddRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatingByUser(ctx context.Context, in *RatingsOptions, opts ...grpc.CallOption) (*Ratings, error) {
	out := new(Ratings)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetRatingByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatingByContent(ctx context.Context, in *RatingsOptions, opts ...grpc.CallOption) (*Ratings, error) {
	out := new(Ratings)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetRatingByContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) HasRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*HasRate, error) {
	out := new(HasRate)
	err := c.cc.Invoke(ctx, "/rating.RatingService/HasRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	DeleteRating(context.Context, *Rating) (*Nothing, error)
	AddRating(context.Context, *Rating) (*Nothing, error)
	GetRatingByUser(context.Context, *RatingsOptions) (*Ratings, error)
	GetRatingByContent(context.Context, *RatingsOptions) (*Ratings, error)
	HasRating(context.Context, *Rating) (*HasRate, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) DeleteRating(context.Context, *Rating) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRating not implemented")
}
func (UnimplementedRatingServiceServer) AddRating(context.Context, *Rating) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRating not implemented")
}
func (UnimplementedRatingServiceServer) GetRatingByUser(context.Context, *RatingsOptions) (*Ratings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingByUser not implemented")
}
func (UnimplementedRatingServiceServer) GetRatingByContent(context.Context, *RatingsOptions) (*Ratings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingByContent not implemented")
}
func (UnimplementedRatingServiceServer) HasRating(context.Context, *Rating) (*HasRate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasRating not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_DeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).DeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/DeleteRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).DeleteRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_AddRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).AddRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/AddRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).AddRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatingByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RatingsOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetRatingByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingByUser(ctx, req.(*RatingsOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatingByContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RatingsOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingByContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetRatingByContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingByContent(ctx, req.(*RatingsOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_HasRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).HasRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/HasRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).HasRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteRating",
			Handler:    _RatingService_DeleteRating_Handler,
		},
		{
			MethodName: "AddRating",
			Handler:    _RatingService_AddRating_Handler,
		},
		{
			MethodName: "GetRatingByUser",
			Handler:    _RatingService_GetRatingByUser_Handler,
		},
		{
			MethodName: "GetRatingByContent",
			Handler:    _RatingService_GetRatingByContent_Handler,
		},
		{
			MethodName: "HasRating",
			Handler:    _RatingService_HasRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating.proto",
}