package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper is required for getting Denom
type AdminKeeper interface {
	Authorize(ctx sdk.Context, patient sdk.AccAddress, id string, recipient sdk.AccAddress, token string) error
}
