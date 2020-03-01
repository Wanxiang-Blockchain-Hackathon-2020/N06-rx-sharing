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
	Doctor    sdk.AccAddress `json:"doctor"`
	Patient   sdk.AccAddress `json:"patient"`
	Encrypted string         `json:"encrypted"` // 处方上链前需加密
	Envelope  string         `json:"envelope"`
	Memo      string         `json:"memo"`
}

// NewMsgRegisterDrugStore is a constructor function for MsgRegisterDrugStore
func NewMsgPrescribe(doctor sdk.AccAddress, patient sdk.AccAddress, encrypted string, envelope string, memo string) MsgPrescribe {
	return MsgPrescribe{
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
	if msg.Doctor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Doctor.String())
	}
	if msg.Patient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Patient.String())
	}
	if len(msg.Encrypted) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Encrypted)
	}
	if len(msg.Memo) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Encrypted)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgPrescribe) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPrescribe) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Doctor}
}
