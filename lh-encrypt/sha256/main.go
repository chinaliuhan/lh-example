package main

import (
	"crypto/sha256"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	src := []byte("this is sha256")

	//创建hash
	hasher := sha256.New()

	//写入数据
	hasher.Write(src)

	//hash sum操作
	hash := hasher.Sum(nil)

	log.Printf("hash : %x\n", hash)
}
