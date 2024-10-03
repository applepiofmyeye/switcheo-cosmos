package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePlaylist{}

func NewMsgCreatePlaylist(creator string, title string, id uint64) *MsgCreatePlaylist {
	return &MsgCreatePlaylist{
		Creator: creator,
		Title:   title,
		Id:      id,
	}
}

func (msg *MsgCreatePlaylist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
