package main

import (
	"log"
	"sort"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//这些排序方法都是针对内置数据类型的 这里的排序方法都是就地排序，也就是说排序改变了切片内容，而不是返回一个新的切片

	//会将字符串按照自然排序,即a->z
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	log.Println(strs)

	//会将整型按照从小到大排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	log.Println(ints)

	//检测切片是否已经排序好
	s := sort.IntsAreSorted(ints)
	log.Println(s)

}
