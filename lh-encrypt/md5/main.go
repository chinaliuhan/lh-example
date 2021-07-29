package main

import (
	"crypto/md5"
	"log"
)

//方式1
//16bytes, 128bit
func md51(info []byte) []byte {
	//对多量数据进行哈希运算

	//1. 创建一个哈希器
	hasher := md5.New()
	hasher.Write(info)

	//2. 执行Sum操作，得到哈希值,这里带了点填充
	//hash := hasher.Sum(nil)
	//sum(b), 如果b不是nil， 那么返回的值为b+hash值,即带有填充内容的md5,比如这里填充了0x,会在开头填入四个ASCII码数字3078,形成36位
	hash := hasher.Sum([]byte("0x"))
	//也可以这样返回一个不带填充的md5,不过%s这个是一个16进制的string,而不是byte,Sum(nil)这个不进行填充
	//hash := hasher.Sum(nil)
	//hash := fmt.Sprintf("%x", hasher.Sum(nil))

	return hash
}

//方式2
func md52(info []byte) []byte {
	hash := md5.Sum(info)

	//将数组转换为切片
	return hash[:]
}
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	/**
	可以发现
	方式1 307860c08e2b03b716a176aeb9c2a2fddb79
	方式2 60c08e2b03b716a176aeb9c2a2fddb79
	长度不一致,这是因为方式1前面的四个数是填充的,带了点四个ASCII码数字,如果不想填充就hasher.Sum([]byte("0x")),把[]byte("0x")该为nil即可
	*/

	//方式1
	src := []byte("this is md5")

	md5S1 := md51(src)
	log.Printf("%x", md5S1)

	//方式2
	md5S2 := md52(src)
	log.Printf("%x", md5S2)
}
