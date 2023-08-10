// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: banks.proto

package protobuf

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
	BankService_GetBanks_FullMethodName = "/bank_proto.BankService/GetBanks"
)

// BankServiceClient is the client API for BankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankServiceClient interface {
	GetBanks(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Banks, error)
}

type bankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankServiceClient(cc grpc.ClientConnInterface) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) GetBanks(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Banks, error) {
	out := new(Banks)
	err := c.cc.Invoke(ctx, BankService_GetBanks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServiceServer is the server API for BankService service.
// All implementations must embed UnimplementedBankServiceServer
// for forward compatibility
type BankServiceServer interface {
	GetBanks(context.Context, *Params) (*Banks, error)
	mustEmbedUnimplementedBankServiceServer()
}

// UnimplementedBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankServiceServer struct {
}

func (UnimplementedBankServiceServer) GetBanks(context.Context, *Params) (*Banks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanks not implemented")
}
func (UnimplementedBankServiceServer) mustEmbedUnimplementedBankServiceServer() {}

// UnsafeBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServiceServer will
// result in compilation errors.
type UnsafeBankServiceServer interface {
	mustEmbedUnimplementedBankServiceServer()
}

func RegisterBankServiceServer(s grpc.ServiceRegistrar, srv BankServiceServer) {
	s.RegisterService(&BankService_ServiceDesc, srv)
}

func _BankService_GetBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BankService_GetBanks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetBanks(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

// BankService_ServiceDesc is the grpc.ServiceDesc for BankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank_proto.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBanks",
			Handler:    _BankService_GetBanks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "banks.proto",
}