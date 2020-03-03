package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// query endpoints supported by the nameservice Querier
const (
	QueryStatus = "status"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryStatus:
			return queryStatus(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown drugstore query endpoint")
		}
	}
}

func queryStatus(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	return nil, nil
}
