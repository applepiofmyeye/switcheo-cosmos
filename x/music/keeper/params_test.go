package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "music/testutil/keeper"
	"music/x/music/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MusicKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
