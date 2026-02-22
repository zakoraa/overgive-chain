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
		opWeightMsgRecordDelivery          = "op_weight_msg_delivery"
		defaultWeightMsgRecordDelivery int = 100
	)

	var weightMsgRecordDelivery int
	simState.AppParams.GetOrGenerate(opWeightMsgRecordDelivery, &weightMsgRecordDelivery, nil,
		func(_ *rand.Rand) {
			weightMsgRecordDelivery = defaultWeightMsgRecordDelivery
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRecordDelivery,
		deliverysimulation.SimulateMsgRecordDelivery(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
