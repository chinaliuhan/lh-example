package main

import (
	"log"
)

type Person struct {
	name string
	age  uint
}

func delElement1() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//循环中删除切片中指定元素
	for i := 0; i < len(nums); i++ {
		if nums[i] == 5 {
			nums = append(nums[:i], nums[i+1:]...)
			i-- // 后面的元素前移了,所以这里i--
		}
	}

	log.Printf("%+v", nums)

}

func delElement2() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//利用新开辟的空间,将符合条件的保留下来,新的切片,len为0,cap为和原来保持一致
	newNums := make([]int, 0, len(nums))
	for _, n := range nums {
		if n != 5 {
			newNums = append(newNums, n)
		}
	}
	log.Printf("%+v", newNums)
}

func delElement3() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0
	//利用新开辟的空间,将符合条件的保留下来,也就相当于删除了不符合条件的
	newNums := make([]int, len(nums))
	for _, n := range nums {
		if n != 5 {
			newNums[i] = n
			i++
		}
	}
	newNums = newNums[:i]
	log.Printf("%+v", newNums)
}

func delElement4() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0
	//查找符合条件的元素,删除指定元素
	for _, n := range nums {
		if n != 5 {
			nums[i] = n
			i++
		}
	}
	nums = nums[:i]

	log.Printf("%+v", nums)
}

func main() {
	log.SetFlags(log.Lshortfile)
	delElement1()
	delElement2()
	delElement3()
	delElement4()
	/**
	利用append 性能较差,但是容易理解
	修改本slice的性能最好,但是复杂点
	*/
}
