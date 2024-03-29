// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: content.proto

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

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	AddRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error)
	DeleteRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error)
	GetFilmByContentID(ctx context.Context, in *ContentID, opts ...grpc.CallOption) (*Film, error)
	GetSeriesByContentID(ctx context.Context, in *ContentID, opts ...grpc.CallOption) (*Series, error)
	GetContentByContentIDs(ctx context.Context, in *ContentIDs, opts ...grpc.CallOption) (*ContentSeq, error)
	GetEpisodesBySeasonNum(ctx context.Context, in *ContentIDSeasonNum, opts ...grpc.CallOption) (*Episodes, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) AddRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/content.ContentService/AddRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteRating(ctx context.Context, in *Rating, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/content.ContentService/DeleteRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetFilmByContentID(ctx context.Context, in *ContentID, opts ...grpc.CallOption) (*Film, error) {
	out := new(Film)
	err := c.cc.Invoke(ctx, "/content.ContentService/GetFilmByContentID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetSeriesByContentID(ctx context.Context, in *ContentID, opts ...grpc.CallOption) (*Series, error) {
	out := new(Series)
	err := c.cc.Invoke(ctx, "/content.ContentService/GetSeriesByContentID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContentByContentIDs(ctx context.Context, in *ContentIDs, opts ...grpc.CallOption) (*ContentSeq, error) {
	out := new(ContentSeq)
	err := c.cc.Invoke(ctx, "/content.ContentService/GetContentByContentIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetEpisodesBySeasonNum(ctx context.Context, in *ContentIDSeasonNum, opts ...grpc.CallOption) (*Episodes, error) {
	out := new(Episodes)
	err := c.cc.Invoke(ctx, "/content.ContentService/GetEpisodesBySeasonNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	AddRating(context.Context, *Rating) (*Nothing, error)
	DeleteRating(context.Context, *Rating) (*Nothing, error)
	GetFilmByContentID(context.Context, *ContentID) (*Film, error)
	GetSeriesByContentID(context.Context, *ContentID) (*Series, error)
	GetContentByContentIDs(context.Context, *ContentIDs) (*ContentSeq, error)
	GetEpisodesBySeasonNum(context.Context, *ContentIDSeasonNum) (*Episodes, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) AddRating(context.Context, *Rating) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRating not implemented")
}
func (UnimplementedContentServiceServer) DeleteRating(context.Context, *Rating) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRating not implemented")
}
func (UnimplementedContentServiceServer) GetFilmByContentID(context.Context, *ContentID) (*Film, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilmByContentID not implemented")
}
func (UnimplementedContentServiceServer) GetSeriesByContentID(context.Context, *ContentID) (*Series, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeriesByContentID not implemented")
}
func (UnimplementedContentServiceServer) GetContentByContentIDs(context.Context, *ContentIDs) (*ContentSeq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentByContentIDs not implemented")
}
func (UnimplementedContentServiceServer) GetEpisodesBySeasonNum(context.Context, *ContentIDSeasonNum) (*Episodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEpisodesBySeasonNum not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_AddRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/AddRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rating)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/DeleteRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteRating(ctx, req.(*Rating))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetFilmByContentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetFilmByContentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/GetFilmByContentID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetFilmByContentID(ctx, req.(*ContentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetSeriesByContentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetSeriesByContentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/GetSeriesByContentID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetSeriesByContentID(ctx, req.(*ContentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetContentByContentIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContentByContentIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/GetContentByContentIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContentByContentIDs(ctx, req.(*ContentIDs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetEpisodesBySeasonNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentIDSeasonNum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetEpisodesBySeasonNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentService/GetEpisodesBySeasonNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetEpisodesBySeasonNum(ctx, req.(*ContentIDSeasonNum))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRating",
			Handler:    _ContentService_AddRating_Handler,
		},
		{
			MethodName: "DeleteRating",
			Handler:    _ContentService_DeleteRating_Handler,
		},
		{
			MethodName: "GetFilmByContentID",
			Handler:    _ContentService_GetFilmByContentID_Handler,
		},
		{
			MethodName: "GetSeriesByContentID",
			Handler:    _ContentService_GetSeriesByContentID_Handler,
		},
		{
			MethodName: "GetContentByContentIDs",
			Handler:    _ContentService_GetContentByContentIDs_Handler,
		},
		{
			MethodName: "GetEpisodesBySeasonNum",
			Handler:    _ContentService_GetEpisodesBySeasonNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}
