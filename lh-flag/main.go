package main

import (
	"flag"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	/**
	启动时使用-h命令会根据usage进行提示
	 go run ./main.go -h
	Usage of /var/folders/js/sqmwsj913q9_cbv032mnc2vm0000gn/T/go-build2530572463/b001/exe/main:
	  -conf string
	        配置文件地址
	  -debug
	        debug模式

	启动时通过-conf=/path -debug=true 来设置参数
	*/

	//指定参数名和默认值,以及提示信息,注意返回的是指针, flag比os.Args稍微人性化一点
	confPath := flag.String("conf", "", "配置文件地址")
	debugMode := flag.Bool("debug", false, "debug模式")
	//解析命令
	flag.Parse()

	//打印信息
	log.Println(*confPath)
	log.Println(*debugMode)

}
