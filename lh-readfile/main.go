package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var filePath string = "./write.log"

	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	方案1
	*/
	bytContent, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	log.Println(string(bytContent))

	/**
	方案2
	*/
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		return
	}

	log.Println(string(all))

	/**
	方案3 逐行读取
	*/

	//打开文件,返回资源句柄
	open, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer open.Close()
	//返回一个新的Scanner来读取传入的文件
	scanner := bufio.NewScanner(open)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	for scanner.Scan() {
		//每遍历一次读取一行字符串
		log.Println("逐行打印: " + scanner.Text())
	}

	/**
	方案4 读取json文件,将JSON文件解析到map
	*/
	filePath = "./tmp.json"
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	bb := new(bytes.Buffer)
	bb.Write(readFile)

	dic := map[int]string{}
	j := json.NewDecoder(bb)
	j.Decode(&dic)

	log.Println(dic)

	/**
	方案5 读取数组
	*/
	filePath = `./tmp.sli`
	data, _ := ioutil.ReadFile(filePath)

	b := new(bytes.Buffer)
	b.Write(data)

	var numbers []int
	var d = json.NewDecoder(b)
	d.Decode(&numbers)

	log.Println(numbers)

	/**
	方案6 读取二进制1
	*/
	filePath = `./tmp_bin.log`
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print(err)
	}
	log.Println(bytes)

	//读取二进制2
	file, err = os.Open(filePath)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	bytes2, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
	}
	log.Println(bytes2)
}
