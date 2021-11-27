package scheduler

import "lh-example/lh-crawler/cocurrent/engine"

//QueuedScheduler 实现了Scheduler接口,并一个请求chan和协程chan
type QueuedScheduler struct {
	requestChan chan engine.Request      //里面是一个URL和回调函数
	workerChan  chan chan engine.Request //同上,不过又被一个chan包裹,利用chan嵌套chan 实现函数异步执行 顺序返回值
}

//WorkerChan 初始化一个request chan
func (qs *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

//Submit 将request推入到 requestChan 中
func (qs *QueuedScheduler) Submit(r engine.Request) {
	qs.requestChan <- r
}

//WorkerReady 将传入的chan request 推入到workerChan中,以供run函数消费
func (qs *QueuedScheduler) WorkerReady(w chan engine.Request) {
	qs.workerChan <- w
}

//Run 执行调度器
func (qs *QueuedScheduler) Run() {
	//初始化qs.workerChan qs.requestChan,方便被引擎调用
	qs.workerChan = make(chan chan engine.Request)
	qs.requestChan = make(chan engine.Request)

	go func() {
		/**
		1. 通过select机制,当有新的新的request推入到requestChan或worker推入到workerChan时
		2. 就会将requestChan和workerChan分别放入到上面的workerQ,requestQ队列中
		3. 再不断的将队列中的 requestQ和workerQ 取出 通过select的左右同时准备好即可运行的机制,将request赋值给activeWorker的worker来执行
		*/

		//初始化请求队列和worker队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request     //激活的请求
			var activeWorker chan engine.Request //激活的worker

			//如果同时有请求和worker,就将两者的第一个元素放入到激活的变量中取
			if len(requestQ) > 0 && len(workerQ) > 0 {
				//准备好一个worker,准备好一个request, 这样下面的activeWorker <- activeRequest才能正常执行,运行activeWorker中的work
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-qs.requestChan:
				//如果有request,就从requestChan中获取一个请求,放入到请求队列中
				requestQ = append(requestQ, r)
			case w := <-qs.workerChan:
				//如果有worker,就从workerChan中获取一个worker,放入到worker队列中
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				//同时又worker和request, 就将激活的请求,放入到激活的worker中开始工作,并删除worker和request队列,的第一个元素
				//这里是利用了select的机制,当左右两边同时准备好且不为nil,并且接收方是已被初始化过的时就会走这个区间,并将activeRequest赋值给worker里的函数来执行
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
