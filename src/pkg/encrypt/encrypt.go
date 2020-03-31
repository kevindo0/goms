package encrypt

import (
	"bytes"
	"crypto/aes"
)

// func main() {
//  data, _ := hex.DecodeString("C42A5FF4DAD4B89353636A137747A4304C40EFA11B6E464BDFE3975248CC218B")
// 	key, _ := hex.DecodeString("88E2EEEEF07E47C93354054AE0608419")
// 	src := "huyong@vsphere.local;D5qXxPaX*T"

//  a := encrypt.DecryptAes128Ecb(data, key)
//  res := PKCS5UnPadding(a)
//  fmt.Println(string(res))

// 	content := []byte(src)
// 	content = PKCS5Padding(content, 16)
// 	a := encrypt.EncryptAes128Ecb(content, key)
// 	fmt.Println(a)
// 	fmt.Println(hex.EncodeToString(a))
// }

func DecryptAes128Ecb(data, key []byte) []byte {
    cipher, _ := aes.NewCipher([]byte(key))
    decrypted := make([]byte, len(data))
    size := 16

    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
        cipher.Decrypt(decrypted[bs:be], data[bs:be])
    }

    return decrypted
}

func EncryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
    encrypted := make([]byte, len(data))
    size := 16
    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
        cipher.Encrypt(encrypted[bs:be], data[bs:be])
    }

    return encrypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    // 去掉最后一个字节 unpadding 次
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}