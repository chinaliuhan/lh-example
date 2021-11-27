package parser

import (
	"lh-example/lh-crawler/cocurrent/engine"
	"regexp"
)

var (
	//预编译性能更好,不过must正则如果失败会panic
	profileRe = regexp.MustCompile(`<a href="(.*localhost:8080/mock/album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(.*localhost:8080/mock/www\.zhenai\.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {

	//处理个人信息
	result := engine.ParseResult{}
	match := profileRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		name := string(m[2])
		//城市名字
		result.Items = append(result.Items, "User "+name)
		//城市URL
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				//因为想传入name,所以这里这样用
				return ParseProfile(bytes, name)
			},
		})
	}

	//处理页面中的其他连接
	match = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result

}
