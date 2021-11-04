package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

// pack 解包
func unpack(reader *bufio.Reader) (string, error) {
	/**
	和封包的理论相同
	*/
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		//防止tcp粘包,解包
		msg, err := unpack(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("解包失败 err:", err)
			return
		}
		log.Println("收到：", msg)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("监听失败 err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("接收失败 err:", err)
			continue
		}
		go process(conn)
	}
}
