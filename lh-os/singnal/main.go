package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		sigs = make(chan os.Signal)
		done = make(chan bool)
	)

	// kill -9 {pid} 无条件结束程序（不能被捕获、阻塞或忽略）
	signal.Notify(
		sigs,
		os.Kill,      // 默认: kill {pid}，向进程发送 SIGTERM 信号，可捕获
		os.Interrupt, // 用户发送INTR字符(Ctrl+C 或 kill -INT {pid})触发
		//syscall.SIGHUP,  // 终端控制进程结束(终端连接断开)。使用nohup运行程序时，nohup 仍会将信号发给程序，忽略此信号即可。
		syscall.SIGTERM, // 结束程序(可以被捕获、阻塞或忽略)
		syscall.SIGQUIT, // 用户发送QUIT字符(Ctrl+\)触发
		syscall.SIGTSTP, // 挂起：ctrl + z
		syscall.SIGINT,  //ctrl +c
		syscall.SIGUSR1, //用户保留
		syscall.SIGUSR2, //用户保留
	)
	go func() {
		sig := <-sigs
		log.Println("收到信号,关闭程序: ", sig.String())
		done <- true
	}()
	log.Println("系统被阻塞,等待信号")
	<-done
	log.Println("done...")
}
