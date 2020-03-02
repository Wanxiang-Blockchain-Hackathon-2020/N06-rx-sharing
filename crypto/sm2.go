package crypto

import (
	"crypto/elliptic"
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"log"
	"math/big"
)

var one = new(big.Int).SetInt64(1)

func randFieldElement(c elliptic.Curve, seed []byte) (k *big.Int, err error) {
	params := c.Params()
	k = new(big.Int).SetBytes(seed)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)
	return
}

func generateKeySM2(seed []byte) (string, error) {
	c := sm2.P256Sm2()
	k, err := randFieldElement(c, seed)
	if err != nil {
		return "", err
	}
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())

	buff, err4 := sm2.WritePublicKeytoMem(priv.Public().(*sm2.PublicKey), nil)
	if err4 != nil {
		return "", err4
	}

	pub := base64.StdEncoding.EncodeToString(buff)
	return pub, nil
}

func encryptSM4() {
	key := []byte("1234567890abcdef")
	fmt.Printf("key = %v\n", key)
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	sm4.WriteKeyToPem("key.pem", key, nil)
	key, err := sm4.ReadKeyFromPem("key.pem", nil)
	fmt.Printf("key = %v\n", key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data = %x\n", data)
	c, err := sm4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	d0 := make([]byte, 16)
	c.Encrypt(d0, data)
	fmt.Printf("d0 = %x\n", d0)
	d1 := make([]byte, 16)
	c.Decrypt(d1, d0)
	fmt.Printf("d1 = %x\n", d1)
}
