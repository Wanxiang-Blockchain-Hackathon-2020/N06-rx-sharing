package keeper

import (
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/drugstore/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	ap       types.AdminKeeper
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the Faucet Keeper
func NewKeeper(
	ap types.AdminKeeper,
	storeKey sdk.StoreKey,
	cdc *codec.Codec) Keeper {
	return Keeper{
		ap:       ap,
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// RegisterPatient register patient on blockchain.
func (k Keeper) SaleDrugs(ctx sdk.Context, patient sdk.AccAddress, id string, drugstore sdk.AccAddress) error {
	return k.ap.SaleDrugs(ctx, patient, id, drugstore)
}
