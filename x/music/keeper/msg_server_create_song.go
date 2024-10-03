package keeper

import (
	"context"

	"music/x/music/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k msgServer) CreateSong(goCtx context.Context, msg *types.MsgCreateSong) (*types.MsgCreateSongResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	 // Check if the sender has the proper authority to perform this action
	 if msg.Creator != k.authority {
        return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "only the authority %s can create songs", k.authority)
    }

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
