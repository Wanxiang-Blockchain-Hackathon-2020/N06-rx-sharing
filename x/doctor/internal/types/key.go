package types

const (
	// ModuleName is the name of the module
	ModuleName = "doctor"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	Rx_ACTIVE  = 1 //有效状态
	Rx_LOCKING = 2 //药店锁定
	Rx_USED    = 3 //完成购买
)
