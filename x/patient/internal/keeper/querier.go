package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// query endpoints supported by the nameservice Querier
const (
	QueryRxs        = "rxs"
	QueryRx         = "rx"
	QueryPermission = "permits"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryRxs:
			return queryRxs(ctx, path[1:], req, keeper)
		case QueryRx:
			return queryRx(ctx, path[1:], req, keeper)
		case QueryPermission:
			return queryPermits(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown patient query endpoint")
		}
	}
}

func queryRxs(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value, e := keeper.ap.GetRxs(ctx, path[0])
	if e != nil {
		return nil, e
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryRx(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value, e := keeper.ap.GetRx(ctx, path[0], path[1])
	if e != nil {
		return nil, e
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryPermits(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value, e := keeper.ap.GetRxPermission(ctx, path[0])
	if e != nil {
		return nil, e
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
