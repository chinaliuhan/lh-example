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

/**
生成并保存私钥和公钥
*/
func generateKeyPair(bits int) error {

	/**
	私钥的生成
	*/
	//使用GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA秘钥
	//rand.Reader 随机数生成器 bits 私钥的长度 返回一个私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	//对生成的私钥进行解码处理,x509 按照规则进行序列化处理,生成der编码的数据
	priDerText := x509.MarshalPKCS1PrivateKey(privateKey)

	//创建block代表PEM编码的结构,并填入编码的数据
	blockPri := pem.Block{
		Type:    "RSA Private Key", //得自前言的类型, 随便填写
		Headers: nil,               //可选信息,包括私钥的加密方式等
		Bytes:   priDerText,        //私钥编码后的数据
	}

	//将PEM block数据写入到文件中
	fileHandler, err := os.Create(PRIVATEKEYFILE)
	if err != nil {
		return err
	}
	defer fileHandler.Close()
	err = pem.Encode(fileHandler, &blockPri)
	if err != nil {
		return err
	}

	/**
	公钥的生成
	*/
	//1. 获取公钥， 通过私钥获取
	pubKey := privateKey.PublicKey

	//2. 要对生成的私钥进行编码处理， x509， 按照规则，进行序列化处理, 生成der编码的数据
	pubDerText := x509.MarshalPKCS1PublicKey(&pubKey)

	//3. 创建Block代表PEM编码的结构, 并填入der编码的数据
	blockPub := pem.Block{
		Type:    "RSA Public key",
		Headers: nil,
		Bytes:   pubDerText,
	}

	//4. 将Pem Block数据写入到磁盘文件
	fileHandler1, err := os.Create(PUBLICKEYFILE)
	if err != nil {
		return err
	}
	defer fileHandler1.Close()
	err = pem.Encode(fileHandler1, &blockPub)
	if err != nil {
		return err
	}

	return nil
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//生成私钥
	err := generateKeyPair(2048)
	if err != nil {
		log.Fatalln(err)
	}

}
