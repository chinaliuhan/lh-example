package main

import (
	"sync"
	"time"
)

func main() {
	//一般情况下main程序执行完毕后不会等待协程,WaitGroup使main程序等待协程执行完毕后一起退出
	var wait = &sync.WaitGroup{}
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		time.Sleep(time.Second * 2)

		wait.Done()

	}(wait)
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		time.Sleep(time.Second * 2)
		wait.Done()

	}(wait)

	for i := 0; i < 10; i++ {
		//add不能在携程中,必须先行add,只有done可以在携程中
		wait.Add(1)
		go func(wait *sync.WaitGroup) {
			time.Sleep(time.Second * 2)
			wait.Done()
		}(wait)
	}

	wait.Wait()
}
