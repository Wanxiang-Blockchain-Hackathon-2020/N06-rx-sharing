package types

import (
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper is required for getting Denom
type AdminKeeper interface {
	Authorize(ctx sdk.Context, patient string, id string, recipient string, envelope string) error
	GetRxs(ctx sdk.Context, pubkey string) (exported.CaseHistory, error)
	GetRx(ctx sdk.Context, pubkey string, id string) (exported.Rx, error)
	GetRxPermission(ctx sdk.Context, rxid string) (exported.RxPermission, error)
}
