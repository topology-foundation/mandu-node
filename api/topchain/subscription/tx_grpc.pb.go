// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: topchain/subscription/tx.proto

package subscription

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
	Msg_UpdateParams_FullMethodName        = "/topchain.subscription.Msg/UpdateParams"
	Msg_CreateDeal_FullMethodName          = "/topchain.subscription.Msg/CreateDeal"
	Msg_CancelDeal_FullMethodName          = "/topchain.subscription.Msg/CancelDeal"
	Msg_UpdateDeal_FullMethodName          = "/topchain.subscription.Msg/UpdateDeal"
	Msg_IncrementDealAmount_FullMethodName = "/topchain.subscription.Msg/IncrementDealAmount"
	Msg_JoinDeal_FullMethodName            = "/topchain.subscription.Msg/JoinDeal"
	Msg_LeaveDeal_FullMethodName           = "/topchain.subscription.Msg/LeaveDeal"
	Msg_Challenge_FullMethodName           = "/topchain.subscription.Msg/Challenge"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateDeal(ctx context.Context, in *MsgCreateDeal, opts ...grpc.CallOption) (*MsgCreateDealResponse, error)
	CancelDeal(ctx context.Context, in *MsgCancelDeal, opts ...grpc.CallOption) (*MsgCancelDealResponse, error)
	UpdateDeal(ctx context.Context, in *MsgUpdateDeal, opts ...grpc.CallOption) (*MsgUpdateDealResponse, error)
	IncrementDealAmount(ctx context.Context, in *MsgIncrementDealAmount, opts ...grpc.CallOption) (*MsgIncrementDealAmountResponse, error)
	JoinDeal(ctx context.Context, in *MsgJoinDeal, opts ...grpc.CallOption) (*MsgJoinDealResponse, error)
	LeaveDeal(ctx context.Context, in *MsgLeaveDeal, opts ...grpc.CallOption) (*MsgLeaveDealResponse, error)
	Challenge(ctx context.Context, in *MsgChallenge, opts ...grpc.CallOption) (*MsgChallengeResponse, error)
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

func (c *msgClient) CreateDeal(ctx context.Context, in *MsgCreateDeal, opts ...grpc.CallOption) (*MsgCreateDealResponse, error) {
	out := new(MsgCreateDealResponse)
	err := c.cc.Invoke(ctx, Msg_CreateDeal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelDeal(ctx context.Context, in *MsgCancelDeal, opts ...grpc.CallOption) (*MsgCancelDealResponse, error) {
	out := new(MsgCancelDealResponse)
	err := c.cc.Invoke(ctx, Msg_CancelDeal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDeal(ctx context.Context, in *MsgUpdateDeal, opts ...grpc.CallOption) (*MsgUpdateDealResponse, error) {
	out := new(MsgUpdateDealResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateDeal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) IncrementDealAmount(ctx context.Context, in *MsgIncrementDealAmount, opts ...grpc.CallOption) (*MsgIncrementDealAmountResponse, error) {
	out := new(MsgIncrementDealAmountResponse)
	err := c.cc.Invoke(ctx, Msg_IncrementDealAmount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinDeal(ctx context.Context, in *MsgJoinDeal, opts ...grpc.CallOption) (*MsgJoinDealResponse, error) {
	out := new(MsgJoinDealResponse)
	err := c.cc.Invoke(ctx, Msg_JoinDeal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LeaveDeal(ctx context.Context, in *MsgLeaveDeal, opts ...grpc.CallOption) (*MsgLeaveDealResponse, error) {
	out := new(MsgLeaveDealResponse)
	err := c.cc.Invoke(ctx, Msg_LeaveDeal_FullMethodName, in, out, opts...)
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

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateDeal(context.Context, *MsgCreateDeal) (*MsgCreateDealResponse, error)
	CancelDeal(context.Context, *MsgCancelDeal) (*MsgCancelDealResponse, error)
	UpdateDeal(context.Context, *MsgUpdateDeal) (*MsgUpdateDealResponse, error)
	IncrementDealAmount(context.Context, *MsgIncrementDealAmount) (*MsgIncrementDealAmountResponse, error)
	JoinDeal(context.Context, *MsgJoinDeal) (*MsgJoinDealResponse, error)
	LeaveDeal(context.Context, *MsgLeaveDeal) (*MsgLeaveDealResponse, error)
	Challenge(context.Context, *MsgChallenge) (*MsgChallengeResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateDeal(context.Context, *MsgCreateDeal) (*MsgCreateDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeal not implemented")
}
func (UnimplementedMsgServer) CancelDeal(context.Context, *MsgCancelDeal) (*MsgCancelDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDeal not implemented")
}
func (UnimplementedMsgServer) UpdateDeal(context.Context, *MsgUpdateDeal) (*MsgUpdateDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeal not implemented")
}
func (UnimplementedMsgServer) IncrementDealAmount(context.Context, *MsgIncrementDealAmount) (*MsgIncrementDealAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementDealAmount not implemented")
}
func (UnimplementedMsgServer) JoinDeal(context.Context, *MsgJoinDeal) (*MsgJoinDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinDeal not implemented")
}
func (UnimplementedMsgServer) LeaveDeal(context.Context, *MsgLeaveDeal) (*MsgLeaveDealResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveDeal not implemented")
}
func (UnimplementedMsgServer) Challenge(context.Context, *MsgChallenge) (*MsgChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Challenge not implemented")
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

func _Msg_CreateDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDeal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDeal(ctx, req.(*MsgCreateDeal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelDeal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelDeal(ctx, req.(*MsgCancelDeal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDeal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDeal(ctx, req.(*MsgUpdateDeal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_IncrementDealAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIncrementDealAmount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IncrementDealAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_IncrementDealAmount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IncrementDealAmount(ctx, req.(*MsgIncrementDealAmount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinDeal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_JoinDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinDeal(ctx, req.(*MsgJoinDeal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LeaveDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLeaveDeal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LeaveDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_LeaveDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LeaveDeal(ctx, req.(*MsgLeaveDeal))
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

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "topchain.subscription.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateDeal",
			Handler:    _Msg_CreateDeal_Handler,
		},
		{
			MethodName: "CancelDeal",
			Handler:    _Msg_CancelDeal_Handler,
		},
		{
			MethodName: "UpdateDeal",
			Handler:    _Msg_UpdateDeal_Handler,
		},
		{
			MethodName: "IncrementDealAmount",
			Handler:    _Msg_IncrementDealAmount_Handler,
		},
		{
			MethodName: "JoinDeal",
			Handler:    _Msg_JoinDeal_Handler,
		},
		{
			MethodName: "LeaveDeal",
			Handler:    _Msg_LeaveDeal_Handler,
		},
		{
			MethodName: "Challenge",
			Handler:    _Msg_Challenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topchain/subscription/tx.proto",
}
