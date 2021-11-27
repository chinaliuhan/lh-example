package parser

import (
	"github.com/sirupsen/logrus"
	"lh-example/lh-crawler/single/engine"
	"regexp"
)

func ParseCity(contents []byte) engine.ParseResult {
	compile, err := regexp.Compile(`<a href="(.*localhost:8080/mock/album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	if err != nil {
		logrus.Errorf("正则匹配失败")
		return engine.ParseResult{}
	}
	result := engine.ParseResult{}
	match := compile.FindAllSubmatch(contents, -1)
	for _, m := range match {
		name := string(m[2])
		//城市名字
		result.Items = append(result.Items, "User "+name)
		//城市URL
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}

	return result

}
