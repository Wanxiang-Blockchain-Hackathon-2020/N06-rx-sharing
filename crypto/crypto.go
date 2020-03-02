package crypto

import (
	"math/rand"
	"time"
)

func GenerateKey(crypto string, seed []byte) (string, error) {
	if crypto == "sm2" {
		return generateKeySM2(seed)
	} else {
		return generateKeyED25519(seed)
	}
}

func RandBytes(size int) []uint8 {
	rand.Seed(time.Now().UnixNano())
	high := 255
	var seed = make([]uint8, size)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(rand.Int() % (high + 1))
	}
	return seed
}

func Encrypt(crypto string, seed []byte, pubkey string, content string) (string, string, error) {
	if crypto == "sm2" {
		return "", "", nil
	} else {
		return encryptX25519(seed, pubkey, content)
	}
}

func Dencrypt(crypto string, pubkey string, seed []byte, envelope string, cryptContent string) (string, error) {
	if crypto == "sm2" {
		return "", nil
	} else {
		return decryptX25519(pubkey, seed, envelope, cryptContent)
	}
}
