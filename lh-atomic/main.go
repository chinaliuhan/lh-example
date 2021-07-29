package main

import (
	"log"
	"sync/atomic"
)

var num int64

//自增
func incr() {

	atomic.AddInt64(&num, 1)
}

//自减
func decr() {
	atomic.AddInt64(&num, -1)

}

//原子值
func value() {
	/**
	sync/atomic.Value类型的变量一旦被声明，不能被复制到别的地方，虽然不会造成编译错误但是go vet会报告此类不正确，
	因为对结构体的复制会生成该值的副本，还会生成其中字段的副本，导致并发安全保护也就失效惹。甚至向副本存储值的操作与原值都无关了
	但是sync/atomic.Value类型的指针类型的变量不存在这个问题。如果Value里面存储的值是引用类型，那么更要注意了。
	*/
	var val atomic.Value
	/**
	接收一个interface{}类型的参数。
	不能把nil作为参数传入；第二次即以后传入的值必须与之前传入值的类型一致。违反了这两个限制会导致一个运行时恐慌。
	*/
	val.Store([]int{1, 2, 3, 4, 5})

	/**
	返回一个interface{}类型的结果。
	如果还没有Store那么返回的会是nil。
	*/
	log.Println(val.Load())
}

func main() {
	log.SetFlags(log.Lshortfile)
	//调用自增
	incr()
	incr()
	incr()
	//读取,打印
	log.Println(atomic.LoadInt64(&num))

	//调用自减
	decr()
	decr()
	log.Println(atomic.LoadInt64(&num))

	//原子值
	value()

}
