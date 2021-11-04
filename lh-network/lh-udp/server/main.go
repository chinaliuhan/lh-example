package main

import (
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8888,
	})
	if err != nil {
		log.Println("监听失败 err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			log.Println("读取UDP失败 err:", err)
			continue
		}
		log.Printf("收到:%v 客户端地址:%v 字节数:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			log.Println("写入UDP失败 err:", err)
			continue
		}
	}
}
