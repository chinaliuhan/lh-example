package parser

import (
	"github.com/sirupsen/logrus"
	"lh-example/lh-crawler/cocurrent/engine"
	"regexp"
)

func ParseCityList(contents []byte) engine.ParseResult {
	compile, err := regexp.Compile(`<a href="(.*localhost:8080/mock/www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	if err != nil {
		logrus.Errorf("正则匹配失败")
		return engine.ParseResult{}
	}
	result := engine.ParseResult{}
	match := compile.FindAllSubmatch(contents, -1)
	for _, m := range match {
		//城市名字
		result.Items = append(result.Items, "City "+string(m[2]))
		//城市URL
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
