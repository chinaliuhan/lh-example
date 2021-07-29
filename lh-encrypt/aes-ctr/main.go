package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
)

/*
需求： 使用aes， ctr

aes :
- 分组长度： 16
- 秘钥：16

ctr:
- 不需要填充
- 需要提供一个数字


1. 创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256。
func NewCipher(key []byte) (cipher.Block, error)
- 包：aes
- 秘钥
- cipher.Block接口


2. 选择分组模式：ctr
返回一个计数器模式的、底层采用block生成key流的Stream接口，初始向量iv的长度必须等于block的块尺寸。
func NewCTR(block Block, iv []byte) Stream
- block
- iv
- 秘钥流

3. 加密操作
type Stream interface {
    // 从加密器的key流和src中依次取出字节二者xor后写入dst，src和dst可指向同一内存地址
    XORKeyStream(dst, src []byte)
}

*/

//加密操作aesCTR
func aesCTREncrypt(src, key []byte) []byte {

	//创建一个cipher.Block的接口
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("block size", block.BlockSize())

	//选择分组方式,ctr,准备向量
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	log.Printf("向量为%s", iv)
	stream := cipher.NewCTR(block, iv)

	//加密操作,dst 为结果即密文,src元数据即明文
	stream.XORKeyStream(src, src)

	return src
}

//解密aesCTR
func aesCTRDecrypt(cipherData, key []byte) []byte {
	//创建cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("block size", block.BlockSize())

	//准备向量
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	log.Printf("向量为%s", iv)

	//选择分组模式
	stream := cipher.NewCTR(block, iv)
	//解密操作
	stream.XORKeyStream(cipherData, cipherData)

	return cipherData
}

func main() {

	/**
	des模式一般只用于解密以前的东西,不再用与新数据的加密,因为早就被破解了
	3des因为加密三次,所以会极大的浪费性能
	推荐AES加密
	分组模式选择
	ECB 不要用,不安全,不推荐,效率高
	CBC 推荐,安全
	CRT 推荐,安全而且效率高
	CFB 不推荐,安全
	OFB 不推荐,安全
	*/
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	src := []byte("德艺双磬")
	key := []byte("1234567887654321")

	cipherData := aesCTREncrypt(src, key)
	log.Printf("加密结果为: %x", cipherData)

	plainText := aesCTRDecrypt(src, key)
	log.Printf("解密结果为: %s", plainText)
}
