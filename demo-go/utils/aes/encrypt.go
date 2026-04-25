package utils

import (
	// bytes 用于构造和比较 PKCS7 补位字节。
	"bytes"
	// crypto/aes 提供 AES 分组加密算法实现。
	"crypto/aes"
	// crypto/cipher 提供 CBC 这样的分组加密工作模式。
	"crypto/cipher"
	// encoding/base64 用于把字符串密文解码为原始字节。
	"encoding/base64"
)

const (
	aesSecretKey = "ailearning#key01"
	aesSecretIV  = "ailearning#iv001"
)

func AESDecrypt(data string) string {
	if data == "" {
		return data
	}

	encryptedBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return data
	}

	block, err := aes.NewCipher([]byte(aesSecretKey))
	if err != nil {
		return data
	}

	if len(encryptedBytes) == 0 || len(encryptedBytes)%aes.BlockSize != 0 {
		return data
	}

	mode := cipher.NewCBCDecrypter(block, []byte(aesSecretIV))
	decryptedBytes := make([]byte, len(encryptedBytes))
	mode.CryptBlocks(decryptedBytes, encryptedBytes)

	unpaddedBytes, err := pkcs7Unpad(decryptedBytes, aes.BlockSize)
	if err != nil {
		return data
	}

	return string(unpaddedBytes)
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 || length%blockSize != 0 {
		return nil, aes.KeySizeError(length)
	}

	paddingLength := int(data[length-1])
	if paddingLength == 0 || paddingLength > blockSize || paddingLength > length {
		return nil, aes.KeySizeError(paddingLength)
	}

	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	if !bytes.Equal(data[length-paddingLength:], padding) {
		return nil, aes.KeySizeError(paddingLength)
	}

	return data[:length-paddingLength], nil
}
