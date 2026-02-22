package delivery

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	deliverysimulation "overgive-chain/x/delivery/simulation"
	"overgive-chain/x/delivery/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	deliveryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&deliveryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateDelivery          = "op_weight_msg_delivery"
		defaultWeightMsgCreateDelivery int = 100
	)

	var weightMsgCreateDelivery int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDelivery, &weightMsgCreateDelivery, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDelivery = defaultWeightMsgCreateDelivery
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDelivery,
		deliverysimulation.SimulateMsgCreateDelivery(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
