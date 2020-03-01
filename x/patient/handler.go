package patient

import (
	"fmt"

	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/patient/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "Rx-sharing" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {

		case types.MsgAuthorizeRx:
			return handleMsgAuthorizeRx(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized admin Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to Authorize Rx
func handleMsgAuthorizeRx(ctx sdk.Context, keeper Keeper, msg types.MsgAuthorizeRx) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.Authorize(ctx, msg.Patient, msg.ID, msg.DrugStore, msg.Envelope)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
}
