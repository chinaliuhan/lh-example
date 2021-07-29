package main

import (
	"fmt"
	"log"
	"time"
)

//计时器
func timer() {
	//等待指定时间
	timer := time.NewTimer(time.Second * 2)
	if true {
		//时间到了之后会向这里推送一个时间数据
		af := <-timer.C
		log.Println("timer", af)

	} else {
		//停止计数器
		timer.Stop()
	}

	log.Println("time1 done")

}

//打点器
func ticker() {
	//每两秒执行一次
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for t := range ticker.C {
			log.Println("ticker", t)
		}
	}()

	time.Sleep(time.Second * 6)
	ticker.Stop()
	log.Println("ticker done...")

}

//超时处理
func after() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "111"
	}()
	select {
	case r := <-c1:
		log.Println("c1", r)
	case af := <-time.After(time.Second * 2):
		log.Println("after", af)
	}
}

func times() {
	//当前时间
	now := time.Now()
	log.Println("当前时间", now)
	log.Println("年", now.Year())
	log.Println("月", now.Month())
	log.Println("日", now.Day())
	log.Println("时", now.Hour())
	log.Println("分", now.Minute())
	log.Println("秒", now.Second())
	log.Println("纳秒", now.Nanosecond())
	log.Println("时区", now.Location())
	log.Println("本地时间", now.Local())
	log.Println("周", now.Weekday())
	log.Println("字符串格式的时间", now.String())
	log.Println("秒时间戳", now.Unix())
	log.Println("纳秒时间戳", now.UnixNano())
	log.Println("将秒时间戳转为时间", time.Unix(now.Unix(), 0))
	log.Println("将纳秒时间戳转为时间", time.Unix(0, now.UnixNano()))

	//指定时间和时区来获取一个时间
	then := time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC)
	log.Println(then)
	diff := then.Sub(now)
	log.Println("默认相差多少,一般是小时开始到秒", diff)
	log.Println("默认相差多少秒", diff.Seconds())

	//用来对比两个时间的顺序
	log.Println(then.Before(now))
	log.Println(then.After(now))
	log.Println(then.Equal(now))

	//格式化时间 如果格式化go生成的时间的话layout可以使用这个常量 time.RFC3339
	t := time.Now()
	log.Println("根据RFC3339来格式化日期将时间格式化为指定的时间格式", t.Format("2006-01-02T15:04:05Z07:00"))
	log.Println("使用Sprintf拼接时间", fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
	//格式化时间,如果是其他时间,比如PHP生成的时间格式,可以使用下面的方式进行处理
	parse, err := time.Parse("2006-01-02 15:04:05", "2019-07-20 18:35:53")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("parse解析时间", parse.String())
	//格式化时间,同时指定时区
	locationTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-07-20 18:35:53", time.Local)
	if err != nil {
		return
	}
	log.Println("ParseInLocation解析时间", locationTime)

	location, err := time.LoadLocation("PRC")
	if err != nil {
		return
	}
	//设置时区格式化时间,go的时区不能全局设置,必须每次都in一下,或者封装起来,LoadLocation在window下不可用
	log.Println("设置时区打印时间", time.Now().In(location).Format(time.RFC3339))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	times()

	//计数器
	//timer()

	//打点器
	//ticker()

	//超时处理
	//after()
	//log.Println("first after ", <-time.After(time.Second*2))

}
