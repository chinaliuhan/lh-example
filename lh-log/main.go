package main

import (
	"io"
	"log"
	"os"
)

func main() {
	log.Println("打印日志")

	//设置日志前缀, 在日志的首行添加前缀
	log.SetPrefix("### this is prefix ###")
	log.Println("打印日志1")

	/**
	设置日志打印格式
	Ldate         = 1 << iota     // 本地时区的日期: 2009/01/23
	Ltime                         // 本地时区的时间: 01:23:23
	Lmicroseconds                 // 微秒级分辨率:01:23:23.123123。如果设置了Ltime。
	Llongfile                     // 完整文件名和行号:/a/b/c/d.go:23
	Lshortfile                    // 最终文件名+行号:d.go:23。覆盖Llongfile
	LUTC                          // 如果设置了“Ldate”或“Ltime”，请使用UTC而不是当地时区
	Lmsgprefix                    // 将“prefix前缀”从日志的行首移动到消息之前
	LstdFlags     = Ldate | Ltime // 带时间和日期, 默认设置
	*/
	//设置日志格式, 带时间日期,带文件名和行号,将前缀从首行设置到消息前
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	log.Println("打印日志2")

	//设置日志输出到文件, 读写|如果不存在则创建|追加
	f, err := os.OpenFile("./tmp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("打印日志3")

	//log.New()可以让我们创建一个新的 Logger 对象,同时设置他的日志输出位置,前缀和格式,同时
	//Discard 是一个 io.Writer 接口，调用它的 Write 方法将不做任何事情并且始终成功返回。
	trace := log.New(io.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	trace.Println("this is trace")

	info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	info.Println("this is info")

	warning := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	warning.Println("this is waring")

	//io.MultiWriter()可以同时指定输出到两个地方,这里指定的时文件和命令行
	//os.Stderr和io.Stdout差不多都是输出到命令行,但是标准错误文件描述符
	erro := log.New(io.MultiWriter(f, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	erro.Printf("this is error")

	/**
	2021/07/02 13:33:32 打印日志
	### this is prefix ###2021/07/02 13:33:32 打印日志1
	2021/07/02 13:33:32 main.go:30: ### this is prefix ###打印日志2
	INFO: 2021/07/02 13:33:32 main.go:50: this is info
	WARNING: 2021/07/02 13:33:32 main.go:53: this is waring
	ERROR: 2021/07/02 13:33:32 main.go:56: this is error
	*/

}
