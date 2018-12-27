package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//输入密钥
	key := "12345678"

	fileName:=""
	file,err:=os.Open(fileName)

	defer file.Close()

	if err!=nil{
		fmt.Println("未找到待处理文件")
		panic(err)
	}
	//读取文件内容
	plain,_:=ioutil.ReadAll(file)
	//创建block
	block,_:=des.NewTripleDESCipher([]byte(key))


	//第三个参数表明是解密
	if len(os.Args)>3{
		DecryptMode:=cipher.NewCBCDecrypter(block,[]byte(key)[:8])
		plain,_ =base64.StdEncoding.DecodeString(string(plain))
		DecryptMode.CryptBlocks(plain,plain)
		plain=PKCS5remove(plain)
		err := ioutil.WriteFile(fileName,plain,0777)
		if err!=nil{
			fmt.Println("保存解密后文件失败!")
		}else{
			fmt.Println("文件已解密!")
		}

	}else{
		EncryptMode:=cipher.NewCBCEncrypter(block,[]byte(key)[:len(key)])
		//明文补足PKCS5Padding
		plain=PKCS5append(plain)
		EncryptMode.CryptBlocks(plain,plain)
		err := ioutil.WriteFile(fileName,[]byte(base64.StdEncoding.EncodeToString(plain)),0777)
		if err!=nil{
			fmt.Println("保存加密后文件失败!")
		}else{
			fmt.Println("文件已加密,务必记住加密key!")
		}
	}
}

func PKCS5append(plaintext []byte) []byte {
	num := 8 - len(plaintext)%8
	for i:=0;i<num;i++{
		plaintext=append(plaintext,byte(num))
	}
	return plaintext
}

func PKCS5remove(plaintext []byte) []byte {
	length := len(plaintext)
	num := int(plaintext[length-1])
	return plaintext[:(length - num)]
}
