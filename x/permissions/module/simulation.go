package permissions

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	permissionssimulation "overgive-chain/x/permissions/simulation"
	"overgive-chain/x/permissions/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	permissionsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AllowedMap: []types.Allowed{{
			Address: "0",
		}, {
			Address: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&permissionsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateAllowed          = "op_weight_msg_permissions"
		defaultWeightMsgCreateAllowed int = 100
	)

	var weightMsgCreateAllowed int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateAllowed, &weightMsgCreateAllowed, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAllowed = defaultWeightMsgCreateAllowed
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAllowed,
		permissionssimulation.SimulateMsgCreateAllowed(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateAllowed          = "op_weight_msg_permissions"
		defaultWeightMsgUpdateAllowed int = 100
	)

	var weightMsgUpdateAllowed int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateAllowed, &weightMsgUpdateAllowed, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAllowed = defaultWeightMsgUpdateAllowed
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAllowed,
		permissionssimulation.SimulateMsgUpdateAllowed(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteAllowed          = "op_weight_msg_permissions"
		defaultWeightMsgDeleteAllowed int = 100
	)

	var weightMsgDeleteAllowed int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteAllowed, &weightMsgDeleteAllowed, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAllowed = defaultWeightMsgDeleteAllowed
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAllowed,
		permissionssimulation.SimulateMsgDeleteAllowed(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgAddPermissions          = "op_weight_msg_permissions"
		defaultWeightMsgAddPermissions int = 100
	)

	var weightMsgAddPermissions int
	simState.AppParams.GetOrGenerate(opWeightMsgAddPermissions, &weightMsgAddPermissions, nil,
		func(_ *rand.Rand) {
			weightMsgAddPermissions = defaultWeightMsgAddPermissions
		},
	)

	const (
		opWeightMsgRemovePermissions          = "op_weight_msg_permissions"
		defaultWeightMsgRemovePermissions int = 100
	)

	var weightMsgRemovePermissions int
	simState.AppParams.GetOrGenerate(opWeightMsgRemovePermissions, &weightMsgRemovePermissions, nil,
		func(_ *rand.Rand) {
			weightMsgRemovePermissions = defaultWeightMsgRemovePermissions
		},
	)

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
