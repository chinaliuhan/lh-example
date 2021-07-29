package main

import (
	"log"
	"time"
)

// 我们将在worker函数里面运行几个并行实例，这个函数从jobs通道
// 里面接受任务，然后把运行结果发送到results通道。每个job我们
// 都休眠一会儿，来模拟一个耗时任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		log.Println("worker", id, "processing job", i)
		//模拟程序处理
		time.Sleep(time.Second * 2)
		results <- i * 2
	}

}

func workerPool() {
	// 要使用工作池，需要发送工作和接受工作的结果， 这里我们定义两个通道，一个jobs，一个results
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 这里启动3个worker协程，一开始的时候worker阻塞执行，因为jobs通道里面还没有工作任务
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	// 发送9个任务，然后关闭通道，告知任务发送完成
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	// 从results里面获得结果
	for k := 1; k <= 9; k++ {
		<-results
	}

}

func easyPool() {

	for i := 0; i < 10; i++ {
		go func(id int) {
			log.Println("this is process id: ", id)
			time.Sleep(time.Second * 2)
		}(i)
	}
}
func main() {
	/**
	工作池最大的作用是为了限制协程无限制的增长,虽然GO的协程开销很小,但是要知道每个携程中是带有很多业务逻辑的,如果协程在执行中被堵住之类的,导致协程一直增长,很容易拖垮服务器
	可以看到,虽然每个协程都休眠了两秒,但是实际上这里一共也只休眠了四秒,可见协程是并行了
	*/
	//建议工作池
	easyPool()
	//工作池
	workerPool()
}
