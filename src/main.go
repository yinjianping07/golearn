package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//DES加密函数，src即明文，key即密钥，得到一个密文返回
func DesEnCrypt(src, key []byte) []byte {

	//new一个cipher.block接口，它对应着要加密的块
	block, e := des.NewCipher(key)

	if e != nil {
		fmt.Print(e)
	}

	//Fill函数即为对明文的填充，即，DES加密的明文长度为64位，
	//少于64位可以填充，多余64位可以根据64位一块，形成多个块，不够的填充。
	//block.BlockSize()就是加密的块长度，fill函数会将明文按照块长度进行分组。
	//这样就形成了多个明文分组。以便于进行DES加密
	srcc := Fill(src, block.BlockSize())


	//make一个密文byte切片用以接收
	dst := make([]byte, len(srcc))


	//使用CBC模式进行加密，只需将加密的块，即初始向量传入就可以得到一个CBC模式：BlockMode
	encrypter := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	//使用CBC模式进行加密，并将其赋给dst。
	encrypter.CryptBlocks(dst, srcc)
	return dst
}

//DES解密函数，src即为密文，key为密钥

func DesDecrypt(src, key []byte) []byte {

	//new一个cipher.block接口，它对应着要加密的块
	block, e := des.NewCipher(key)

	if e != nil {
		fmt.Print(e)
	}


	decrypter := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])

	dst := make([]byte, len(src))

	decrypter.CryptBlocks(dst, src)

	out := Out(dst, block.BlockSize())

	return out
}

func main() {
	src := []byte("iamwekadfasafdsfasdf")
	key := []byte("12345679")

	crypt := fmt.Sprintf("%x", DesEnCrypt(src, key))
	fmt.Print(crypt)
	fmt.Println()

	fmt.Printf("%s", DesDecrypt(DesEnCrypt(src, key), key))
	fmt.Println()
}

func Fill(src []byte, blocksize int) []byte {
	fillsize := blocksize - len(src)%blocksize
	repeat := bytes.Repeat([]byte{0}, fillsize)
	return append(src, repeat...)
}

func Out(src []byte, blocksize int) []byte {
	return bytes.TrimRightFunc(src, func(r rune) bool {
		return r == rune(0)
	})
}

