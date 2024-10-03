package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"music/x/music/types"
)

func (k msgServer) UpdateSong(goCtx context.Context, msg *types.MsgUpdateSong) (*types.MsgUpdateSongResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var song = types.Song{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
	}
	val, found := k.GetSong(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetSong(ctx, song)
	return &types.MsgUpdateSongResponse{}, nil

}
