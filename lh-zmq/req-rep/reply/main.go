package main

import (
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	响应
	*/
	socket, err := zmq4.NewSocket(zmq4.REP)
	if err != nil {
		log.Fatalln("NewSocket", err)
	}

	//绑定端口
	err = socket.Bind("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Bind", err)
	}
	defer socket.Close()

	for {
		//Send和Recv必须交替使用,recv读取之后然后必须send发送数据
		msg, err := socket.Recv(0) //0表示阻塞模式
		if err != nil {
			log.Fatalln("Send", err)
		}
		log.Println("收到", msg)
		send, err := socket.Send("reply send this is reply", 0) //同步发送
		if err != nil {
			log.Fatalln("Send", err)
		}
		log.Println("发送长度", send)
	}

}
