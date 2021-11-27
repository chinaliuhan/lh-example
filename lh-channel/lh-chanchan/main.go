package main

import (
	"log"
	"time"
)

type chanchanDemo struct {
	cc chan chan string
}

//把chan作为参数使用
func (r *chanchanDemo) in(in chan string) {
	//将接收到的in 参数chan,推入到cc chan 中
	r.cc <- in
}

func (r *chanchanDemo) out() {
	//取出cc chan
	c := <-r.cc

	//因为 cc 的结构是chan chan 取出的结果依然可以推入数据
	//UTC时间
	//c <- time.Now().In(time.FixedZone("CST", 0)).String()
	//东八区时间
	c <- time.Now().In(time.FixedZone("CST", 8*3600)).String()

}

func main() {
	/**
	chan chan 嵌套结构
	*/
	//实例化,并初始化 channel
	ccd := chanchanDemo{}
	ccd.cc = make(chan chan string)

	//将该chan 作为参数传入
	in := make(chan string)

	for {

		go ccd.in(in)

		go ccd.out()

		log.Println(<-in)
		time.Sleep(time.Second * 2)
	}

}
