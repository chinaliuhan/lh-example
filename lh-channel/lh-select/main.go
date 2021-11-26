package main

import "log"

//如果intChan1是带bufer的,就会走到case1
//var intChan1 = make(chan int, 10)
var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1, 2, 3, 4, 5}

func getNumber(i int) int {
	log.Println("number", i)
	return numbers[i]
}
func getChan(i int) chan int {
	log.Println("chan", i)
	return channels[i]
}
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	/**
	在开始select语句时,所有跟case关键字右边的发送和接收语句的表达式和元素表达式都会先求值(从左到右,从上到下),无论该case时候会被选中
	下面前四个打印的结果展示了表达式的求值顺序
	因为intChan1和intChan2都没有被初始化,向他们发送数据会被阻塞,所以所有的case的走不通,只会走default.
	select在开始时会(从左到右,从上到下)判断每个发送和接收是否可以立即执行,且不会因为该操作而阻塞goroutine,该判断还要根据通道的具体特性(缓冲,非缓冲)以及那一刻的具体情况来进行.
	如果存在通道值为nil的读写操作，则该分支将被忽略
	如果intChan1是带bufer的就会走到case1,否则只能走到default
	*/

	select {
	case getChan(0) <- getNumber(0):
		log.Println("case 1")
	case getChan(1) <- getNumber(1):
		log.Println("case 2")
	default:
		log.Println("default")
	}

	log.Println(channels) //如果intChan1是带bufer的,就会走到case1  这里就会返回[0xc0000c2000 <nil>]

	/**
	如果intChan1和intChan2都没有被初始化,向他们发送数据会被阻塞,所以所有的case的走不通,只会走default.
	chan 0
	number 0
	chan 1
	number 1
	default
	[<nil> <nil>]
	*/
}
