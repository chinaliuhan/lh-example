package main

import (
	"crypto/sha1"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	s := "this is sha1"

	//产生一个散列值
	h := sha1.New()

	//写入一个字符串
	h.Write([]byte(s))

	//得到最终的散列值的字符切片
	bs := h.Sum(nil)

	log.Printf("%x", bs)
}
