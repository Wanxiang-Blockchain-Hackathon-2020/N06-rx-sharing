package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SupplyKeeper is required for mining coin
type AdminKeeper interface {
	// Prescribe is used for doctor to prescribe on blockchain.
	Prescribe(ctx sdk.Context, doctor string, patient string, encrypted string, memo string, token string) error
}
