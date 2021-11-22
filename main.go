package main

import (
	encryptionProvider "encryiption-service/encryption/provider"
	encryptionService "encryiption-service/encryption/service"
	keyProvider "encryiption-service/keystore/provider"
	keyService "encryiption-service/keystore/service"
	"fmt"
)

var (
	localKeys []keyProvider.LocalKey
)

func init() {}

func main() {
	// Encryption Service
	// Receives plaintext STRING
	// Outputs base64 encoded ciphertext STRING

	// Encryption service will leverage cloud based KMS (AWS, GCP)
	// Create Interface Encrypter(Encrypt, Decrypt) which will be implemented by EncryiptionProvider
	// Create EncryptionService which will use EncryptionProvider to do encryption and decryption
	// Create a presentation layer (API, HTTP) which will receive plaintext or ciphertext in string in order to encrypt or decrypt
	localKeyStore := keyProvider.NewLocalKeyStore(localKeys)
	keyService := keyService.NewKeyService(localKeyStore)

	aesEncryptionProvider := encryptionProvider.NewAESEncryptionProvider()
	encryptionService := encryptionService.NewEncryptionService(aesEncryptionProvider)

	testKey01, err := keyService.FetchKey("test-key-01")
	if err != nil {
		panic(err)
	}

	encrypted, err := encryptionService.Encrypt(
		`The encoding pads the output to a multiple of 4 bytes so Encode is not 
		appropriate for use on individual blocks of a large data stream. 
		Use NewEncoder() instead.
		`,
		testKey01,
	)
	fmt.Println(string(encrypted))

	dec, err := encryptionService.Decrypt(encrypted, testKey01)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dec))
}
