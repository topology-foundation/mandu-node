package challenge

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "mandu/api/mandu/challenge"
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
				// Is this the proof for a single vertex hash? Are they submitted independently? Or is this the hash of the vertices in the epoch?
				// If it is the latter, we don't need this method, and the QueryProofs method can simply replace this one.
				{
					RpcMethod:      "Proof",
					Use:            "proof [challenge_id] [hash]",
					Short:          "Query proof",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "challenge_id"}, {ProtoField: "hash"}},
				},
				{
					RpcMethod:      "Proofs",
					Use:            "proofs [challenge_id]",
					Short:          "Query proofs",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "challenge_id"}},
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
					RpcMethod:      "Challenge",
					Use:            "challenge [subscriber_id] [epoch]",
					Short:          "Send a challenge tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subscriber_id"}, {ProtoField: "epoch"}},
				},
				{
					RpcMethod:      "SubmitProof",
					Use:            "submit-proof [challenge_id] [vertices]",
					Short:          "Submit a submit-proof tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "challenge_id"}, {ProtoField: "vertices"}},
				},
				{
					RpcMethod:      "SettleChallenge",
					Use:            "settle-challenge [challenge_id]",
					Short:          "Send a settle-challenge tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "challenge_id"}},
				},
			},
		},
	}
}
