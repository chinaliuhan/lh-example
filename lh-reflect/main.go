package main

import (
	"log"
	"reflect"
)

type Demo struct {
	Name string `json:"name,omitempty"` //omitempty 在使用时如果不给值则在json中不显示该字段
	Sex  int    `json:"sex,omitempty"`
	Age  int    `json:"age,omitempty"`
	Addr string `json:"addr,omitempty"`
}

func (receiver *Demo) One() string {

	return "one"
}
func (receiver *Demo) Two() string {

	return "two"
}

func valueOf(v reflect.Value) {
	log.Println("###########ValueOf##############")

	//当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)。
	//v := reflect.ValueOf(d)
	log.Println("数据内容", v.Type())
	log.Println("数据内容", v)
	log.Println("字符串形式", v.String())
	log.Println("类别", v.Kind())
	if v.Kind() == reflect.Ptr {
		log.Println("指针类型的所有值", v.Elem()) //仅限指针类型,否则会panic
	}
	if v.Kind() == reflect.Struct {
		log.Println("字段数量", v.NumField()) //仅限首字母大写的函数名,也就是public级别的
		if v.NumField() > 0 {
			log.Println("通过字段名称获取字段的值", v.FieldByName("Name")) //仅限首字母大写的函数名,也就是public级别的
			log.Println("通过字段编号获取字段的值", v.Field(0))
			//遍历字段
			for i := 0; i < v.NumField(); i++ {
				log.Printf("###%v", v.Field(i))
			}
		}

	}
	log.Println("是否可以获取地址", v.CanAddr())
	log.Println("是否可以修改", v.CanSet())
	if v.CanAddr() {
		log.Println("内容地址", v.Addr())
		log.Println("不安全的地址", v.UnsafeAddr())
	}
	log.Println("方法数量", v.NumMethod()) //仅限首字母大写的函数名,也就是public级别的
	if v.NumMethod() > 0 {
		log.Println("通过名称获取方法", v.MethodByName("One").String()) //仅限首字母大写的函数名,也就是public级别的
		log.Println("通过编号获取方法", v.Method(0).String())           //仅限首字母大写的函数名,也就是public级别的
		// 遍历方法
		for i := 0; i < v.NumMethod(); i++ {
			log.Printf("###%v", v.Method(i).String())
		}
	}

	switch v.Kind() {
	case reflect.Struct:
		log.Println("=== 结构体 ===")
	}
}

func typeOf(t reflect.Type) {
	log.Println("###########TypeOf##############")
	log.Println("分类", t.Kind())
	log.Println("方法数量", t.NumMethod())
	//指针
	if t.Kind() == reflect.Ptr {
		log.Println("指针的字段值", t.Elem())
	}

	if t.Kind() == reflect.Struct {
		log.Println("字段数量", t.NumField())
		//遍历字段并打印字段名
		for i := 0; i < t.NumField(); i++ {
			log.Println("### 字段名", t.Field(i).Name)
			log.Println("### tag全部", t.Field(i).Tag)              //json:"addr,omitempty"
			log.Println("### tag的内容", t.Field(i).Tag.Get("json")) //addr,omitempty
			if tagStr, ok := t.Field(i).Tag.Lookup("json"); ok {
				log.Println(tagStr) //addr,omitempty
			}
		}
	}
	//遍历防范并打印方法名
	for i := 0; i < t.NumMethod(); i++ {
		log.Println("###方法名", t.Method(i).Name)
	}

}
func main() {
	log.SetFlags(log.Lshortfile)

	//反射结构体
	d := Demo{
		Name: "peter",
		Sex:  1,
		Age:  19,
		Addr: "PRC",
	}
	d.One()
	d.Two()

	valueOf(reflect.ValueOf(d))
	typeOf(reflect.TypeOf(d))

}
