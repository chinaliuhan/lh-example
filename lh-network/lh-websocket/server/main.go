package main

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

type WsJson struct {
	Time string `json:"time"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func upgrade(ws *websocket.Conn) {
	//将接收的数据再发送回去
	for {
		var reply string

		//接收字符串
		err := websocket.Message.Receive(ws, &reply)
		if err != nil {
			log.Println(err)
			break
		}

		go func() {
			for {
				websocket.Message.Send(ws, "ping..."+time.Now().String())
				time.Sleep(time.Second * 2)
			}
		}()

		//接收字符串
		err = websocket.Message.Send(ws, "服务器接收到数据: "+reply+time.Now().String())
		if err != nil {
			log.Println(err)
			break
		}
	}

}

func upgradeJson(ws *websocket.Conn) {
	//将接收的数据再发送回去
	for {
		//接收JSON
		wsj := WsJson{}
		err := websocket.JSON.Receive(ws, &wsj)
		if err != nil {
			log.Println(err)
			break
		}
		//发送JSON
		wsj.Time = time.Now().String()
		err = websocket.JSON.Send(ws, wsj)
		if err != nil {
			log.Println(err)
			break
		}
	}

}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.Handle("/upgrade", websocket.Handler(upgrade))
	http.Handle("/upgradeJson", websocket.Handler(upgradeJson))

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
