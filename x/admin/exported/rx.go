package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type Rx interface {
	GetID() string
	SetID(id string)
	GetPatient() string
	SetPatient(patient string)
	GetStatus() sdk.Int
	SetStatus(id sdk.Int)
	GetTime() time.Time
	SetTime(t time.Time)
	GetEncrypted() string
	SetEncrypted(en string)
	GetMemo() string
	SetMemo(memo string)
	GetSaleStore() string
	SetSaleStore(store string)
	GetSaleTime() time.Time
	SetSaleTime(t time.Time)

	String() string
}
