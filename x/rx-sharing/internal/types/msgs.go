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
	Address   sdk.AccAddress `json:"address"`
	Name      string         `json:"name"`
	Sex       string         `json:"sex"`
	Birthday  time.Time      `json:"birthday"`
	Encrypted string         `json:"encrypted"` //加密信息，如，疾病史，家族史，过敏药物等等
	Envelope  string         `json:"envelope"`  //密码信封
}

// NewMsgRegisterPatient is a constructor function for MsgRegisterPatient
func NewMsgRegisterPatient(address sdk.AccAddress, name string, sex string, birthday time.Time, encrypted string) MsgRegisterPatient {
	return MsgRegisterPatient{
		Address:   address,
		Name:      name,
		Sex:       sex,
		Birthday:  birthday,
		Encrypted: encrypted,
	}
}

// Route should return the name of the module
func (msg MsgRegisterPatient) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterPatient) Type() string { return "register_patient" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterPatient) ValidateBasic() error {
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Address.String())
	}
	if len(msg.Name) < 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Name)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterPatient) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterPatient) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

/// Doctor message

// MsgRegisterDocter defines a register doctor message
type MsgRegisterDocter struct {
	Address      sdk.AccAddress `json:"address"`
	Name         string         `json:"name"`
	Sex          string         `json:"sex"`
	Hospital     string         `json:"hospital"`     //就职医院
	Department   string         `json:"department"`   //所在科室
	Title        string         `json:"title"`        //职称
	Introduction string         `json:"introduction"` //介绍
}

// NewMsgRegisterDocter is a constructor function for MsgRegisterDocter
func NewMsgRegisterDocter(address sdk.AccAddress, name string, sex string, hospital string, department string, title string, introduction string) MsgRegisterDocter {
	return MsgRegisterDocter{
		Address:      address,
		Name:         name,
		Sex:          sex,
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
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Address.String())
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
	return []sdk.AccAddress{msg.Address}
}

/// Drag Store

// MsgRegisterDrugStore defines a register drugstore message
type MsgRegisterDrugStore struct {
	Address  sdk.AccAddress `json:"address"`
	Name     string         `json:"name"`
	Phone    string         `json:"phone"`
	Group    string         `json:"group"`    //所属连锁集团
	BizTime  string         `json:"biz_time"` //营业时间
	Location string         `json:"location"` //门店地址
}

// NewMsgRegisterDrugStore is a constructor function for MsgRegisterDrugStore
func NewMsgRegisterDrugStore(address sdk.AccAddress, name string, phone string, group string, biztime string, location string) MsgRegisterDrugStore {
	return MsgRegisterDrugStore{
		Address:  address,
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
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Address.String())
	}
	if len(msg.Name) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Name)
	}
	if len(msg.Phone) < 1 {
		return sdkerrors.Wrap(ErrInputInvalid, msg.Phone)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterDrugStore) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterDrugStore) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

/// Rx

// MsgCreateRx defines a create rx message
type MsgCreateRx struct {
	Doctor    sdk.AccAddress `json:"doctor"`
	Patient   sdk.AccAddress `json:"patient"`
	Encrypted string         `json:"encrypted"` // 处方上链前需加密
	Memo      string         `json:"memo"`
}

// NewMsgRegisterDrugStore is a constructor function for MsgRegisterDrugStore
func NewMsgCreateRx(doctor sdk.AccAddress, patient sdk.AccAddress, encrypted string, memo string) MsgCreateRx {
	return MsgCreateRx{
		Doctor:    doctor,
		Patient:   patient,
		Encrypted: encrypted,
		Memo:      memo,
	}
}

// Route should return the name of the module
func (msg MsgCreateRx) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateRx) Type() string { return "create_rx" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateRx) ValidateBasic() error {
	if msg.Doctor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Doctor.String())
	}
	if msg.Patient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Doctor.String())
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
func (msg MsgCreateRx) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateRx) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Doctor}
}

// MsgAuthorizeRx defines a message to authorize dragstore accessing rx
type MsgAuthorizeRx struct {
	Patient   sdk.AccAddress `json:"patient"`
	DrugStore sdk.AccAddress `json:"drugstore"`
	ID        string         `json:"id"` // 处方ID
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

///

// MsgAuthorizeRx defines a message to authorize dragstore accessing rx
type MsgSaleDrugs struct {
	Patient   sdk.AccAddress `json:"patient"`
	DrugStore sdk.AccAddress `json:"drugstore"`
	ID        string         `json:"id"` // 处方ID
}

// NewMsgSaleDrugs is a constructor function for MsgSaleDrugs
func NewMsgSaleDrugs(patient sdk.AccAddress, drugstore sdk.AccAddress, id string) MsgSaleDrugs {
	return MsgSaleDrugs{
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
func (msg MsgSaleDrugs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaleDrugs) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Patient}
}
