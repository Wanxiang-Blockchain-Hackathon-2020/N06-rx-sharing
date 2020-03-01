package drugstore

import (
	"fmt"

	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/drugstore/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "Rx-sharing" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {

		case types.MsgSaleDrugs:
			return handleMsgSaleDrugs(ctx, keeper, msg)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized admin Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to Sale Drugs
func handleMsgSaleDrugs(ctx sdk.Context, keeper Keeper, msg types.MsgSaleDrugs) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.SaleDrugs(ctx, msg.Patient, msg.ID, msg.DrugStore)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
}
