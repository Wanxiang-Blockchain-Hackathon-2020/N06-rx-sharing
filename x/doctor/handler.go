package doctor

import (
	"fmt"

	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/doctor/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "Rx-sharing" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgPrescribe:
			return handleMsgPrescribe(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized admin Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to Prescribe
func handleMsgPrescribe(ctx sdk.Context, keeper Keeper, msg types.MsgPrescribe) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.Prescribe(ctx, msg.Doctor, msg.Patient, msg.Encrypted, msg.Memo, msg.Envelope)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
}
