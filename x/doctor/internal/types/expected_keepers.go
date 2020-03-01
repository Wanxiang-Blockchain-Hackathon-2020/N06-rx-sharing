package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SupplyKeeper is required for mining coin
type AdminKeeper interface {
	// Prescribe is used for doctor to prescribe on blockchain.
	Prescribe(ctx sdk.Context, doctor sdk.AccAddress, patient sdk.AccAddress, encrypted string, memo string, token string) error
}
