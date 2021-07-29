package main

import (
	"encoding/base64"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	str := "this is a good person"
	//将字符串转为base64
	enStr := base64.StdEncoding.EncodeToString([]byte(str))
	log.Println()
	//将base64转码为字符串
	deStr, _ := base64.StdEncoding.DecodeString(enStr)
	log.Println(string(deStr))

	//Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	uEnc := base64.URLEncoding.EncodeToString([]byte(str))
	log.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	log.Println(string(uDec))

}
