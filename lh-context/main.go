package main

import (
	"context"
	"log"
	"time"
)

func do(ctx context.Context) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			//收到取消,打印信息
			log.Println("canceled")
			return
		default:
			//未收到取消,持续工作
			log.Println("doing")
		}
	}
}

func withCancel(rootContext context.Context) {
	// 新建一个可取消的ctx
	ctx, cancel := context.WithCancel(rootContext)

	//调用协程
	go do(ctx)

	time.Sleep(3 * time.Second)

	//调用context.WithCancel 返回的CancelFunc
	cancel()

	time.Sleep(2 * time.Second)
	log.Println("end")

}

func withValue(rootContext context.Context) {
	ctx := context.WithValue(rootContext, "val00", "this is value00")
	value01(ctx)
}
func value01(ctx context.Context) {
	ctx01 := context.WithValue(ctx, "val01", "this is value01")
	value02(ctx01)
}
func value02(ctx context.Context) {
	log.Println(ctx.Value("val00"))
	log.Println(ctx.Value("val01"))
}

func main() {
	log.SetFlags(log.Lshortfile)

	//background和todo其实是一样的
	rootContext := context.TODO()

	withCancel(rootContext)

	withValue(rootContext)

}
