package crypto

import (
	"crypto/elliptic"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

var one = new(big.Int).SetInt64(1)

func randFieldElement(c elliptic.Curve, seed []byte) (k *big.Int, err error) {
	params := c.Params()
	//b := make([]byte, params.BitSize/8+8)
	//_, err = io.ReadFull(rand, b)
	//if err != nil {
	//	return
	//}
	k = new(big.Int).SetBytes(seed)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)
	return
}

func GenerateKey(seed []byte) (*sm2.PrivateKey, error) {
	c := sm2.P256Sm2()
	k, err := randFieldElement(c, seed)
	if err != nil {
		return nil, err
	}
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}
