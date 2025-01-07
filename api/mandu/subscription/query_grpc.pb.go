// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: mandu/subscription/query.proto

package subscription

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Query_Params_FullMethodName                    = "/mandu.subscription.Query/Params"
	Query_SubscriptionRequest_FullMethodName       = "/mandu.subscription.Query/SubscriptionRequest"
	Query_SubscriptionRequestStatus_FullMethodName = "/mandu.subscription.Query/SubscriptionRequestStatus"
	Query_SubscriptionRequests_FullMethodName      = "/mandu.subscription.Query/SubscriptionRequests"
	Query_Subscription_FullMethodName              = "/mandu.subscription.Query/Subscription"
	Query_Subscriptions_FullMethodName             = "/mandu.subscription.Query/Subscriptions"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	SubscriptionRequest(ctx context.Context, in *QuerySubscriptionRequestRequest, opts ...grpc.CallOption) (*QuerySubscriptionRequestResponse, error)
	SubscriptionRequestStatus(ctx context.Context, in *QuerySubscriptionRequestStatusRequest, opts ...grpc.CallOption) (*QuerySubscriptionRequestStatusResponse, error)
	SubscriptionRequests(ctx context.Context, in *QuerySubscriptionRequestsRequest, opts ...grpc.CallOption) (*QuerySubscriptionRequestsResponse, error)
	Subscription(ctx context.Context, in *QuerySubscriptionRequest, opts ...grpc.CallOption) (*QuerySubscriptionResponse, error)
	Subscriptions(ctx context.Context, in *QuerySubscriptionsRequest, opts ...grpc.CallOption) (*QuerySubscriptionsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SubscriptionRequest(ctx context.Context, in *QuerySubscriptionRequestRequest, opts ...grpc.CallOption) (*QuerySubscriptionRequestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySubscriptionRequestResponse)
	err := c.cc.Invoke(ctx, Query_SubscriptionRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SubscriptionRequestStatus(ctx context.Context, in *QuerySubscriptionRequestStatusRequest, opts ...grpc.CallOption) (*QuerySubscriptionRequestStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySubscriptionRequestStatusResponse)
	err := c.cc.Invoke(ctx, Query_SubscriptionRequestStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SubscriptionRequests(ctx context.Context, in *QuerySubscriptionRequestsRequest, opts ...grpc.CallOption) (*QuerySubscriptionRequestsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySubscriptionRequestsResponse)
	err := c.cc.Invoke(ctx, Query_SubscriptionRequests_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Subscription(ctx context.Context, in *QuerySubscriptionRequest, opts ...grpc.CallOption) (*QuerySubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySubscriptionResponse)
	err := c.cc.Invoke(ctx, Query_Subscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Subscriptions(ctx context.Context, in *QuerySubscriptionsRequest, opts ...grpc.CallOption) (*QuerySubscriptionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySubscriptionsResponse)
	err := c.cc.Invoke(ctx, Query_Subscriptions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query defines the gRPC querier service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	SubscriptionRequest(context.Context, *QuerySubscriptionRequestRequest) (*QuerySubscriptionRequestResponse, error)
	SubscriptionRequestStatus(context.Context, *QuerySubscriptionRequestStatusRequest) (*QuerySubscriptionRequestStatusResponse, error)
	SubscriptionRequests(context.Context, *QuerySubscriptionRequestsRequest) (*QuerySubscriptionRequestsResponse, error)
	Subscription(context.Context, *QuerySubscriptionRequest) (*QuerySubscriptionResponse, error)
	Subscriptions(context.Context, *QuerySubscriptionsRequest) (*QuerySubscriptionsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) SubscriptionRequest(context.Context, *QuerySubscriptionRequestRequest) (*QuerySubscriptionRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionRequest not implemented")
}
func (UnimplementedQueryServer) SubscriptionRequestStatus(context.Context, *QuerySubscriptionRequestStatusRequest) (*QuerySubscriptionRequestStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionRequestStatus not implemented")
}
func (UnimplementedQueryServer) SubscriptionRequests(context.Context, *QuerySubscriptionRequestsRequest) (*QuerySubscriptionRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscriptionRequests not implemented")
}
func (UnimplementedQueryServer) Subscription(context.Context, *QuerySubscriptionRequest) (*QuerySubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscription not implemented")
}
func (UnimplementedQueryServer) Subscriptions(context.Context, *QuerySubscriptionsRequest) (*QuerySubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscriptions not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SubscriptionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscriptionRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SubscriptionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SubscriptionRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SubscriptionRequest(ctx, req.(*QuerySubscriptionRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SubscriptionRequestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscriptionRequestStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SubscriptionRequestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SubscriptionRequestStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SubscriptionRequestStatus(ctx, req.(*QuerySubscriptionRequestStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SubscriptionRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscriptionRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SubscriptionRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SubscriptionRequests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SubscriptionRequests(ctx, req.(*QuerySubscriptionRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Subscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Subscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Subscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Subscription(ctx, req.(*QuerySubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Subscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Subscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Subscriptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Subscriptions(ctx, req.(*QuerySubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mandu.subscription.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SubscriptionRequest",
			Handler:    _Query_SubscriptionRequest_Handler,
		},
		{
			MethodName: "SubscriptionRequestStatus",
			Handler:    _Query_SubscriptionRequestStatus_Handler,
		},
		{
			MethodName: "SubscriptionRequests",
			Handler:    _Query_SubscriptionRequests_Handler,
		},
		{
			MethodName: "Subscription",
			Handler:    _Query_Subscription_Handler,
		},
		{
			MethodName: "Subscriptions",
			Handler:    _Query_Subscriptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mandu/subscription/query.proto",
}