package donation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	donationsimulation "overgive-chain/x/donation/simulation"
	"overgive-chain/x/donation/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	donationGenesis := types.GenesisState{
		Params: types.DefaultParams(),
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&donationGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgRecordDonation          = "op_weight_msg_donation"
		defaultWeightMsgRecordDonation int = 100
	)

	var weightMsgRecordDonation int
	simState.AppParams.GetOrGenerate(opWeightMsgRecordDonation, &weightMsgRecordDonation, nil,
		func(_ *rand.Rand) {
			weightMsgRecordDonation = defaultWeightMsgRecordDonation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRecordDonation,
		donationsimulation.SimulateMsgRecordDonation(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
