package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// AccountKeeper is required for getting Account
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account
}
