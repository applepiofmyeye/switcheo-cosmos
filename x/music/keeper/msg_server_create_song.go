package keeper

import (
	"context"

	"music/x/music/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSong(goCtx context.Context, msg *types.MsgCreateSong) (*types.MsgCreateSongResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var song = types.Song{
		Creator: msg.Creator,
		Title:   msg.Title,
	}
	id := k.AppendSong(
		ctx,
		song,
	)
	return &types.MsgCreateSongResponse{
		Id: id,
	}, nil
}
