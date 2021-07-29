package main

import (
	"bytes"
	"log"
	"regexp"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 测试模式是否匹配字符串,可以简单概括为在指定的字符串中判断是否包含该匹配规则
	match, _ := regexp.MatchString("s([a-z]+)ck", "suck")
	log.Println(match)

	//预先编译解析一个正则表达式，如果成功返回一个可以用来匹配文本的Regexp对象,后面可以一直服用该表达式
	compile, err := regexp.Compile("s([a-z]+)ck")
	if err != nil {
		return
	}
	//使用一个编译解析的正则表达式,判断该字符串中是否包含该匹配规则
	log.Println(compile.MatchString("suck"))

	//使用一个编译解析的正则表达式 从指定的字符串中中查询指定的字符串
	log.Println(compile.FindString("suck"))

	//这个方法查找第一次匹配的索引，并返回匹配字符串的起始索引和结束索引
	log.Println(compile.FindStringIndex("suck blood"))

	//返回全局匹配的字符串和局部匹配的字符
	log.Println(compile.FindStringSubmatch("suck blood"))

	//和上面的方法一样，不同的是返回全局匹配和局部匹配的起始索引和结束索引
	log.Println(compile.FindStringSubmatchIndex("suck blood"))

	//这个方法返回所有正则匹配的字符，不仅仅是第一个
	log.Println(compile.FindAllString("suck blood capital suuck", -1))

	//这个方法返回所有全局匹配和局部匹配的字符串起始索引和结束索引
	log.Println(compile.FindAllStringSubmatchIndex("suck blood capital", -1))

	//为这个方法提供一个正整数参数来限制匹配数量
	log.Println(compile.FindAllString("suck blood capital", 2))

	//上面我们都是用了诸如MatchString这样的方法，其实也可以使用[]byte作为参数，并且使用Match这样的方法名
	log.Println(compile.Match([]byte("suck")))

	// 当使用正则表达式来创建常量的时候，可以使用MustCompile  因为Compile返回两个值,但如果表达式不能被解析MustCompile就会panic
	compile = regexp.MustCompile("s([a-z]+)er")
	log.Println(compile)
	//regexp包也可以用来将字符串的一部分替换为指定的字符串
	log.Println(compile.ReplaceAllString("a sucker", "replace"))
	// ReplaceAllFunc可以让将所有匹配的字符串都经过该函数处理,转变为所需要的值
	in := []byte("a sucker")
	out := compile.ReplaceAllFunc(in, bytes.ToUpper)
	log.Println(string(out))

}
