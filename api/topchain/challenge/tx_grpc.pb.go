// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: topchain/challenge/tx.proto

package challenge

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
	Msg_UpdateParams_FullMethodName        = "/topchain.challenge.Msg/UpdateParams"
	Msg_Challenge_FullMethodName           = "/topchain.challenge.Msg/Challenge"
	Msg_SubmitProof_FullMethodName         = "/topchain.challenge.Msg/SubmitProof"
	Msg_RequestDependencies_FullMethodName = "/topchain.challenge.Msg/RequestDependencies"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	Challenge(ctx context.Context, in *MsgChallenge, opts ...grpc.CallOption) (*MsgChallengeResponse, error)
	SubmitProof(ctx context.Context, in *MsgSubmitProof, opts ...grpc.CallOption) (*MsgSubmitProofResponse, error)
	RequestDependencies(ctx context.Context, in *MsgRequestDependencies, opts ...grpc.CallOption) (*MsgRequestDependenciesResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Challenge(ctx context.Context, in *MsgChallenge, opts ...grpc.CallOption) (*MsgChallengeResponse, error) {
	out := new(MsgChallengeResponse)
	err := c.cc.Invoke(ctx, Msg_Challenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitProof(ctx context.Context, in *MsgSubmitProof, opts ...grpc.CallOption) (*MsgSubmitProofResponse, error) {
	out := new(MsgSubmitProofResponse)
	err := c.cc.Invoke(ctx, Msg_SubmitProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestDependencies(ctx context.Context, in *MsgRequestDependencies, opts ...grpc.CallOption) (*MsgRequestDependenciesResponse, error) {
	out := new(MsgRequestDependenciesResponse)
	err := c.cc.Invoke(ctx, Msg_RequestDependencies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	Challenge(context.Context, *MsgChallenge) (*MsgChallengeResponse, error)
	SubmitProof(context.Context, *MsgSubmitProof) (*MsgSubmitProofResponse, error)
	RequestDependencies(context.Context, *MsgRequestDependencies) (*MsgRequestDependenciesResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) Challenge(context.Context, *MsgChallenge) (*MsgChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Challenge not implemented")
}
func (UnimplementedMsgServer) SubmitProof(context.Context, *MsgSubmitProof) (*MsgSubmitProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProof not implemented")
}
func (UnimplementedMsgServer) RequestDependencies(context.Context, *MsgRequestDependencies) (*MsgRequestDependenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDependencies not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Challenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgChallenge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Challenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Challenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Challenge(ctx, req.(*MsgChallenge))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitProof)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SubmitProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProof(ctx, req.(*MsgSubmitProof))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestDependencies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RequestDependencies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestDependencies(ctx, req.(*MsgRequestDependencies))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "topchain.challenge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "Challenge",
			Handler:    _Msg_Challenge_Handler,
		},
		{
			MethodName: "SubmitProof",
			Handler:    _Msg_SubmitProof_Handler,
		},
		{
			MethodName: "RequestDependencies",
			Handler:    _Msg_RequestDependencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topchain/challenge/tx.proto",
}
