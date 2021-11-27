package engine

import (
	"github.com/sirupsen/logrus"
	"lh-example/lh-crawler/cocurrent/fetcher"
	"lh-example/lh-crawler/cocurrent/model"
)

//ConcurrentEngine 内含一个调度器和指定协程数量
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//Run 执行引擎
func (e *ConcurrentEngine) Run(seeds ...Request) {
	//执行传入的调度器
	e.Scheduler.Run()

	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		//创建指定数量的协程,并将worker通过WorkerChan()方法初始化一个channel
		//WorkerChan() =  return make(chan engine.Request)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//将传入的所有request,通过submit逐个推入到requestChan中
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	profileCount := 0
	for {
		//获取worker推出的chan
		result := <-out

		//计数器
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				logrus.Infof("获取到item ###%d %v", profileCount, item)
				profileCount++
			}
		}

		//遍历worker推出的chan里的所有的request,将其通过submit推入到requestChan中
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//URL去重提示
				logrus.Warnln("跳过重复的链接 ##########", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

//createWorker 创建协程
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//将一个空的chan推入到workerChan中,里面的run函数会不断的从中获取
			//里面是chan chan 结构, 会将in推入到另一个chan中,所以下面request := <-in可以接收到里面推送的数据
			ready.WorkerReady(in)

			//调用worker处理传入的request,并将处理结果推入到out中,传出
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//worker  通过传入的r 使用r的url调用下载器下载页面, 再调用r的回调函数
func worker(r Request) (ParseResult, error) {
	logrus.Infoln("读取URL", r.Url)

	//调用下载器,获取指定页面的内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		logrus.Error("fetch失败", err)
		return ParseResult{}, err
	}

	//调用回调函数处理返回的页面
	return r.ParserFunc(body), nil
}

var visitedUrls = make(map[string]bool)

//isDuplicate URL去重
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true

	return false
}
