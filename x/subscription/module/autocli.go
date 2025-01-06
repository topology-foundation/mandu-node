package subscription

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "mandu/api/mandu/subscription"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "SubscriptionRequest",
					Use:            "sub-req [id]",
					Short:          "Query subscription request",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "SubscriptionRequestStatus",
					Use:            "sub-req-status [id]",
					Short:          "Query subscription request status",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "SubscriptionRequests",
					Use:            "sub-reqs [requester]",
					Short:          "Query deals",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "requester"}},
				},
				{
					RpcMethod:      "Subscription",
					Use:            "subscription [id]",
					Short:          "Query subscription",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "Subscriptions",
					Use:            "subscriptions [subscriber]",
					Short:          "Query subscriptions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscriber"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateSubscriptionRequest",
					Use:            "create-sub-req [cro_id] [amount] [start_block] [end_block] [initial_frontier]",
					Short:          "Send a create-sub-req tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "start_block"}, {ProtoField: "epoch_size"}, {ProtoField: "duration"}, {ProtoField: "drp_ids"}, {ProtoField: "writers"}, {ProtoField: "initial_frontier"}},
				},
				{
					RpcMethod:      "CancelSubscriptionRequest",
					Use:            "cancel-sub-req [subscription_request_id]",
					Short:          "Send a cancel-sub-req tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_request_id"}},
				},
				{
					RpcMethod:      "UpdateSubscriptionRequest",
					Use:            "update-sub-req [subscription_request_id] [amount] [start_block] [end_block]",
					Short:          "Send a update-sub-req tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_request_id"}, {ProtoField: "amount"}, {ProtoField: "start_block"}, {ProtoField: "duration"}, {ProtoField: "writers"}},
				},
				{
					RpcMethod:      "IncrementSubscriptionRequestAmount",
					Use:            "increment-sub-req-amount [subscription_request_id] [amount]",
					Short:          "Send a increment-sub-req-amount tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_request_id"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "JoinSubscriptionRequest",
					Use:            "join-sub-req [subscription_request_id]",
					Short:          "Send a join-sub-req tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_request_id"}},
				},
				{
					RpcMethod:      "LeaveSubscriptionRequest",
					Use:            "leave-sub-req [subscription_request_id]",
					Short:          "Send a leave-sub-req tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_request_id"}},
				},
				{
					RpcMethod:      "SubmitProgress",
					Use:            "submit-progress [subscription_id] [previous_progress (hashes)] [obfuscated_hash]",
					Short:          "Send a submit-progress tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_id"}, {ProtoField: "previous_vertices_hashes"}, {ProtoField: "obfuscated_vertices_hash"}},
				},
			},
		},
	}
}
