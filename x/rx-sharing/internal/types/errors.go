package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrWithdrawTooOften withdraw too often
	ErrInputInvalid = sdkerrors.Register(ModuleName, 1, "Your input is not valid")
)
