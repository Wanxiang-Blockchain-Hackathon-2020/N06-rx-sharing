package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAuthorizeRx{}, "admin/authorize", nil)
	cdc.RegisterConcrete(MsgPrescribe{}, "admin/prescribe", nil)
	cdc.RegisterConcrete(MsgRegisterDocter{}, "admin/register-doctor", nil)
	cdc.RegisterConcrete(MsgRegisterDrugStore{}, "admin/register-drugstore", nil)
	cdc.RegisterConcrete(MsgRegisterPatient{}, "admin/register-patient", nil)
	cdc.RegisterConcrete(MsgSaleDrugs{}, "admin/saledrugs", nil)
}
