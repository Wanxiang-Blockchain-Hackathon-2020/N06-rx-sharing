package types

import (
	"fmt"
	"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/x/admin/exported"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Patient is a struct that contains all the metadata of a patient
type Patient struct {
	Pubkey    string    `json:"pubkey"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	Encrypted string    `json:"encrypted"` //加密信息，如，疾病史，家族史，过敏药物等等
	Envelope  string    `json:"envelope"`  //密码信封
}

// NewPatient returns a new Patient
func NewPatient(pubkey string, name string, gender string, birthday time.Time, encrypted string, envelope string) Patient {
	return Patient{
		Pubkey:    pubkey,
		Name:      name,
		Gender:    gender,
		Birthday:  birthday,
		Encrypted: encrypted,
		Envelope:  envelope,
	}
}

// implement fmt.Stringer
func (p Patient) String() string {
	return strings.TrimSpace(fmt.Sprintf(`pubkey: %s, Name: %s, Sex: %s, Birthday: %s, Encrypted: %s`,
		p.Pubkey, p.Name, p.Gender, p.Birthday, p.Encrypted))
}

/// Doctor

// Doctor is a struct that contains all the metadata of a doctor
type Doctor struct {
	Pubkey       string `json:"pubkey"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	Hospital     string `json:"hospital"`     //就职医院
	Department   string `json:"department"`   //所在科室
	Title        string `json:"title"`        //职称
	Introduction string `json:"introduction"` //介绍
}

// NewDoctor returns a new Doctor
func NewDoctor(pubkey string, name string, gender string, hospital string, department string, title string, introduction string) Doctor {
	return Doctor{
		Pubkey:       pubkey,
		Name:         name,
		Gender:       gender,
		Hospital:     hospital,
		Department:   hospital,
		Title:        title,
		Introduction: introduction,
	}
}

// implement fmt.Stringer
func (d Doctor) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Pubkey: %s, Name: %s, Sex: %s, Hospital: %s, Department: %s, Title: %s, Introduction: %s`,
		d.Pubkey, d.Name, d.Gender, d.Hospital, d.Department, d.Title, d.Introduction))
}

/// DrugStore

// DrugStore is a struct that contains all the metadata of a DrugStore
type DrugStore struct {
	Pubkey   string `json:"pubkey"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Group    string `json:"group"`    //所属连锁集团
	BizTime  string `json:"biz_time"` //营业时间
	Location string `json:"location"` //门店地址
}

// NewDrugStore returns a new DrugStore
func NewDrugStore(pubkey string, name string, phone string, group string, biztime string, location string) DrugStore {
	return DrugStore{
		Pubkey:   pubkey,
		Name:     name,
		Phone:    phone,
		Group:    group,
		BizTime:  biztime,
		Location: location,
	}
}

// implement fmt.Stringer
func (d DrugStore) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Address: %s, Name: %s, Phone: %s, Group: %s, BizTime: %s, Location: %s`,
		d.Pubkey, d.Name, d.Phone, d.Group, d.BizTime, d.Location))
}

const (
	Rx_ACTIVE  = 1 //有效状态
	Rx_LOCKING = 2 //药店锁定
	Rx_USED    = 3 //完成购买
)

/// Case History 病历

// Rx is a struct that contains all the metadata of a Rx
type CaseHistory struct {
	Patient string                 `json:"patient"`
	Rxs     map[string]exported.Rx `json:"rxs"`
}

func NewCaseHistory(patient string) CaseHistory {
	return CaseHistory{
		Patient: patient,
		Rxs:     make(map[string]exported.Rx),
	}
}

func (ch CaseHistory) AddRx(rx exported.Rx) {
	ch.Rxs[rx.GetID()] = rx
}

func (ch CaseHistory) SetRx(id string, rx exported.Rx) {
	ch.Rxs[id] = rx
}

func (ch CaseHistory) GetRx(id string) (exported.Rx, bool) {
	rx, ok := ch.Rxs[id]
	return rx, ok
}

func (ch CaseHistory) UpdateStatus(id string, status sdk.Int) {
	rx := ch.Rxs[id]
	rx.SetStatus(status)
	ch.Rxs[id] = rx
}

/// Rx 处方

// Rx is a struct that contains all the metadata of a Rx
type rx struct {
	ID        string            `json:"id"`
	Patient   string            `json:"patient"`
	Status    sdk.Int           `json:"status"`
	Time      time.Time         `json:"time"`
	Encrypted string            `json:"encrypted"` //加密处方数据
	tokens    map[string]string `json:"tokens"`    //秘钥信封
	Memo      string            `json:"memo"`
	SaleStore string            `json:"sale_store"` //在哪个门店使用的
	SaleTime  time.Time         `json:"sale_time"`  //销售时间
}

func genRxId(pubkey string) string {
	time.Now().Unix()
	id := []string{}
	id = append(id, pubkey[:2])
	id = append(id, "-")
	id = append(id, string(time.Now().Unix()))
	return strings.Join(id, "")
}

// NewRx returns a new Rx
func NewRx(pubkey string, encrypted string, memo string) exported.Rx {
	return rx{
		ID:        genRxId(pubkey),
		Patient:   pubkey,
		Status:    sdk.NewInt(Rx_ACTIVE),
		Time:      time.Now(),
		Encrypted: encrypted,
		tokens:    make(map[string]string),
		Memo:      memo,
	}
}

// implement fmt.Stringer
func (r rx) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Patient: %s, Status: %s, Time: %s, Encrypted: %s, Token: %s, Memo: %s`,
		r.Patient, r.Status, r.Time, r.Encrypted, r.tokens, r.Memo))
}

//impliment exported.Rx
func (r rx) GetID() string {
	return r.ID
}
func (r rx) SetID(id string) {
	r.ID = id
}
func (r rx) GetPatient() string {
	return r.Patient
}
func (r rx) SetPatient(p string) {
	r.Patient = p
}
func (r rx) GetStatus() sdk.Int {
	return r.Status
}
func (r rx) SetStatus(status sdk.Int) {
	r.Status = status
}
func (r rx) GetTime() time.Time {
	return r.Time
}
func (r rx) SetTime(t time.Time) {
	r.Time = t
}
func (r rx) GetEncrypted() string {
	return r.Encrypted
}
func (r rx) SetEncrypted(en string) {
	r.Encrypted = en
}
func (r rx) GetMemo() string {
	return r.Memo
}
func (r rx) SetMemo(memo string) {
	r.Memo = memo
}
func (r rx) GetSaleStore() string {
	return r.SaleStore
}
func (r rx) SetSaleStore(s string) {
	r.SaleStore = s
}
func (r rx) GetSaleTime() time.Time {
	return r.SaleTime
}
func (r rx) SetSaleTime(t time.Time) {
	r.SaleTime = t
}
func (r rx) AddAccessToken(recipient string, token string) {
	r.tokens[recipient] = token
}
func (r rx) GetAccessToken(recipient string) (string, bool) {
	v, ok := r.tokens[recipient]
	return v, ok
}
