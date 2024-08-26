// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: friends.proto

package friends_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FriendsService_SetFriends_FullMethodName       = "/FriendsService/SetFriends"
	FriendsService_GetPeerFollow_FullMethodName    = "/FriendsService/GetPeerFollow"
	FriendsService_GetWhoFollowPeer_FullMethodName = "/FriendsService/GetWhoFollowPeer"
	FriendsService_RemoveSubscribe_FullMethodName  = "/FriendsService/RemoveSubscribe"
	FriendsService_GetInvitePeer_FullMethodName    = "/FriendsService/GetInvitePeer"
)

// FriendsServiceClient is the client API for FriendsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for friends
type FriendsServiceClient interface {
	// Add friends method
	SetFriends(ctx context.Context, in *SetFriendsIn, opts ...grpc.CallOption) (*SetFriendsOut, error)
	GetPeerFollow(ctx context.Context, in *GetPeerFollowIn, opts ...grpc.CallOption) (*GetPeerFollowOut, error)
	GetWhoFollowPeer(ctx context.Context, in *GetWhoFollowPeerIn, opts ...grpc.CallOption) (*GetWhoFollowPeerOut, error)
	RemoveSubscribe(ctx context.Context, in *RemoveSubscribeIn, opts ...grpc.CallOption) (*RemoveSubscribeOut, error)
	GetInvitePeer(ctx context.Context, in *GetInvitePeerIn, opts ...grpc.CallOption) (*GetInvitePeerOut, error)
}

type friendsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendsServiceClient(cc grpc.ClientConnInterface) FriendsServiceClient {
	return &friendsServiceClient{cc}
}

func (c *friendsServiceClient) SetFriends(ctx context.Context, in *SetFriendsIn, opts ...grpc.CallOption) (*SetFriendsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetFriendsOut)
	err := c.cc.Invoke(ctx, FriendsService_SetFriends_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) GetPeerFollow(ctx context.Context, in *GetPeerFollowIn, opts ...grpc.CallOption) (*GetPeerFollowOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPeerFollowOut)
	err := c.cc.Invoke(ctx, FriendsService_GetPeerFollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) GetWhoFollowPeer(ctx context.Context, in *GetWhoFollowPeerIn, opts ...grpc.CallOption) (*GetWhoFollowPeerOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWhoFollowPeerOut)
	err := c.cc.Invoke(ctx, FriendsService_GetWhoFollowPeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) RemoveSubscribe(ctx context.Context, in *RemoveSubscribeIn, opts ...grpc.CallOption) (*RemoveSubscribeOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveSubscribeOut)
	err := c.cc.Invoke(ctx, FriendsService_RemoveSubscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsServiceClient) GetInvitePeer(ctx context.Context, in *GetInvitePeerIn, opts ...grpc.CallOption) (*GetInvitePeerOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvitePeerOut)
	err := c.cc.Invoke(ctx, FriendsService_GetInvitePeer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendsServiceServer is the server API for FriendsService service.
// All implementations must embed UnimplementedFriendsServiceServer
// for forward compatibility
//
// Service for friends
type FriendsServiceServer interface {
	// Add friends method
	SetFriends(context.Context, *SetFriendsIn) (*SetFriendsOut, error)
	GetPeerFollow(context.Context, *GetPeerFollowIn) (*GetPeerFollowOut, error)
	GetWhoFollowPeer(context.Context, *GetWhoFollowPeerIn) (*GetWhoFollowPeerOut, error)
	RemoveSubscribe(context.Context, *RemoveSubscribeIn) (*RemoveSubscribeOut, error)
	GetInvitePeer(context.Context, *GetInvitePeerIn) (*GetInvitePeerOut, error)
	mustEmbedUnimplementedFriendsServiceServer()
}

// UnimplementedFriendsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendsServiceServer struct {
}

func (UnimplementedFriendsServiceServer) SetFriends(context.Context, *SetFriendsIn) (*SetFriendsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriends not implemented")
}
func (UnimplementedFriendsServiceServer) GetPeerFollow(context.Context, *GetPeerFollowIn) (*GetPeerFollowOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerFollow not implemented")
}
func (UnimplementedFriendsServiceServer) GetWhoFollowPeer(context.Context, *GetWhoFollowPeerIn) (*GetWhoFollowPeerOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhoFollowPeer not implemented")
}
func (UnimplementedFriendsServiceServer) RemoveSubscribe(context.Context, *RemoveSubscribeIn) (*RemoveSubscribeOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubscribe not implemented")
}
func (UnimplementedFriendsServiceServer) GetInvitePeer(context.Context, *GetInvitePeerIn) (*GetInvitePeerOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvitePeer not implemented")
}
func (UnimplementedFriendsServiceServer) mustEmbedUnimplementedFriendsServiceServer() {}

// UnsafeFriendsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendsServiceServer will
// result in compilation errors.
type UnsafeFriendsServiceServer interface {
	mustEmbedUnimplementedFriendsServiceServer()
}

func RegisterFriendsServiceServer(s grpc.ServiceRegistrar, srv FriendsServiceServer) {
	s.RegisterService(&FriendsService_ServiceDesc, srv)
}

func _FriendsService_SetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendsIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).SetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendsService_SetFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).SetFriends(ctx, req.(*SetFriendsIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_GetPeerFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeerFollowIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).GetPeerFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendsService_GetPeerFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).GetPeerFollow(ctx, req.(*GetPeerFollowIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_GetWhoFollowPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWhoFollowPeerIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).GetWhoFollowPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendsService_GetWhoFollowPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).GetWhoFollowPeer(ctx, req.(*GetWhoFollowPeerIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_RemoveSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubscribeIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).RemoveSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendsService_RemoveSubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).RemoveSubscribe(ctx, req.(*RemoveSubscribeIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendsService_GetInvitePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvitePeerIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServiceServer).GetInvitePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FriendsService_GetInvitePeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServiceServer).GetInvitePeer(ctx, req.(*GetInvitePeerIn))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendsService_ServiceDesc is the grpc.ServiceDesc for FriendsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FriendsService",
	HandlerType: (*FriendsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetFriends",
			Handler:    _FriendsService_SetFriends_Handler,
		},
		{
			MethodName: "GetPeerFollow",
			Handler:    _FriendsService_GetPeerFollow_Handler,
		},
		{
			MethodName: "GetWhoFollowPeer",
			Handler:    _FriendsService_GetWhoFollowPeer_Handler,
		},
		{
			MethodName: "RemoveSubscribe",
			Handler:    _FriendsService_RemoveSubscribe_Handler,
		},
		{
			MethodName: "GetInvitePeer",
			Handler:    _FriendsService_GetInvitePeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friends.proto",
}
