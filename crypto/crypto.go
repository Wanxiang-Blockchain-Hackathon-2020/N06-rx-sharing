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
		return encryptSM2(seed, pubkey, content)
	} else {
		return encryptX25519(seed, pubkey, content)
	}
}

func Dencrypt(crypto string, pubkey string, seed []byte, envelope string, cryptContent string) (string, error) {
	if crypto == "sm2" {
		return decryptSM2(pubkey, seed, envelope, cryptContent)
	} else {
		return decryptX25519(pubkey, seed, envelope, cryptContent)
	}
}

func RenewEnvelope(crypto string, seed []byte, old string, news string, envelope string) (string, error) {
	if crypto == "sm2" {
		return renewEnvelopeSM2(seed, old, news, envelope)
	} else {
		return renewEnvelope(seed, old, news, envelope)
	}
}
