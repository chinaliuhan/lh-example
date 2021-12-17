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

	//二维数组排序
	type person struct {
		name string
		age  uint8
	}
	var sli []person
	a := person{
		name: "ZhangSan",
		age:  1,
	}
	b := person{
		name: "LisSi",
		age:  3,
	}
	c := person{
		name: "WangEr",
		age:  2,
	}
	d := person{
		name: "MaZi",
		age:  4,
	}
	sli = append(sli, a, b, c, d)
	sort.Slice(sli, func(i, j int) bool {
		//从小到大
		return sli[i].age < sli[j].age
	})
	log.Printf("从小到大 %+v", sli) //从小到大 [{name:ZhangSan age:1} {name:WangEr age:2} {name:LisSi age:3} {name:MaZi age:4}]

	sort.Slice(sli, func(i, j int) bool {
		//从大到大
		return sli[i].age > sli[j].age
	})
	log.Printf("从大到小 %+v", sli) // 从大到小 [{name:MaZi age:4} {name:LisSi age:3} {name:WangEr age:2} {name:ZhangSan age:1}]

}
