package main

import (
	"bytes"
	"crypto/des"
	"fmt"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//padtext := bytes.Repeat([]byte{0}, padding)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}


// PKCS5Unpadding func
func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


// desEncrypt func
func DESEncryptECB(key []byte, src []byte) []byte {
	desBlockEncrypter, err := des.NewCipher(key[0:8])
	if err != nil {
		panic(err)
	}


	bs := desBlockEncrypter.BlockSize()
	srcPadding := PKCS5Padding(src, bs)
	out := make([]byte, len(srcPadding))
	dst := out
	for len(srcPadding) > 0 {
		desBlockEncrypter.Encrypt(dst, srcPadding[:bs])
		srcPadding = srcPadding[bs:]
		dst = dst[bs:]
	}
	return out
}

// desDecrypt func
func DESDecryptECB(key []byte, src []byte) []byte {
	desBlockDecrypter, err := des.NewCipher(key[0:8])
	if err != nil {
		panic(err)
	}

	bs := desBlockDecrypter.BlockSize()
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		desBlockDecrypter.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}

	out = PKCS5Unpadding(out)

	return out

}

func main(){
	src := "acbd"
	key := "12345678"

	first := DESEncryptECB([]byte(src),[]byte(key))
	fmt.Println(first)

}
