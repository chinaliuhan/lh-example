package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
)

// pack 封包
func pack(message string) ([]byte, error) {
	/**
	为数据加上包头,此时据包就分为包头和包体两部分内容了(如果要过滤非法包时封包会加入包尾即可).
	包头部分的长度是固定的,并且存储了包体的长度,根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包.
	可以自定义一个协议,此处是将数据包的前4个字节为包头,里面存储的是发送的数据的长度.
	*/
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("拨号失败 err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `你好服务器` + strconv.FormatInt(int64(i), 10)
		//为防止tcp粘包,封包
		data, err := pack(msg)
		if err != nil {
			fmt.Println("编码失败 err:", err)
			return
		}
		conn.Write(data)
	}
}
