package service

import "encoding/base64"

type EncryptionService struct {
	encrypter Encrypter
}

type Encrypter interface {
	Encrypt(plainText, key []byte) ([]byte, error)
	Decrypt(cipherText, key []byte) ([]byte, error)
}

func NewEncryptionService(e Encrypter) *EncryptionService {
	return &EncryptionService{
		encrypter: e,
	}
}

func (c *EncryptionService) Encrypt(plainText string, key []byte) (string, error) {
	encrypted, err := c.encrypter.Encrypt([]byte(plainText), key)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return encoded, nil
}

func (c *EncryptionService) Decrypt(encodedCipherText string, key []byte) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encodedCipherText)
	if err != nil {
		return "", nil
	}

	decrypted, err := c.encrypter.Decrypt(decoded, key)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
