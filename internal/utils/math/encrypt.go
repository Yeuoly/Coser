package math

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

/*
This file is used to implement the encryption algorithm.
*/
const AESBlockSize = aes.BlockSize

// padPKCS7 pads the input data to the specified block size using PKCS7 padding.
func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// unpadPKCS7 removes PKCS7 padding from the input data.
func unpadPKCS7(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

// EncryptWithCBCAndIV encrypts the plaintext using AES in CBC mode with variable key lengths and a dynamic IV.
func EncryptWithCBCAndIV(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	keySize := len(key)
	if keySize != 16 && keySize != 24 && keySize != 32 {
		return nil, errors.New("invalid key size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedPlaintext := padPKCS7(plaintext, AESBlockSize)

	ciphertext := make([]byte, len(paddedPlaintext)+AESBlockSize)
	copy(ciphertext[:AESBlockSize], iv)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[AESBlockSize:], paddedPlaintext)

	return ciphertext, nil
}

// DecryptWithCBCAndIV decrypts the ciphertext using AES in CBC mode with variable key lengths and a dynamic IV.
func DecryptWithCBCAndIV(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	keySize := len(key)
	if keySize != 16 && keySize != 24 && keySize != 32 {
		return nil, errors.New("invalid key size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < AESBlockSize*2 {
		return nil, errors.New("ciphertext too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decryptedText := make([]byte, len(ciphertext)-AESBlockSize)
	mode.CryptBlocks(decryptedText, ciphertext[AESBlockSize:])

	return unpadPKCS7(decryptedText), nil
}

/*
Encrypt the string with AES-CBC algorithm.
key should be 16, 24 or 32 bytes long.
iv should be 16 bytes long.
result is base64 encoded.
*/
func AesEncrypt(str string, key []byte, iv []byte) (string, error) {
	data, err := EncryptWithCBCAndIV([]byte(str), key, iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

/*
Decrypt the string with AES-CBC algorithm.
key should be 16, 24 or 32 bytes long.
iv should be 16 bytes long.
str is base64 encoded.
*/
func AesDecrypt(str string, key []byte, iv []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	dnData, err := DecryptWithCBCAndIV(data, key, iv)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}
