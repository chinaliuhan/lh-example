package main

import (
	"lh-example/lh-crawler/cocurrent/engine"
	"lh-example/lh-crawler/cocurrent/scheduler"
	"lh-example/lh-crawler/cocurrent/zhenai/parser"
	_ "net/http/pprof" //在初始化导包的时候，pprof 包会自动注册 handler
)

func main() {

	//为引擎注入调度器并指定协程数量
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 1,
	}

	//启动引擎,支持同时传入多个request
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
