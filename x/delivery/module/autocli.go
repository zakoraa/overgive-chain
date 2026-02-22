package delivery

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"overgive-chain/x/delivery/types"
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
					RpcMethod:      "CreateDelivery",
					Use:            "create-delivery [id] [campaign-id] [title] [note-hash] [created-at] [created-by]",
					Short:          "Send a create-delivery tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "campaign_id"}, {ProtoField: "title"}, {ProtoField: "note_hash"}, {ProtoField: "created_at"}, {ProtoField: "created_by"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
