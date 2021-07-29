package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

const PRIVATEKEYFILE = "./rsaPrivateKey.pem"
const PUBLICKEYFILE = "./rsaPublicKey.pem"

//公钥加密
func rsaPubEncrypt(filename string, plainText []byte) (error, []byte) {
	//1. 通过公钥文件，读取公钥信息 -> pem encode 的数据
	file, err := os.ReadFile(filename)
	if err != nil {
		return err, nil
	}

	//2. pem decode， 得到block中的der编码数据
	//返回值 ：pem.block
	//返回值：rest参加是未解码完的数据，存储在这里
	block, _ := pem.Decode(file)

	//3. 解码der，得到公钥
	derText := block.Bytes
	publicKey, err := x509.ParsePKCS1PublicKey(derText)
	if err != nil {
		return err, nil
	}
	//4. 公钥加密
	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)

	if err != nil {
		return err, nil
	}
	return nil, cipherData
}

//私钥解密
func rsaPriKeyDecrypt(filename string, cipherData []byte) (error, []byte) {
	//1. 通过私钥文件，读取私钥信息 -> pem encode 的数据
	file, err := os.ReadFile(filename)
	if err != nil {
		return err, nil
	}

	//2. pem decode， 得到block中的der编码数据
	//返回值1 ：pem.block
	//返回值2：rest参加是未解码完的数据，存储在这里
	block, _ := pem.Decode(file)

	//3. 解码der，得到私钥
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)

	if err != nil {
		return err, nil
	}

	//4. 私钥解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherData)

	if err != nil {
		return err, nil
	}

	return nil, plainText

}
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//公钥加密
	src := []byte("this is encrypt content")
	err, cipherData := rsaPubEncrypt(PUBLICKEYFILE, src)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%x", cipherData)

	//私钥解密
	err, plainText := rsaPriKeyDecrypt(PRIVATEKEYFILE, cipherData)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s", plainText)

}
