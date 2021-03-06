// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ChittyChat

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

// ChittyChatServiceClient is the client API for ChittyChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Response, error)
	GetBroadcast(ctx context.Context, in *GetBroadcastRequest, opts ...grpc.CallOption) (*Response, error)
	JoinChat(ctx context.Context, in *JoinChatRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	LeaveChat(ctx context.Context, in *LeaveChatRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ChittyChat.ChittyChatService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) GetBroadcast(ctx context.Context, in *GetBroadcastRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ChittyChat.ChittyChatService/GetBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) JoinChat(ctx context.Context, in *JoinChatRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/ChittyChat.ChittyChatService/JoinChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) LeaveChat(ctx context.Context, in *LeaveChatRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, "/ChittyChat.ChittyChatService/LeaveChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittyChatServiceServer is the server API for ChittyChatService service.
// All implementations must embed UnimplementedChittyChatServiceServer
// for forward compatibility
type ChittyChatServiceServer interface {
	Publish(context.Context, *PublishRequest) (*Response, error)
	GetBroadcast(context.Context, *GetBroadcastRequest) (*Response, error)
	JoinChat(context.Context, *JoinChatRequest) (*JoinResponse, error)
	LeaveChat(context.Context, *LeaveChatRequest) (*LeaveResponse, error)
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServiceServer struct {
}

func (UnimplementedChittyChatServiceServer) Publish(context.Context, *PublishRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittyChatServiceServer) GetBroadcast(context.Context, *GetBroadcastRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBroadcast not implemented")
}
func (UnimplementedChittyChatServiceServer) JoinChat(context.Context, *JoinChatRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChat not implemented")
}
func (UnimplementedChittyChatServiceServer) LeaveChat(context.Context, *LeaveChatRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChat not implemented")
}
func (UnimplementedChittyChatServiceServer) mustEmbedUnimplementedChittyChatServiceServer() {}

// UnsafeChittyChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServiceServer will
// result in compilation errors.
type UnsafeChittyChatServiceServer interface {
	mustEmbedUnimplementedChittyChatServiceServer()
}

func RegisterChittyChatServiceServer(s grpc.ServiceRegistrar, srv ChittyChatServiceServer) {
	s.RegisterService(&ChittyChatService_ServiceDesc, srv)
}

func _ChittyChatService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChat.ChittyChatService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_GetBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).GetBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChat.ChittyChatService/GetBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).GetBroadcast(ctx, req.(*GetBroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_JoinChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).JoinChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChat.ChittyChatService/JoinChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).JoinChat(ctx, req.(*JoinChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_LeaveChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).LeaveChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChat.ChittyChatService/LeaveChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).LeaveChat(ctx, req.(*LeaveChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittyChatService_ServiceDesc is the grpc.ServiceDesc for ChittyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittyChat.ChittyChatService",
	HandlerType: (*ChittyChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _ChittyChatService_Publish_Handler,
		},
		{
			MethodName: "GetBroadcast",
			Handler:    _ChittyChatService_GetBroadcast_Handler,
		},
		{
			MethodName: "JoinChat",
			Handler:    _ChittyChatService_JoinChat_Handler,
		},
		{
			MethodName: "LeaveChat",
			Handler:    _ChittyChatService_LeaveChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ChittyChat/ChittyChat.proto",
}
