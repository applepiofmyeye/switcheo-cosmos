package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateSong{}

func NewMsgUpdateSong(creator string, title string, id uint64) *MsgUpdateSong {
	return &MsgUpdateSong{
		Creator: creator,
		Title:   title,
		Id:      id,
	}
}

func (msg *MsgUpdateSong) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
