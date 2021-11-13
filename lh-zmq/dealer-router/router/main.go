package main

import (
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	ROUTER
	可以认为是服务端
	*/
	socket, err := zmq4.NewSocket(zmq4.ROUTER)
	if err != nil {
		log.Fatalln("NewSocket", err)
	}
	//Bind 绑定端口，并指定传输层协议
	err = socket.Bind("tcp://127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Bind", err)
	}
	defer socket.Close()

	for {
		//Send和Recv可以不交替使用
		addr, _ := socket.RecvBytes(0) //接收到的第一帧表示对方的地址
		resp, _ := socket.Recv(0)
		log.Println("addr", addr, "resp", resp)
		_, err := socket.SendBytes(addr, zmq4.SNDMORE) //第一帧需要指明对方的地址,SNDMORE表示消息还没发完
		if err != nil {
			log.Fatalln("SendBytes", err)
		}
		_, err = socket.Send("router send1", zmq4.SNDMORE) //如果不用SNDMORE表示这已经是最后一帧了,下一次Send就是下一段消息的第一帧了
		if err != nil {
			log.Fatalln("Send1", err)
		}
		_, err = socket.Send("router send2", 0) //最后一帧
		if err != nil {
			log.Fatalln("Send2", err)
		}
	}

}
