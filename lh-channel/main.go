package main

import (
	"log"
	"time"
)

func blockChan() {

	//一个不带容量的channel,必须同时拥有生产者和消费者且都准备完毕,否则,如果只有消费者的话,会被阻塞在消费者那一行
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "ping"
	}()
	log.Println(<-ch)
}

func bufferChannel() {
	//一个带有容量的channel,长度为10,特性就是即使没有消费者,只要buffer还没有满,就可以继续生产,但是如果一旦超过长度则会报错
	ch := make(chan string, 2)
	ch <- "first"
	ch <- "second"
	//ch <- "third"//容量为2,这里继续推送则会超出容量,会报错
	log.Println(<-ch)
	log.Println(<-ch)
}

func worker() {
	ch := make(chan string, 1)
	go func() {
		//将处理后的结果推送到channel中和外面交互,一般不会写在一起,这里是因为函数多,为了方便
		log.Println("start...")
		time.Sleep(time.Second * 2)
		log.Println("end...")
		ch <- "done..."
	}()
	log.Println(<-ch)
}

func direction() {
	//带有方向的channel,这类channel,只接受指定方向的channel,比如ch <-只接受推送,<-chan 只接受生成

	//ping 函数定义了一个只允许发送数据的通道。尝试使用这个通道来接收数据将会得到一个编译时错误。
	ping := func(pings chan<- string, msg string) {
		pings <- msg
	}
	//pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。
	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "this is message")
	pong(pings, pongs)
	log.Println(<-pongs)
}

func selectChan() {
	//通道channel的选择器
	ch1 := make(chan string)
	ch2 := make(chan string)

	//通过协程向channel推送数据
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "two"
	}()

	//同样,如果这个i的遍历此处超过2也会报错,deadlock,如果要避免的话,可以close掉该channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			log.Println("received", msg1)
		case msg2 := <-ch2:
			log.Println("received", msg2)
		}
	}
}

func forChannel() {
	//循环遍历channel

	//声明一个buffer channel,向该channel不断推送如数据,后面使用for range 不断的遍历该channel中的数据
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//如果这里不close掉该channel,将在这个循环中继续阻塞执行继续等待接收
		close(ch)
	}()

	for j := range ch {
		log.Println(j)
	}
}

func timeOutChannel() {
	//声明一个buffer channel,向该channel推送数据,但是休眠2秒,
	//该select获取channel时,如果1秒内读取不到数据则不再读取,打印timeout 1
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "111"
	}()
	select {
	case res := <-c1:
		log.Println(res)
	case <-time.After(time.Second * 1):
		log.Println("timeout 1")
	}

	//和上面的类似,不过这里的sleep和timeout设置比较合理,让他不会出现timeout
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "222"
	}()
	select {
	case res := <-c2:
		log.Println(res)
	case <-time.After(time.Second * 3):
		log.Println("timeout 2")
	}

}

func unblockChannel() {
	//常规的通过通道发送和接收数据是阻塞的。但是可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。

	//这里是一个非阻塞接收的例子。如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。如果不是，就直接到 default 分支中。
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		log.Println("received message", msg)
	default:
		log.Println("no message received")
	}

	//一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	select {
	case messages <- msg:
		log.Println("sent message", msg)
	default:
		log.Println("no message sent")
	}
	//我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作。
	select {
	case msg := <-messages:
		log.Println("received message", msg)
	case sig := <-signals:
		log.Println("received signal", sig)
	default:
		log.Println("no activity")
	}
}

func closeChannel() {
	//关闭 一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这个通道的接收方传达工作已经完成的信息。

	//在这个例子中，我们将使用一个 jobs 通道来传递 main() 中 Go协程任务执行的结束信息到一个工作 Go 协程中。当我们没有多余的任务给这个工作 Go 协程时，我们将 close 这个 jobs 通道。
	jobs := make(chan int, 5)
	done := make(chan bool)

	//这是工作 Go 协程。使用 j, more := <- jobs 循环的从jobs 接收数据。在接收的这个特殊的二值形式的值中，如果 jobs 已经关闭了，并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。当我们完成所有的任务时，将使用这个特性通过 done 通道去进行通知。
	go func() {
		for {
			j, more := <-jobs
			if more {
				log.Println("received job", j)
			} else {
				log.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//这里使用 jobs 发送 3 个任务到工作函数中，然后关闭 jobs。
	for j := 1; j <= 3; j++ {
		jobs <- j
		log.Println("sent job", j)
	}
	close(jobs)
	log.Println("sent all jobs")

	<-done

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//声明一个unbuffer channel,向channel推入数据,而后消费掉
	//blockChan()

	//声明一个buffer channel ,想channel推入数据,而后逐一消费
	//bufferChannel()

	//传入一个channel,经过函数处理后再传出,这个东西很有用,比如对一个数据进行不同流程的处理,对一个函数处理完之后传入第二个函数,以此类推
	//worker()

	//当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值。这个特性提升了程序的类型安全性。
	//direction()

	//channel选择器
	//selectChan()

	//遍历通道
	//forChannel()

	//channel的超时处理
	//timeOutChannel()

	//不阻塞的channel
	//unblockChannel()

	//关闭通道
	//closeChannel()
}
