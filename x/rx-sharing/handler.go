package rx_sharing

import (
	"fmt"

	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/rx-sharing/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "Rx-sharing" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgPrescribe:
			return handleMsgPrescribe(ctx, keeper, msg)
		case types.MsgSaleDrugs:
			return handleMsgSaleDrugs(ctx, keeper, msg)
		case types.MsgRegisterPatient:
			return handleMsgRegisterPatient(ctx, keeper, msg)
		case types.MsgRegisterDrugStore:
			return handleMsgRegisterDrugStore(ctx, keeper, msg)
		case types.MsgRegisterDocter:
			return handleMsgRegisterDocter(ctx, keeper, msg)
		case types.MsgAuthorizeRx:
			return handleMsgAuthorizeRx(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized rx-sharing Msg type: %v", msg.Type()))
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

// Handle a message to Sale Drugs
func handleMsgSaleDrugs(ctx sdk.Context, keeper Keeper, msg types.MsgSaleDrugs) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.SaleDrugs(ctx, msg.Patient, msg.ID, msg.DrugStore)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
}

// Handle a message to Register Docter
func handleMsgRegisterDocter(ctx sdk.Context, keeper Keeper, msg types.MsgRegisterDocter) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.RegisterDoctor(ctx, msg.Address, msg.Name, msg.Gender, msg.Hospital, msg.Department, msg.Title, msg.Introduction)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
}

// Handle a message to Register DrugStore
func handleMsgRegisterDrugStore(ctx sdk.Context, keeper Keeper, msg types.MsgRegisterDrugStore) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.RegisterDrugstore(ctx, msg.Address, msg.Name, msg.Phone, msg.Group, msg.BizTime, msg.Location)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
}

// Handle a message to Register Patient
func handleMsgRegisterPatient(ctx sdk.Context, keeper Keeper, msg types.MsgRegisterPatient) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received message: %s", msg)
	err := keeper.RegisterPatient(ctx, msg.Address, msg.Name, msg.Gender, msg.Birthday, msg.Encrypted, msg.Envelope)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil // return
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
