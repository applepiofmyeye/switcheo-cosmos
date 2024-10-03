package keeper

import (
	"context"

	"music/x/music/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePlaylist(goCtx context.Context, msg *types.MsgCreatePlaylist) (*types.MsgCreatePlaylistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var song = types.Song{
		Creator: msg.Creator,
		Title:   msg.Title,
	}
	id := k.AppendSong(
		ctx,
		song,
	)
	return &types.MsgCreatePlaylistResponse{
		Playlist: id,
	}, nil
}
