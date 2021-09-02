package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	/**
	NewReader
	*/

	s := strings.NewReader("this is bufio 1 ")
	// NewReaderSize 将 rd 封装成一个拥有 size 大小缓存的 bufio.Reader 对象
	// 如果 rd 的基类型就是 bufio.Reader 类型，而且拥有足够的缓存
	// 则直接将 rd 转换为基类型并返回
	//bufio.NewReaderSize(nr, 100)

	// NewReader 相当于 NewReaderSize(rd, 4096)
	reader := bufio.NewReader(s)
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 字节数据
	// 该操作不会将数据读出，只是引用
	// 引用的数据在下一次读取操作之前是有效的
	// 如果引用的数据长度小于 n，则返回一个错误信息
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull
	// 通过 Peek 的返回值，可以修改缓存中的数据
	// 但是不能修改底层 io.Reader 中的数据
	b, _ := reader.Peek(100)
	log.Printf("%s", b)

	s2 := strings.NewReader("this is bufio 2")
	br := bufio.NewReader(s2)
	b2 := make([]byte, 20)
	// Read 读出数据到 p 中，返回读出的字节数
	_, _ = br.Read(b2)
	log.Printf("%s", b2)

	s3 := strings.NewReader("this is bufio 3")
	br3 := bufio.NewReader(s3)

	// ReadByte 读出一个字节并返回
	// 如果 无可读数据，则返回一个错误
	//还有ReadRune 什么的就不在写了,套路都一样
	bb, _ := br3.ReadByte()
	log.Printf("%c", bb)

	/**
	NewWriter
	*/
	bw := bytes.NewBuffer(make([]byte, 0))
	ww := bufio.NewWriter(bw)
	//写入字符串
	ww.WriteString("this")
	ww.WriteString(" is")
	ww.WriteString(" bufio")
	ww.WriteString(" new writer")
	// Flush 将缓存中的数据提交到底层的 io.Writer 中
	ww.Flush()
	fmt.Printf("%s", bw)

	/**
	NewScanner
	*/

	// NewScanner 创建一个 Scanner 来扫描 r,默认匹配函数为 ScanLines。
	input := bufio.NewScanner(os.Stdin)
	//逐行遍历输入的内容
	// Scan 在遇到下面的情况时会终止扫描并返回 false（扫描一旦终止，将无法再继续）：
	// 1、遇到 io.EOF
	// 2、遇到读写错误
	// 3、“匹配部分”的长度超过了缓存的长度
	for input.Scan() {
		//将最后一次扫描出的“匹配部分”作为字符串返回（返回副本）
		text := input.Text()
		log.Println(text)
		// Bytes 将最后一次扫描出的“匹配部分”作为一个切片引用返回，下一次的 Scan 操作会覆
		bb := input.Bytes()
		log.Println(bb)
	}
	// Err 返回扫描过程中遇到的非 EOF 错误，供用户调用，以便获取错误信息。
	if err := input.Err(); err != nil {
		log.Println(err)
	}

}
