package dollar

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"dollar/testutil/sample"
	dollarsimulation "dollar/x/dollar/simulation"
	"dollar/x/dollar/types"
)

// avoid unused import issue
var (
	_ = dollarsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePost int = 100

	opWeightMsgUpdatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePost int = 100

	opWeightMsgDeletePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePost int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dollarGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PostList: []types.Post{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PostCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dollarGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		dollarsimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePost, &weightMsgUpdatePost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePost = defaultWeightMsgUpdatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePost,
		dollarsimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		dollarsimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePost,
			defaultWeightMsgCreatePost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				dollarsimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePost,
			defaultWeightMsgUpdatePost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				dollarsimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePost,
			defaultWeightMsgDeletePost,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				dollarsimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
