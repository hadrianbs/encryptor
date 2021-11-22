package provider

// This is a local encryption provider
// Which means that the key is stored locally, and loaded to memory on build time

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"time"
)

type AESEncryptionProvider struct {
}

func NewAESEncryptionProvider() *AESEncryptionProvider {
	return &AESEncryptionProvider{}
}

func (c *AESEncryptionProvider) Encrypt(plainText, key []byte) ([]byte, error) {
	defer trackTime(time.Now(), "encryption")
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(aesCipher, 512)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nonce, nonce, plainText, nil)
	return encrypted, nil
}

func (c *AESEncryptionProvider) Decrypt(cipherText, key []byte) ([]byte, error) {
	defer trackTime(time.Now(), "decryption")

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(aesCipher, 512)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("ciphertext and nonce size missmatch")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed.String())
}
