package haes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func Encrypt(src []byte, priKey []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, errors.New("haes.Encrypt: src is empty")
	}
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return nil, errors.New("haes.Encrypt: aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	padding := aes.BlockSize - len(src)%aes.BlockSize
	paddedPlaintext := append(src, bytes.Repeat([]byte{byte(padding)}, padding)...)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.New("haes.Encrypt: init iv failed: " + err.Error())
	}
	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)
	copy(ciphertext[:aes.BlockSize], iv)
	return ciphertext, nil
}

func Decrypt(src []byte, priKey []byte) ([]byte, error) {
	if len(src) < aes.BlockSize {
		return nil, errors.New("haes.Decrypt: src is too short")
	}
	if len(src)%aes.BlockSize != 0 {
		return nil, errors.New("haes.Decrypt: invalid ciphertext length")
	}
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return nil, errors.New("haes.Decrypt: aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	iv := src[:aes.BlockSize]
	decrypted := make([]byte, len(src)-aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, src[aes.BlockSize:])
	padding := int(decrypted[len(decrypted)-1])
	decrypted = decrypted[:len(decrypted)-padding]
	return decrypted, nil
}
