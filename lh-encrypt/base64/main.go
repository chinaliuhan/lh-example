package main

import (
	"encoding/base64"
	"log"
)

func base64Encode() {

	info := []byte("https://baidu.com/user/login?name=peter")

	//将字符串转为base64
	encode := base64.StdEncoding.EncodeToString(info)
	log.Println("base64 encode", encode)
	//将base64转码为字符串
	deStr, _ := base64.StdEncoding.DecodeString(encode)
	log.Println("base64 decode", string(deStr))

	//Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	urlEncode := base64.URLEncoding.EncodeToString(info)
	log.Println("url base64 encode", urlEncode)
	uDec, _ := base64.URLEncoding.DecodeString(urlEncode)
	log.Println("url base64 decode ", string(uDec))

}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	/**
	base64不是加密,是编码,放在这里是因为汇总而已
	*/

	base64Encode()
}
