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
	cdc.RegisterConcrete(MsgAuthorizeRx{}, "rx-sharing/authorize", nil)
	cdc.RegisterConcrete(MsgCreateRx{}, "rx-sharing/create", nil)
	cdc.RegisterConcrete(MsgRegisterDocter{}, "rx-sharing/register-doctor", nil)
	cdc.RegisterConcrete(MsgRegisterDrugStore{}, "rx-sharing/register-drugstore", nil)
	cdc.RegisterConcrete(MsgRegisterPatient{}, "rx-sharing/register-patient", nil)
	cdc.RegisterConcrete(MsgSaleDrugs{}, "rx-sharing/saledrugs", nil)
}
