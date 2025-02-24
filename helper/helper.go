package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var encryptionKey = []byte("your_secret_key_12345678")

func EncryptData(data string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func GetJWTKey() string {
	return "your_secret_key"
}
