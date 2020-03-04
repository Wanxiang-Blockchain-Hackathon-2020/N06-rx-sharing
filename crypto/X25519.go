package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	ed25519 "github.com/miguelsandro/curve25519-go/axlsign"
	"io"
	"strings"
)

func generateKeyED25519(seed []byte) (string, error) {
	keypare := ed25519.GenerateKeyPair(seed)
	return hex.EncodeToString(keypare.PublicKey), nil
}

func aesEncrypt(key []byte, raw []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, raw, nil)

	ret := []string{
		base64.StdEncoding.EncodeToString(ciphertext),
		base64.StdEncoding.EncodeToString(nonce),
	}

	return strings.Join(ret, "#"), nil
}

func aesDecrypt(key []byte, content string) ([]byte, error) {
	text := strings.Split(content, "#")
	if len(text) != 2 {
		return nil, errors.New("your input is invalid")
	}
	//key, _ := hex.DecodeString(keytext)
	ciphertext, _ := base64.StdEncoding.DecodeString(text[0])
	nonce, _ := base64.StdEncoding.DecodeString(text[1])

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("decrypt error")
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.New("decrypt error")
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("decrypt error")
	}

	return plaintext, nil
}

func encryptX25519(seed []byte, pubkey string, content string) (string, string, error) {

	//随机密码加密数据内容
	datakey := RandBytes(16) //随机数据秘钥
	cipherContent, err := aesEncrypt(datakey, []byte(content))

	if err != nil {
		return "", "", err
	}

	keypair := ed25519.GenerateKeyPair(seed)
	b64pk, _ := hex.DecodeString(pubkey)
	sk := ed25519.SharedKey(keypair.PrivateKey, []uint8(b64pk))

	envelope, err2 := aesEncrypt([]byte(sk), datakey) //生成密码信封
	if err2 != nil {
		return "", "", err2
	}

	return cipherContent, envelope, nil
}

func decryptX25519(pubkey string, seed []byte, envelope string, cryptText string) (string, error) {

	keypair := ed25519.GenerateKeyPair(seed)
	b64pk, _ := hex.DecodeString(pubkey)
	sk := ed25519.SharedKey(keypair.PrivateKey, []uint8(b64pk))

	datakey, err := aesDecrypt(sk, envelope)
	if err != nil {
		return "", err
	}

	content, err2 := aesDecrypt(datakey, cryptText)
	if err2 != nil {
		return "", err2
	}
	return string(content), nil
}

func renewEnvelope(seed []byte, o string, n string, envelope string) (string, error) {
	keypair := ed25519.GenerateKeyPair(seed)
	pko, _ := hex.DecodeString(o)
	sko := ed25519.SharedKey(keypair.PrivateKey, []uint8(pko))

	datakey, err := aesDecrypt(sko, envelope)
	if err != nil {
		return "", err
	}

	pkn, _ := hex.DecodeString(n)
	skn := ed25519.SharedKey(keypair.PrivateKey, []uint8(pkn))
	news, err2 := aesEncrypt([]byte(skn), datakey) //生成密码信封
	if err2 != nil {
		return "", err2
	}
	return news, nil
}
