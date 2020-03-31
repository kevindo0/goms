package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"flag"
	"fmt"
	"goms/pkg/encrypt"
)

var (
	hexKey = flag.String("key", "", "the key")
	src    = flag.String("src", "", "the content")
)

func main() {
	flag.Parse()

	var hasError = false

	if hexKey == nil || src == nil {
		fmt.Println("args error")
		hasError = true
	}

	key, err := hex.DecodeString(*hexKey)
	if err != nil {
		fmt.Println("key error, please input hex type aes key")
		hasError = true
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error")
		hasError = true
	}

	if *src == "" {
		fmt.Println("plain content empty")
		hasError = true
	}

	if hasError {
		flag.Usage()
		return
	}

	// encrypt
	ecb := encrypt.NewECBEncrypter(block)
	content := []byte(*src)
	content = PKCS5Padding(content, block.BlockSize())
	fmt.Println("content1:", content)
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	fmt.Println("result1:", crypted)
	fmt.Println("result:", hex.EncodeToString(crypted))

	// decrypt
	// ecb := encrypt.NewECBDecrypter(block)
	// content, err := hex.DecodeString(*src)
	// decrypted := make([]byte, len(content))
	// ecb.CryptBlocks(decrypted, content)
	// fmt.Println("result:", string(decrypted))

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}