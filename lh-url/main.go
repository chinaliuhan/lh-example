package main

import (
	"log"
	"net/url"
)

func parseUrl() {

	urlPath := "https://user:pass@www.baidu.com:80/user/login?name=peter&age=19&sex=1#aaa"
	parseUrl, err := url.Parse(urlPath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("登录信息账号:", parseUrl.User.Username())
	password, _ := parseUrl.User.Password()
	log.Println("登录信息的密码:", password)
	log.Println("协议:", parseUrl.Scheme)
	log.Println("Host:", parseUrl.Host)
	log.Println("Path:", parseUrl.Path)
	log.Println("锚点:", parseUrl.Fragment)
	log.Println("Get参数:", parseUrl.RawQuery)

	//将URL地址中的get参数解析
	query, err := url.ParseQuery(parseUrl.RawQuery)
	if err != nil {
		log.Fatalln(query)
	}
	log.Println("将GET参数解析为Map:", query)

}

func urlEncode1() {
	//URL encode
	values := url.Values{}
	values.Set("name", "peter")
	values.Set("age", "19")
	values.Set("sex", "1")
	values.Set("addr", "PRC")
	values.Set("CNName", "拉布拉多国王")
	log.Println(values.Encode())
}
func urlEncode2() {
	str := "name=peter&age=19&sex=1&addr=PRC&common=傻了吧,人家会飞!"
	//encode,会对字符串进行转义，使其能够安全地放置在URL查询中
	en := url.QueryEscape(str)
	log.Println(en)

	//decode
	log.Println(url.QueryUnescape(en))
}

func main() {
	//解析URL地址
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	parseUrl()

	urlEncode1()

	urlEncode2()

}
