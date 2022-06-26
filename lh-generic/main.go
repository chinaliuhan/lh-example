package main

import (
	"log"
)

/**
不使用泛型
*/
func sumNoGenericInt(a, b int) int {
	//只接受int参数
	return a + b
}
func sumNoGenericFloat64(a, b float64) float64 {
	//只接受float64参数
	return a + b
}
func sumNoGenericInterface(a, b interface{}) interface{} {
	//需要断言
	aint, oka := a.(int)
	bint, okb := b.(int)
	if oka && okb {
		return aint + bint
	}

	return nil
}

/**
使用泛型
*/
func sumGeneric[T int | float64](a, b T) T {
	//只接受int或float64

	return a + b
}
func sumGenericInts[T ~int](a, b T) T {
	//~int表示只接受int,int32,int64,uint 等等,整型家族衍生类型

	return a + b
}

func sumGenericComparable[K comparable, V ~int | string](m map[K]V) V {
	//comparable: 表示go里面所有内置的可比较类型：int、uint、float、bool、struct、指针等一切可以比较的类型
	//comparable不能是一个变量,所以这里把他作为一个map的下标使用
	var c V
	for _, v := range m {

		c += v
	}

	return c
}

/**
泛型类型约束
*/
type Number interface {
	~int | ~float64
}

func sumNumbers[T Number](a, b T) T {
	return a + b
}

/**
约束
*/

/**
泛型
version >=1.18
1.17需要单独配置才能开启泛型
*/
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//会自动推断返回类型,这里[]中的int可以不写
	sumRes := sumGeneric[int](1, 2)
	log.Println(sumRes)
	sumRes1 := sumGeneric(1, 2)
	log.Println(sumRes1)

	//可以是整数也可以是浮点数
	sumNumbersRes := sumNumbers(1, 2)
	log.Println(sumNumbersRes)
	sumNumbersRes1 := sumNumbers(1.1, 2.2)
	log.Println(sumNumbersRes1)

}
