package engine

//Request 请求的结构体,包含一个URL地址和回调函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//ParseResult 解析的结构体,包含请求列表和返回的节点
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// Scheduler 调度器接口
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}
