package main

import (
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	/**
	DEALER
	可以认为是客户端
	*/

	socket, err := zmq4.NewSocket(zmq4.DEALER)
	if err != nil {
		log.Fatalln("NewSocket", err)
	}
	//建立连接
	err = socket.Connect("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Connect", err)
	}
	defer socket.Close()

	for i := 0; i < 10; i++ {
		//Send和Recv可以不交替使用
		_, err = socket.Send("dealer send", 0) //非阻塞模式,只是将数据写入本地buffer,并没有真正发送
		if err != nil {
			log.Fatalln("Send", err)
		}
		resp1, err := socket.Recv(0)
		if err != nil {
			log.Fatalln("Recv1", err)
		}
		resp2, err := socket.Recv(0)
		if err != nil {
			log.Fatalln("Recv2", err)
		}
		log.Println("recv1", resp1, "recv2", resp2)
	}
}
