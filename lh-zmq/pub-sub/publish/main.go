package main

import (
	"github.com/pebbe/zmq4"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	发布
	*/

	//实例化发布
	context, err := zmq4.NewContext()
	if err != nil {
		log.Fatalln("NewContext error", err)
	}
	socket, err := context.NewSocket(zmq4.PUB)
	defer func(socket *zmq4.Socket) {
		err := socket.Close()
		if err != nil {
			log.Fatalln("close error", err)
		}
	}(socket)
	if err != nil {
		log.Fatalln("NewSocket error", err)
	}
	//设置发布
	err = socket.Bind("tcp://*:8888")
	if err != nil {
		log.Fatalln("bind error", err)
	}

	//不然会丢消息,不光初始连接这里会丢第一条,而且消费者慢了的话也会丢消息,怕丢消息就不要用这个模型
	time.Sleep(time.Millisecond * 200)

	for i := 0; i < 10; i++ {
		_, err = socket.SendMessage(i)
		if err != nil {
			log.Fatalln("SendMessage error", err)
		}
	}
}
