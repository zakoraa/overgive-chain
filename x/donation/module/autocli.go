package donation

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"overgive-chain/x/donation/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "RecordDonation",
					Use:            "record-donation [campaign-id] [amount] [currency] [xendit-reference-id] [donation-hash] [metadata-hash]",
					Short:          "Send a record-donation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "campaign_id"}, {ProtoField: "amount"}, {ProtoField: "currency"}, {ProtoField: "xendit_reference_id"}, {ProtoField: "donation_hash"}, {ProtoField: "metadata_hash"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
