package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgAuthorizeRx defines a message to authorize dragstore accessing rx
type MsgAuthorizeRx struct {
	Patient   sdk.AccAddress `json:"patient"`
	DrugStore sdk.AccAddress `json:"drugstore"`
	ID        string         `json:"id"` // 处方ID
	Envelope  string         `json:"envelope"`
}

// NewMsgAuthorizeRx is a constructor function for MsgAuthorizeRx
func NewMsgAuthorizeRx(patient sdk.AccAddress, drugstore sdk.AccAddress, id string) MsgAuthorizeRx {
	return MsgAuthorizeRx{
		Patient:   patient,
		DrugStore: drugstore,
		ID:        id,
	}
}

// Route should return the name of the module
func (msg MsgAuthorizeRx) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAuthorizeRx) Type() string { return "authorize_rx" }

// ValidateBasic runs stateless checks on the message
func (msg MsgAuthorizeRx) ValidateBasic() error {
	if msg.Patient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Patient.String())
	}
	if msg.DrugStore.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.DrugStore.String())
	}
	if len(msg.ID) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.ID)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAuthorizeRx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAuthorizeRx) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Patient}
}
