package main

import (
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func simpleWS() {

	wsUrl := "ws://localhost:8081/upgrade"
	origin := "http://localhost/"
	dial, err := websocket.Dial(wsUrl, "ws", origin)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		//发送数据
		write, err := dial.Write([]byte("啊哈哈"))
		if err != nil {
			log.Fatalln(err, write)
		}

		//接收数据
		data := make([]byte, 1024)
		read, err := dial.Read(data)
		if err != nil {
			log.Fatalln(err, read)
		}

		log.Println("客户端接收到数据:", string(data))
		time.Sleep(time.Second * 5)
	}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	simpleWS()

}
