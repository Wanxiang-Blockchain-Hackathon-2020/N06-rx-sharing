package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrWithdrawTooOften withdraw too often
	ErrInputInvalid        = sdkerrors.Register(ModuleName, 1, "Your input is not valid")
	ErrPatientExisted      = sdkerrors.Register(ModuleName, 2, "Patient already exists")
	ErrDoctorExisted       = sdkerrors.Register(ModuleName, 3, "Doctor already exists")
	ErrDrugStoreExisted    = sdkerrors.Register(ModuleName, 4, "Drugstore already exists")
	ErrDontHaveRx          = sdkerrors.Register(ModuleName, 5, "Don't have any rx in this address")
	ErrDrugstoreNotExisted = sdkerrors.Register(ModuleName, 6, "Drugstore does not existed")
	ErrIllegalAccess       = sdkerrors.Register(ModuleName, 7, "Illegal Rx Access")
	ErrDuplicatedUse       = sdkerrors.Register(ModuleName, 8, "Rx can be used only once.")
	ErrRxDoesNotExists     = sdkerrors.Register(ModuleName, 9, "Rx does not exists.")
)
