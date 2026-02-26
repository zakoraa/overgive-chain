package permissions

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"overgive-chain/x/permissions/types"
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
				{
					RpcMethod: "ListAllowed",
					Use:       "list-allowed",
					Short:     "List all allowed",
				},
				{
					RpcMethod:      "GetAllowed",
					Use:            "get-allowed [id]",
					Short:          "Gets a allowed",
					Alias:          []string{"show-allowed"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
					RpcMethod:      "CreateAllowed",
					Use:            "create-allowed [index] [address]",
					Short:          "Create a new allowed",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateAllowed",
					Use:            "update-allowed [index] [address]",
					Short:          "Update allowed",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "DeleteAllowed",
					Use:            "delete-allowed [index]",
					Short:          "Delete allowed",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
