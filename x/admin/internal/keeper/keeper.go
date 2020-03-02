package keeper

import (
	"bytes"
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the Faucet Keeper
func NewKeeper(
	storeKey sdk.StoreKey,
	cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// RegisterPatient register patient on blockchain.
func (k Keeper) RegisterPatient(ctx sdk.Context, pubkey string, name string, sex string, birthday time.Time, encrypted string, envelope string) error {

	if k.hasPatient(ctx, pubkey) {
		return types.ErrPatientExisted
	}

	p := types.NewPatient(pubkey, name, sex, birthday, encrypted, envelope)

	k.SavePatient(ctx, p)

	return nil
}

// RegisterPatient register patient on blockchain.
func (k Keeper) RegisterDoctor(ctx sdk.Context, pubkey string, name string, sex string, hospital string, department string, title string, introduction string) error {

	if k.hasDoctor(ctx, pubkey) {
		return types.ErrDoctorExisted
	}

	d := types.NewDoctor(pubkey, name, sex, hospital, department, title, introduction)

	k.SaveDoctor(ctx, d)

	return nil
}

// RegisterPatient register patient on blockchain.
func (k Keeper) RegisterDrugstore(ctx sdk.Context, pubkey string, name string, phone string, group string, biztime string, location string) error {

	if k.hasDrugstore(ctx, pubkey) {
		return types.ErrDrugStoreExisted
	}

	p := types.NewDrugStore(pubkey, name, phone, group, biztime, location)

	k.SaveDrugStore(ctx, p)

	return nil
}

// RegisterPatient register patient on blockchain.
func (k Keeper) Prescribe(ctx sdk.Context, doctor string, patient string, encrypted string, memo string, token string) error {

	var ch types.CaseHistory

	if k.hasCaseHistory(ctx, patient) {
		chs, err := k.GetCaseHistory(ctx, patient)
		if err != nil {
			return err
		} else {
			ch = chs
		}
	} else {
		ch = types.NewCaseHistory(patient)
	}
	rx := types.NewRx(patient, encrypted, memo)

	// make sure patient has access right
	rx.AddAccessToken(patient, token)
	ch.AddRx(rx)

	k.SaveCaseHistory(ctx, ch)

	return nil
}

func (k Keeper) Authorize(ctx sdk.Context, patient string, id string, recipient string, token string) error {

	if !k.hasCaseHistory(ctx, patient) {
		return types.ErrDontHaveRx
	}
	ch, err := k.GetCaseHistory(ctx, patient)
	if err != nil {
		return err
	}

	rx, ok := ch.GetRx(id)
	if !ok {
		return types.ErrDontHaveRx
	}

	rx.AddAccessToken(recipient, token)
	rx.Status = sdk.NewInt(types.Rx_LOCKING)

	ch.SetRx(rx.ID, rx)

	k.SaveCaseHistory(ctx, ch)

	return nil
}

func (k Keeper) SaleDrugs(ctx sdk.Context, patient string, id string, drugstore string) error {

	if !k.hasCaseHistory(ctx, patient) {
		return types.ErrDontHaveRx
	}
	ch, err := k.GetCaseHistory(ctx, patient)
	if err != nil {
		return err
	}

	//处方是否存在
	rx, ok := ch.GetRx(id)
	if !ok {
		return types.ErrDontHaveRx
	}

	//药店是否存在
	if !k.hasDrugstore(ctx, drugstore) {
		return types.ErrDrugstoreNotExisted
	}

	ds, err2 := k.GetDrugstore(ctx, drugstore)
	if err2 != nil {
		return types.ErrDrugstoreNotExisted
	}

	//是否获取用户授权
	_, ok2 := rx.GetAccessToken(drugstore)
	if !ok2 {
		return types.ErrIllegalAccess
	}

	//处方转态是否可用
	if rx.Status == sdk.NewInt(types.Rx_USED) {
		return types.ErrDuplicatedUse
	}

	rx.Status = sdk.NewInt(types.Rx_USED)
	rx.SaleStore = ds.Name
	rx.SaleTime = time.Now()

	ch.SetRx(rx.ID, rx)

	k.SaveCaseHistory(ctx, ch)

	return nil
}

func (k Keeper) GetPatient(ctx sdk.Context, pubkey string) (types.Patient, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_Patient, pubkey))
	var data types.Patient
	err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, &data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (k Keeper) SavePatient(ctx sdk.Context, patient types.Patient) {
	store := ctx.KVStore(k.storeKey)
	store.Set(keymap(types.Prefix_Patient, patient.Pubkey), k.cdc.MustMarshalBinaryBare(patient))
}

func (k Keeper) GetDoctor(ctx sdk.Context, pubkey string) (types.Doctor, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_Doctor, pubkey))
	var data types.Doctor
	err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, &data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (k Keeper) SaveDoctor(ctx sdk.Context, doctor types.Doctor) {
	store := ctx.KVStore(k.storeKey)
	store.Set(keymap(types.Prefix_Doctor, doctor.Pubkey), k.cdc.MustMarshalBinaryBare(doctor))
}

func (k Keeper) GetDrugstore(ctx sdk.Context, pubkey string) (types.DrugStore, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_DrugStore, pubkey))
	var data types.DrugStore
	err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, &data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (k Keeper) SaveDrugStore(ctx sdk.Context, drugstore types.DrugStore) {
	store := ctx.KVStore(k.storeKey)
	store.Set(keymap(types.Prefix_DrugStore, drugstore.Pubkey), k.cdc.MustMarshalBinaryBare(drugstore))
}

func (k Keeper) GetCaseHistory(ctx sdk.Context, pubkey string) (types.CaseHistory, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_CaseHistory, pubkey))
	var data types.CaseHistory
	err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, &data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (k Keeper) SaveCaseHistory(ctx sdk.Context, history types.CaseHistory) {
	store := ctx.KVStore(k.storeKey)
	store.Set(keymap(types.Prefix_CaseHistory, history.Patient), k.cdc.MustMarshalBinaryBare(history))
}

func keymap(prefix string, key string) []byte {
	keyBuf := bytes.NewBufferString(prefix)
	keyBuf.Write([]byte(key))
	return keyBuf.Bytes()
}

// has check if the key is present in the store or not
func (k Keeper) has(ctx sdk.Context, key string, prefix string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(keymap(prefix, key))
}

func (k Keeper) hasDoctor(ctx sdk.Context, key string) bool {
	return k.has(ctx, key, types.Prefix_Doctor)
}
func (k Keeper) hasPatient(ctx sdk.Context, key string) bool {
	return k.has(ctx, key, types.Prefix_Patient)
}
func (k Keeper) hasDrugstore(ctx sdk.Context, key string) bool {
	return k.has(ctx, key, types.Prefix_DrugStore)
}
func (k Keeper) hasCaseHistory(ctx sdk.Context, key string) bool {
	return k.has(ctx, key, types.Prefix_CaseHistory)
}
