package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

/// Rx

// MsgCreateRx defines a create rx message
type MsgPrescribe struct {
	From      sdk.AccAddress `json:"from"`
	Doctor    string         `json:"doctor"`
	Patient   string         `json:"patient"`
	Encrypted string         `json:"encrypted"` // 处方上链前需加密
	Envelope  string         `json:"envelope"`
	Memo      string         `json:"memo"`
}

// NewMsgRegisterDrugStore is a constructor function for MsgRegisterDrugStore
func NewMsgPrescribe(from sdk.AccAddress, doctor string, patient string, encrypted string, envelope string, memo string) MsgPrescribe {
	return MsgPrescribe{
		From:      from,
		Doctor:    doctor,
		Patient:   patient,
		Encrypted: encrypted,
		Envelope:  envelope,
		Memo:      memo,
	}
}

// Route should return the name of the module
func (msg MsgPrescribe) Route() string { return RouterKey }

// Type should return the action
func (msg MsgPrescribe) Type() string { return "prescribe" }

// ValidateBasic runs stateless checks on the message
func (msg MsgPrescribe) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.From.String())
	}
	if len(msg.Doctor) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Doctor)
	}
	if len(msg.Patient) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Patient)
	}
	if len(msg.Encrypted) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Encrypted)
	}
	if len(msg.Envelope) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Envelope)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPrescribe) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPrescribe) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}
