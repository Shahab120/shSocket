package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
)

// EncryptECB encrypts data using AES-128 in ECB mode
func EncryptECB(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return []byte{}, err
	}

	if len(plaintext)%aes.BlockSize != 0 {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}

	ciphertext := make([]byte, len(plaintext))
	for bs, be := 0, block.BlockSize(); bs < len(plaintext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(ciphertext[bs:be], plaintext[bs:be])
	}

	return ciphertext, nil
}

// DecryptECB decrypts data using AES-128 in ECB mode
func DecryptECB(cipher []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	if len(cipher)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(cipher))
	for bs, be := 0, block.BlockSize(); bs < len(cipher); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(plaintext[bs:be], cipher[bs:be])
	}

	plaintext = PKCS5Unpadding(plaintext)

	return plaintext, nil
}

// PKCS5Padding pads the plaintext to be a multiple of the block size
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// PKCS5Unpadding removes padding from the plaintext
func PKCS5Unpadding(src []byte) []byte {
	if len(src) > 0 {
		length := len(src)
		unpadding := int(src[length-1])
		return src[:(length - unpadding)]
	}
	return src
}
