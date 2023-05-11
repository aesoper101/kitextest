// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: trans.proto

package api_transfrom

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

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferServiceClient interface {
	TransIn(ctx context.Context, opts ...grpc.CallOption) (TransferService_TransInClient, error)
	TransOut(ctx context.Context, opts ...grpc.CallOption) (TransferService_TransOutClient, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) TransIn(ctx context.Context, opts ...grpc.CallOption) (TransferService_TransInClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransferService_ServiceDesc.Streams[0], "/tests.TransferService/TransIn", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferServiceTransInClient{stream}
	return x, nil
}

type TransferService_TransInClient interface {
	Send(*TransferRequest) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type transferServiceTransInClient struct {
	grpc.ClientStream
}

func (x *transferServiceTransInClient) Send(m *TransferRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferServiceTransInClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transferServiceClient) TransOut(ctx context.Context, opts ...grpc.CallOption) (TransferService_TransOutClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransferService_ServiceDesc.Streams[1], "/tests.TransferService/TransOut", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferServiceTransOutClient{stream}
	return x, nil
}

type TransferService_TransOutClient interface {
	Send(*TransferRequest) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type transferServiceTransOutClient struct {
	grpc.ClientStream
}

func (x *transferServiceTransOutClient) Send(m *TransferRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferServiceTransOutClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServiceServer is the server API for TransferService service.
// All implementations must embed UnimplementedTransferServiceServer
// for forward compatibility
type TransferServiceServer interface {
	TransIn(TransferService_TransInServer) error
	TransOut(TransferService_TransOutServer) error
	mustEmbedUnimplementedTransferServiceServer()
}

// UnimplementedTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransferServiceServer struct {
}

func (UnimplementedTransferServiceServer) TransIn(TransferService_TransInServer) error {
	return status.Errorf(codes.Unimplemented, "method TransIn not implemented")
}
func (UnimplementedTransferServiceServer) TransOut(TransferService_TransOutServer) error {
	return status.Errorf(codes.Unimplemented, "method TransOut not implemented")
}
func (UnimplementedTransferServiceServer) mustEmbedUnimplementedTransferServiceServer() {}

// UnsafeTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServiceServer will
// result in compilation errors.
type UnsafeTransferServiceServer interface {
	mustEmbedUnimplementedTransferServiceServer()
}

func RegisterTransferServiceServer(s grpc.ServiceRegistrar, srv TransferServiceServer) {
	s.RegisterService(&TransferService_ServiceDesc, srv)
}

func _TransferService_TransIn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServiceServer).TransIn(&transferServiceTransInServer{stream})
}

type TransferService_TransInServer interface {
	SendAndClose(*Message) error
	Recv() (*TransferRequest, error)
	grpc.ServerStream
}

type transferServiceTransInServer struct {
	grpc.ServerStream
}

func (x *transferServiceTransInServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferServiceTransInServer) Recv() (*TransferRequest, error) {
	m := new(TransferRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TransferService_TransOut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServiceServer).TransOut(&transferServiceTransOutServer{stream})
}

type TransferService_TransOutServer interface {
	SendAndClose(*Message) error
	Recv() (*TransferRequest, error)
	grpc.ServerStream
}

type transferServiceTransOutServer struct {
	grpc.ServerStream
}

func (x *transferServiceTransOutServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferServiceTransOutServer) Recv() (*TransferRequest, error) {
	m := new(TransferRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferService_ServiceDesc is the grpc.ServiceDesc for TransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tests.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransIn",
			Handler:       _TransferService_TransIn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TransOut",
			Handler:       _TransferService_TransOut_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "trans.proto",
}