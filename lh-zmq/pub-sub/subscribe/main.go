package main

import (
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	订阅
	*/

	//实例化订阅
	socket, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		log.Fatalln("NewSocket error", err)
	}
	//设置订阅
	err = socket.SetSubscribe("")
	if err != nil {
		log.Fatalln("SetSubscribe error", err)
	}
	//连接
	err = socket.Connect("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Connect error", err)
	}

	for {
		//循环读取,不断打印
		//zmq4.DONTWAIT　　对于当socket不可使用就要执行阻塞方式的socket类型来说（DEALER，PUSH），此选项可以指定这个操作以非阻塞模式执行。如果无法添加消息到socket的消息队列上，zmq_send()函数将会执行失败并设置errno为EAGAIN。
		//这里使用0,默认阻塞
		msg, err := socket.Recv(0)
		if err != nil {
			log.Fatalln("Recv error", err)
		}
		log.Println(msg)
	}
}
