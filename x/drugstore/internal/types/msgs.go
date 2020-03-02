package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

///

// MsgAuthorizeRx defines a message to authorize dragstore accessing rx
type MsgSaleDrugs struct {
	From      sdk.AccAddress `json:"from"`
	Patient   string         `json:"patient"`
	DrugStore string         `json:"drugstore"`
	ID        string         `json:"id"` // 处方ID
}

// NewMsgSaleDrugs is a constructor function for MsgSaleDrugs
func NewMsgSaleDrugs(from sdk.AccAddress, patient string, drugstore string, id string) MsgSaleDrugs {
	return MsgSaleDrugs{
		From:      from,
		Patient:   patient,
		DrugStore: drugstore,
		ID:        id,
	}
}

// Route should return the name of the module
func (msg MsgSaleDrugs) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSaleDrugs) Type() string { return "sale_drugs" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSaleDrugs) ValidateBasic() error {
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
func (msg MsgSaleDrugs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaleDrugs) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}
