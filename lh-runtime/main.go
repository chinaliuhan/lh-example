package main

import (
	"fmt"
	"log"
	"runtime"
)

func callerDemo() {
	//可以看到,如果设置为1的话,就是这个文件的这一行
	caller(1) // 17629481   /Users/liuhao/go/src/lh-example/lh-runtime/main.go   10   true   main.callerDemo
}

func caller(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name() //获取函数名
	log.Println(fmt.Sprintf("%v   %s   %d   %t   %s", pc, file, line, ok, pcName))
}

func main() {
	log.SetFlags(log.Lshortfile)

	log.Println("设置最大的可同时使用的CPU核数", runtime.GOMAXPROCS(4))
	log.Println("当前系统的 CPU 核数量", runtime.NumCPU())
	log.Println("正在执行和排队的任务总数", runtime.NumGoroutine())
	log.Println("GoRoot路径", runtime.GOROOT())
	log.Println("当前操作系统", runtime.GOOS)
	log.Println("当前GO版本", runtime.Version())

	//获取代码执行的信息
	//参数：skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
	//pc是uintptr这个返回的是函数指针 file是函数所在文件名目录 line所在行号 ok 是否可以获取到信息
	pc, file, line, ok := runtime.Caller(0)
	//runtime.FuncForPC,获取函数信息,接收的时一个uintptr的指针
	log.Println("获取函数一级调用者", runtime.FuncForPC(pc).Name(), file, line, ok)
	callerDemo()

	//GC执行一次垃圾回收
	runtime.GC()
	//出让时间片,当前线程未来会继续执行
	runtime.Gosched()
	//退出当前 goroutine(但是defer语句会照常执行) 不可再main中执行
	//runtime.Goexit()
}
