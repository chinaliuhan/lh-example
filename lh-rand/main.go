package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//如果你出于加密目的，需要使用随机数的话，请使用 crypto/rand 包，此方法不够安全

	//随机数2
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	log.Println("随机数1", r1.Intn(100))

	//随机数1
	rand.Seed(time.Now().UnixNano())
	log.Println("随机数", rand.Intn(100))

	//错误的用法,每次随机的结果都一样,一定要记得给种子
	//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
	s2 := rand.NewSource(100)
	r2 := rand.New(s2)
	log.Println("随机数2", r2.Intn(100))

}
