package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"log"
)

/*
需求：算法：des ， 分组模式：CBC

des :
秘钥：8bytes
分组长度：8bytes

cbc:
1. 提供初始化向量，长度与分组长度相同，8bytes
2. 需要填充


加密分析

1. 创建并返回一个使用DES算法的cipher.Block接口。

	func NewCipher(key []byte) (cipher.Block, error)
	- 包名：des
	- 参数：秘钥，8bytes
	- 返回值：一个cipher.Block接口

	type Block interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Encrypt(dst, src []byte)
		// 解密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Decrypt(dst, src []byte)
	}

2. 进行数据填充


3. 引入CBC模式, 返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。
	func NewCBCEncrypter(b Block, iv []byte) BlockMode
	- 包名：cipher
	- 参数1：cipher.Block
	- 参数2：iv， initialize vector
	- 返回值：分组模式，里面提供加解密方法

	type BlockMode interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
		CryptBlocks(dst, src []byte)
	}

解密分析
1. 创建并返回一个使用DES算法的cipher.Block接口。

	func NewCipher(key []byte) (cipher.Block, error)
	- 包名：des
	- 参数：秘钥，8bytes
	- 返回值：一个cipher.Block接口

	type Block interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Encrypt(dst, src []byte)
		// 解密src的第一块数据并写入dst，src和dst可指向同一内存地址
		Decrypt(dst, src []byte)
	}


2. 返回一个密码分组链接模式的、底层用b解密的BlockMode接口，初始向量iv必须和加密时使用的iv相同。
	func NewCBCDecrypter(b Block, iv []byte) BlockMode
	- 包名：cipher
	- 参数1：cipher.Block
	- 参数2：iv， initialize vector
	- 返回值：分组模式，里面提供加解密方法

	type BlockMode interface {
		// 返回加密字节块的大小
		BlockSize() int
		// 加密或解密连续的数据块，src的尺寸必须是块大小的整数倍，src和dst可指向同一内存地址
		CryptBlocks(dst, src []byte)
	}

3. 解密操作

4. 去除填充
*/

//加密,输入明文和秘钥, des加密key最小为8位数
func desCBCEncrypt(src, key []byte) []byte {
	log.Printf("开始加密,输入数据为%s", src)

	//创建并返回一个des算法的cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("block size", block.BlockSize())

	//进行数据填充,如果不填充的话只能对8个字符长度的数据进行加密
	src = paddingInfo(src, block.BlockSize())

	//引入cbc模式,返回一个秘密分组链接模式,底层用b密码的blockModel接口,初始化向量iv的长度必须等于b的块尺寸,一般des是8位,aes是16位
	iv := bytes.Repeat([]byte("1"), block.BlockSize()) //将1自动追加block.BlockSize()次,这里也就是8个1
	log.Println("向量为", string(iv))
	blockModel := cipher.NewCBCEncrypter(block, iv)

	//加密操作,dst是加密后的秘闻,src是加密前的明文
	blockModel.CryptBlocks(src, src)

	log.Printf("加密结束,加密数据为%x", src)
	return src
}

//填充函数,输入明文,分组长度,输出填充后的数据
func paddingInfo(src []byte, blockSize int) []byte {
	//明文长度
	length := len(src)
	//要填充的数量
	remains := length % blockSize     //3
	paddingNum := blockSize - remains //5
	//把填充的数值转为字符
	s1 := byte(paddingNum) //'5'
	//把字符转为数组
	s2 := bytes.Repeat([]byte{s1}, paddingNum) //多次填充的结果为[]byte{'5','5','5','5','5'}
	//把拼成的数组住家到src后面
	newSrc := append(src, s2...)
	//返回新的数组
	return newSrc
}

//解密,输入秘闻和秘钥, des加密key最小为8位数
func desCBCDecrypt(cipherData, key []byte) []byte {
	log.Printf("开始解密,输入数据为%x", cipherData)

	//创建并返回一个des算法的cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("block size", block.BlockSize())

	//进行数据填充,如果不填充的话只能对8个字符长度的数据进行加密
	//cipherData = paddingInfo(cipherData, block.BlockSize())

	//引入cbc模式,返回一个秘密分组链接模式,底层用b密码的blockModel接口,初始化向量iv的长度必须等于b的块尺寸,一般des是8位,aes是16位
	iv := bytes.Repeat([]byte("1"), block.BlockSize()) //将1自动追加block.BlockSize()次,这里也就是8个1
	log.Println("向量为", string(iv))
	blockModel := cipher.NewCBCDecrypter(block, iv)

	//解密操作,dst是解密后的明文,src是解密前的密文
	blockModel.CryptBlocks(cipherData, cipherData)

	log.Printf("解密结束,加密数据为%x", cipherData)

	//去除填充
	plainText := unPaddingInfo(cipherData) //此时的cipherData已是解密后的明文了

	return plainText
}

//去除填充
func unPaddingInfo(plainText []byte) []byte {
	//获取长度
	length := len(plainText)
	if length == 0 {
		log.Fatalln("去除填充的长度为0")
	}
	//获取最后一个字符
	lastByte := plainText[length-1]

	//将字符转为数字
	unPaddingNum := int(lastByte)

	//切片获取需要的数据
	return plainText[:length-unPaddingNum]
}

func main() {

	/**
	des模式一般只用于解密以前的东西,不再用与新数据的加密,因为早就被破解了
	3des因为加密三次,所以会极大的浪费性能
	推荐AES加密
	分组模式选择
	ECB 不要用
	CBC 推荐
	CRT 推荐
	CFB
	OFB
	*/
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	src := []byte("德艺双磬")
	key := []byte("12345678")

	//des加密
	cipherDta := desCBCEncrypt(src, key)
	log.Printf("%x", cipherDta)

	log.Println("##########################")

	//des解密
	plainText := desCBCDecrypt(cipherDta, key)
	log.Printf("%s", plainText)
	log.Printf("%x", plainText)
}
