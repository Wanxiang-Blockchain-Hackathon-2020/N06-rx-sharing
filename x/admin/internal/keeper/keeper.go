package keeper

import (
	"bytes"
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/exported"
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

	var ch exported.CaseHistory

	if k.hasCaseHistory(ctx, patient) {
		chs, err := k.GetCaseHistory(ctx, patient)
		if err != nil {
			return err
		} else {
			ch = chs
		}
	} else {
		ch = exported.NewCaseHistory()
	}
	rx := exported.NewRx(doctor, patient, encrypted, memo)
	ch = append(ch, rx)

	k.SaveCaseHistory(ctx, patient, ch)

	var rp exported.RxPermission
	if k.hasRxPermission(ctx, rx.GetID()) {
		rp, _ = k.GetRxPermission(ctx, rx.GetID())
	} else {
		rp = exported.NewRxPermission()
	}
	rp = append(rp, exported.NewPermission(patient, token))
	k.SaveRxPermission(ctx, rx.GetID(), rp)

	return nil
}

func (k Keeper) Authorize(ctx sdk.Context, patient string, id string, recipient string, token string) error {

	permits, _ := k.GetRxPermission(ctx, id)

	for _, t := range permits { //过滤重复授权
		if t.Visitor == recipient {
			return nil
		}
	}

	permits = append(permits, exported.NewPermission(recipient, token))

	k.SaveRxPermission(ctx, id, permits)

	return nil
}

func (k Keeper) SaleDrugs(ctx sdk.Context, patient string, rxid string, drugstore string) error {

	if k.hasCaseHistory(ctx, patient) {
		chs, err := k.GetCaseHistory(ctx, patient)
		if err == nil {
			for i, t := range chs {
				if t.ID == rxid && t.Status != sdk.NewInt(types.Rx_USED) {
					permits, err := k.GetRxPermission(ctx, rxid)
					if err == nil {
						for _, p := range permits {
							if p.Visitor == drugstore {
								t.Status = sdk.NewInt(types.Rx_USED)
								t.SaleTime = time.Now()
								t.SaleStore = drugstore

								chs[i] = t
								k.SaveCaseHistory(ctx, patient, chs)
							}
						}
					}
				}
			}

		}
	}

	return types.ErrRxDoesNotExists
}

func (k Keeper) GetPatient(ctx sdk.Context, pubkey string) types.Patient {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_Patient, pubkey))
	var data types.Patient
	k.cdc.MustUnmarshalBinaryBare(bz, &data)
	return data
}

func (k Keeper) SavePatient(ctx sdk.Context, patient types.Patient) {
	store := ctx.KVStore(k.storeKey)
	store.Set(keymap(types.Prefix_Patient, patient.Pubkey), k.cdc.MustMarshalBinaryBare(patient))
}

func (k Keeper) GetDoctor(ctx sdk.Context, pubkey string) (types.Doctor, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_Doctor, pubkey))
	var data types.Doctor
	err := k.cdc.UnmarshalBinaryBare(bz, &data)
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
	err := k.cdc.UnmarshalBinaryBare(bz, &data)
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

func (k Keeper) GetCaseHistory(ctx sdk.Context, pubkey string) (exported.CaseHistory, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_CaseHistory, pubkey))
	var data exported.CaseHistory
	err := k.cdc.UnmarshalBinaryBare(bz, &data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (k Keeper) SaveCaseHistory(ctx sdk.Context, pubkey string, history exported.CaseHistory) {
	store := ctx.KVStore(k.storeKey)
	byt, err := k.cdc.MarshalBinaryBare(history)
	if err != nil {
		fmt.Println(err)
	}
	store.Set(keymap(types.Prefix_CaseHistory, pubkey), byt)
}

func (k Keeper) GetRxs(ctx sdk.Context, patient string) (exported.CaseHistory, error) {
	ch, err := k.GetCaseHistory(ctx, patient)
	if err != nil {
		return nil, err
	} else {
		return ch, nil
	}
}

func (k Keeper) GetRx(ctx sdk.Context, patient string, id string) (exported.Rx, error) {
	rxs, err := k.GetRxs(ctx, patient)
	if err != nil {
		return exported.Rx{}, err
	}

	for _, t := range rxs {
		if t.ID == id {
			return t, nil
		}
	}
	return exported.Rx{}, types.ErrRxDoesNotExists
}

func (k Keeper) GetRxPermission(ctx sdk.Context, rxid string) (exported.RxPermission, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(keymap(types.Prefix_RxPermission, rxid))
	var data exported.RxPermission
	err := k.cdc.UnmarshalBinaryBare(bz, &data)
	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (k Keeper) SaveRxPermission(ctx sdk.Context, rxid string, rxPermission exported.RxPermission) {
	store := ctx.KVStore(k.storeKey)
	store.Set(keymap(types.Prefix_RxPermission, rxid), k.cdc.MustMarshalBinaryBare(rxPermission))
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
func (k Keeper) hasRxPermission(ctx sdk.Context, key string) bool {
	return k.has(ctx, key, types.Prefix_RxPermission)
}
