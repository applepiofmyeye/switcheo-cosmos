package music_test

import (
	"testing"

	keepertest "music/testutil/keeper"
	"music/testutil/nullify"
	music "music/x/music/module"
	"music/x/music/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MusicKeeper(t)
	music.InitGenesis(ctx, k, genesisState)
	got := music.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
