package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgAuthorizeRx defines a message to authorize dragstore accessing rx
type MsgAuthorizeRx struct {
	From      sdk.AccAddress `json:"from"`
	Patient   string         `json:"patient"`
	DrugStore string         `json:"drugstore"`
	ID        string         `json:"id"` // 处方ID
	Envelope  string         `json:"envelope"`
}

// NewMsgAuthorizeRx is a constructor function for MsgAuthorizeRx
func NewMsgAuthorizeRx(from sdk.AccAddress, patient string, drugstore string, id string, envelope string) MsgAuthorizeRx {
	return MsgAuthorizeRx{
		From:      from,
		Patient:   patient,
		DrugStore: drugstore,
		ID:        id,
		Envelope:  envelope,
	}
}

// Route should return the name of the module
func (msg MsgAuthorizeRx) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAuthorizeRx) Type() string { return "authorize_rx" }

// ValidateBasic runs stateless checks on the message
func (msg MsgAuthorizeRx) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.From.String())
	}
	if len(msg.ID) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.ID)
	}

	if len(msg.Patient) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Patient)
	}

	if len(msg.DrugStore) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.DrugStore)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAuthorizeRx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAuthorizeRx) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}
