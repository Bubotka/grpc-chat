// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: protos/chat/chat.proto

package pb

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

const (
	ChatService_Chat_FullMethodName       = "/chat.ChatService/Chat"
	ChatService_ChatStream_FullMethodName = "/chat.ChatService/ChatStream"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Chat(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	ChatStream(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (ChatService_ChatStreamClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Chat(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, ChatService_Chat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ChatStream(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (ChatService_ChatStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_ChatStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ChatStreamClient interface {
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type chatServiceChatStreamClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatStreamClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Chat(context.Context, *MessageRequest) (*MessageResponse, error)
	ChatStream(*MessageRequest, ChatService_ChatStreamServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Chat(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServiceServer) ChatStream(*MessageRequest, ChatService_ChatStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatStream not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Chat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Chat(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ChatStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ChatStream(m, &chatServiceChatStreamServer{stream})
}

type ChatService_ChatStreamServer interface {
	Send(*MessageResponse) error
	grpc.ServerStream
}

type chatServiceChatStreamServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatStreamServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chat",
			Handler:    _ChatService_Chat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatStream",
			Handler:       _ChatService_ChatStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/chat/chat.proto",
}