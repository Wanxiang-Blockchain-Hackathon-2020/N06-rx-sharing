package exported

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"
	"time"
)

//type Rx interface {
//	GetID() string
//	SetID(id string)
//	GetPatient() string
//	SetPatient(patient string)
//	GetStatus() sdk.Int
//	SetStatus(id sdk.Int)
//	GetTime() time.Time
//	SetTime(t time.Time)
//	GetEncrypted() string
//	SetEncrypted(en string)
//	GetMemo() string
//	SetMemo(memo string)
//	GetSaleStore() string
//	SetSaleStore(store string)
//	GetSaleTime() time.Time
//	SetSaleTime(t time.Time)
//
//	String() string
//}

type CaseHistory = []Rx

func NewCaseHistory(rxs ...Rx) CaseHistory {
	if len(rxs) == 0 {
		return CaseHistory{}
	}

	return rxs
}

// Rx is a struct that contains all the metadata of a Rx
type Rx struct {
	ID        string    `json:"id"`
	Patient   string    `json:"patient"`
	Status    sdk.Int   `json:"status"`
	Time      time.Time `json:"time"`
	Encrypted string    `json:"encrypted"` //加密处方数据
	Memo      string    `json:"memo"`
	SaleStore string    `json:"sale_store"` //在哪个门店使用的
	SaleTime  time.Time `json:"sale_time"`  //销售时间
}

func genRxId(pubkey string) string {
	time.Now().Unix()
	id := make([]string, 2)
	id[0] = pubkey[:2]
	id[1] = strconv.FormatInt(time.Now().Unix(), 10)
	return strings.Join(id, "-")
}

// NewRx returns a new Rx
func NewRx(pubkey string, encrypted string, memo string) Rx {
	return Rx{
		ID:        genRxId(pubkey),
		Patient:   pubkey,
		Status:    sdk.NewInt(1),
		Time:      time.Now(),
		Encrypted: encrypted,
		Memo:      memo,
	}
}

// implement fmt.Stringer
func (r Rx) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s,Patient: %s, Status: %s, Time: %s, Encrypted: %s, Memo: %s`,
		r.ID, r.Patient, r.Status, r.Time, r.Encrypted, r.Memo))
}

//impliment exported.Rx
func (r Rx) GetID() string {
	return r.ID
}
func (r Rx) SetID(id string) {
	r.ID = id
}
func (r Rx) GetPatient() string {
	return r.Patient
}
func (r Rx) SetPatient(p string) {
	r.Patient = p
}
func (r Rx) GetStatus() sdk.Int {
	return r.Status
}
func (r Rx) SetStatus(status sdk.Int) {
	r.Status = status
}
func (r Rx) GetTime() time.Time {
	return r.Time
}
func (r Rx) SetTime(t time.Time) {
	r.Time = t
}
func (r Rx) GetEncrypted() string {
	return r.Encrypted
}
func (r Rx) SetEncrypted(en string) {
	r.Encrypted = en
}
func (r Rx) GetMemo() string {
	return r.Memo
}
func (r Rx) SetMemo(memo string) {
	r.Memo = memo
}
func (r Rx) GetSaleStore() string {
	return r.SaleStore
}
func (r Rx) SetSaleStore(s string) {
	r.SaleStore = s
}
func (r Rx) GetSaleTime() time.Time {
	return r.SaleTime
}
func (r Rx) SetSaleTime(t time.Time) {
	r.SaleTime = t
}

type Permission struct {
	Visitor  string
	Envelope string
}

type RxPermission = []Permission

func NewPermission(visitor string, envelope string) Permission {
	return Permission{
		Visitor:  visitor,
		Envelope: envelope,
	}
}

func NewRxPermission(prs ...Permission) RxPermission {
	if len(prs) == 0 {
		return RxPermission{}
	}

	return prs
}
