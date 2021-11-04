package main

import (
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8888,
	})
	if err != nil {
		log.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("你好,这里是客户端")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		log.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		log.Println("接收数据失败，err:", err)
		return
	}
	log.Printf("收到:%v 地址:%v 字节数:%v\n", string(data[:n]), remoteAddr, n)
}
