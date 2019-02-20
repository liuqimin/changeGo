package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func Db_passwd_encode(password string) string {
	key := []byte("`?.F(fHbN6XK|j!t")
	result, err := AesEncrypt([]byte(password), key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(result)
}

func Db_passwd_decode(encodePassword string) string {
	key := []byte("`?.F(fHbN6XK|j!t")
	d, _ := base64.StdEncoding.DecodeString(encodePassword)
	origData, err := AesDecrypt(d, key)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	fmt.Println(string(origData))
	return string(origData)
}
func test() {
	key := []byte("0123456789abcdef")
	result, err := AesEncrypt([]byte("hello world"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))

	bb := Db_passwd_encode("lqm11111")
	fmt.Println(bb)
	Db_passwd_decode(bb)
}
