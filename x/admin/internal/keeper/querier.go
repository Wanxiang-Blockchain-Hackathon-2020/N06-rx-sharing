package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// query endpoints supported by the nameservice Querier
const (
	QueryPatient = "patient"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryPatient:
			return queryPatient(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown admin query endpoint")
		}
	}
}

func queryPatient(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {

	value := keeper.GetPatient(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
