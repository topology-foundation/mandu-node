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
					RpcMethod:      "Subscription",
					Use:            "subscription [id]",
					Short:          "Query subscription",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "Subscriptions",
					Use:            "subscriptions [provider]",
					Short:          "Query subscriptions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "provider"}},
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
					RpcMethod:      "SubmitProgress",
					Use:            "submit-progress [subscription_id] [previous_progress (hashes)] [obfuscated_hash]",
					Short:          "Send a submit-progress tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscription_id"}, {ProtoField: "previous_vertices_hashes"}, {ProtoField: "obfuscated_vertices_hash"}},
				},
			},
		},
	}
}
