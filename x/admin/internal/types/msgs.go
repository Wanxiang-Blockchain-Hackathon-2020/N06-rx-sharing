package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgRegisterPatient defines a register patient message
type MsgRegisterPatient struct {
	From      sdk.AccAddress `json:"from"`
	Pubkey    string         `json:"address"`
	Name      string         `json:"name"`
	Gender    string         `json:"gender"`
	Birthday  time.Time      `json:"birthday"`
	Encrypted string         `json:"encrypted"` //加密信息，如，疾病史，家族史，过敏药物等等
	Envelope  string         `json:"envelope"`  //密码信封
}

// NewMsgRegisterPatient is a constructor function for MsgRegisterPatient
func NewMsgRegisterPatient(from sdk.AccAddress, pubkey string, name string, gender string, birthday time.Time, encrypted string, envelope string) MsgRegisterPatient {
	return MsgRegisterPatient{
		From:      from,
		Pubkey:    pubkey,
		Name:      name,
		Gender:    gender,
		Birthday:  birthday,
		Encrypted: encrypted,
		Envelope:  envelope,
	}
}

// Route should return the name of the module
func (msg MsgRegisterPatient) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterPatient) Type() string { return "register_patient" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterPatient) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.From.String())
	}
	if len(msg.Pubkey) < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, msg.Pubkey)
	}
	if len(msg.Name) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Name)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterPatient) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterPatient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

/// Doctor message

// MsgRegisterDocter defines a register doctor message
type MsgRegisterDocter struct {
	From         sdk.AccAddress `json:"from"`
	PubKey       string         `json:"pubkey"`
	Name         string         `json:"name"`
	Gender       string         `json:"gender"`
	Hospital     string         `json:"hospital"`     //就职医院
	Department   string         `json:"department"`   //所在科室
	Title        string         `json:"title"`        //职称
	Introduction string         `json:"introduction"` //介绍
}

// NewMsgRegisterDocter is a constructor function for MsgRegisterDocter
func NewMsgRegisterDocter(from sdk.AccAddress, pubkey string, name string, gender string, hospital string, department string, title string, introduction string) MsgRegisterDocter {
	return MsgRegisterDocter{
		From:         from,
		PubKey:       pubkey,
		Name:         name,
		Gender:       gender,
		Hospital:     hospital,
		Department:   hospital,
		Title:        title,
		Introduction: introduction,
	}
}

// Route should return the name of the module
func (msg MsgRegisterDocter) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterDocter) Type() string { return "register_doctor" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterDocter) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.From.String())
	}
	if len(msg.PubKey) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.PubKey)
	}
	if len(msg.Name) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Name)
	}
	if len(msg.Hospital) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Hospital)
	}
	if len(msg.Department) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Department)
	}
	if len(msg.Title) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Title)
	}
	if len(msg.Introduction) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Introduction)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterDocter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterDocter) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

/// Drag Store

// MsgRegisterDrugStore defines a register drugstore message
type MsgRegisterDrugStore struct {
	From     sdk.AccAddress `json:"from"`
	Pubkey   string         `json:"pubkey"`
	Name     string         `json:"name"`
	Phone    string         `json:"phone"`
	Group    string         `json:"group"`    //所属连锁集团
	BizTime  string         `json:"biz_time"` //营业时间
	Location string         `json:"location"` //门店地址
}

// NewMsgRegisterDrugStore is a constructor function for MsgRegisterDrugStore
func NewMsgRegisterDrugStore(from sdk.AccAddress, pubkey string, name string, phone string, group string, biztime string, location string) MsgRegisterDrugStore {
	return MsgRegisterDrugStore{
		From:     from,
		Pubkey:   pubkey,
		Name:     name,
		Phone:    phone,
		Group:    group,
		BizTime:  biztime,
		Location: location,
	}
}

// Route should return the name of the module
func (msg MsgRegisterDrugStore) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterDrugStore) Type() string { return "register_drug_store" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterDrugStore) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.From.String())
	}
	if len(msg.Pubkey) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Pubkey)
	}
	if len(msg.Name) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Name)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterDrugStore) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterDrugStore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}
