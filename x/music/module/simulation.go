package music

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"music/testutil/sample"
	musicsimulation "music/x/music/simulation"
	"music/x/music/types"
)

// avoid unused import issue
var (
	_ = musicsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateSong = "op_weight_msg_create_song"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSong int = 100

	opWeightMsgUpdateSong = "op_weight_msg_update_song"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSong int = 100

	opWeightMsgDeleteSong = "op_weight_msg_delete_song"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSong int = 100

	opWeightMsgCreatePlaylist = "op_weight_msg_create_playlist"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePlaylist int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	musicGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&musicGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateSong int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSong, &weightMsgCreateSong, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSong = defaultWeightMsgCreateSong
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSong,
		musicsimulation.SimulateMsgCreateSong(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSong int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSong, &weightMsgUpdateSong, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSong = defaultWeightMsgUpdateSong
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSong,
		musicsimulation.SimulateMsgUpdateSong(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSong int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteSong, &weightMsgDeleteSong, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSong = defaultWeightMsgDeleteSong
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSong,
		musicsimulation.SimulateMsgDeleteSong(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePlaylist int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePlaylist, &weightMsgCreatePlaylist, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePlaylist = defaultWeightMsgCreatePlaylist
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePlaylist,
		musicsimulation.SimulateMsgCreatePlaylist(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSong,
			defaultWeightMsgCreateSong,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				musicsimulation.SimulateMsgCreateSong(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSong,
			defaultWeightMsgUpdateSong,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				musicsimulation.SimulateMsgUpdateSong(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteSong,
			defaultWeightMsgDeleteSong,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				musicsimulation.SimulateMsgDeleteSong(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePlaylist,
			defaultWeightMsgCreatePlaylist,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				musicsimulation.SimulateMsgCreatePlaylist(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
