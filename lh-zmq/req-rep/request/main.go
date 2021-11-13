package main

import (
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	请求
	*/
	socket, _ := zmq4.NewSocket(zmq4.REQ)
	//发送连接
	err := socket.Connect("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Connect", err)
	}
	defer socket.Close()

	for i := 0; i < 10; i++ {
		//Send和Recv必须交替使用,也就是send发送之后必须recv读取数据
		send, err := socket.Send("req send: this is request "+string(rune(i)), zmq4.DONTWAIT) //非阻塞模式,只是将数据写入本地buffer,并没有真正发送
		if err != nil {
			log.Fatalln("Send", err)
		}
		log.Println("发送长度", send)
		msg, err := socket.Recv(0) //接收消息
		if err != nil {
			log.Fatalln("Recv", err)
		}
		log.Println("收到", msg)
	}

}
