package main

import (
	"lh-example/lh-crawler/single/engine"
	"lh-example/lh-crawler/single/zhenai/parser"
)

func main() {

	url := "http://localhost:8080/mock/www.zhenai.com/zhenghun"
	engine.Run(
		engine.Request{
			Url:        url,
			ParserFunc: parser.ParseCityList,
		})

}
