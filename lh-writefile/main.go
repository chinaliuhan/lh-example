package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var (
		filePath string
		n        int
		err      error
		content  []byte
	)

	//设置日志输出格式,带文件名,行号,时间
	log.SetFlags(log.Lshortfile | log.Ldate)

	/**
	方案1
	*/
	filePath = "./write.log"
	content = []byte(time.Now().String() + " this is WriteFile bytes  line1\n")
	//覆盖性的写入文件,如果不存在则创建,如果存在则覆盖
	err = os.WriteFile(filePath, content, 0755)
	if err != nil {
		log.Println(err)
		return
	}

	/**
	方案2
	*/
	filePath = "./write1.log"
	content = []byte(time.Now().String() + " this is Create bytes line1\n")
	//覆盖性的创建文件,返回资源句柄,创建创建或截断命名文件。如果文件已经存在被截断。如果文件不存在，它将以0666模式创建
	create, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer create.Close()
	//写入[]byte内容,返回写入长度
	n, err = create.Write(content)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write bytes done", n)
	//写入字符串内容,返回写入字符长度
	n, err = create.WriteString(time.Now().String() + " this is Create string line2\n")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("writeString done", n)

	/**
	方案3 打开文件,然后在后面追加写入内容
	*/
	filePath = "./write3.log"
	content = []byte(time.Now().String() + " this is OpenFile bytes line1\n")
	//读取文件,如果不存在就创建,追加,读写,如果不存在则创建
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//写入[]byte并返回长度
	n, err = file.Write(content)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write bytes done", n)
	//写入字符串并返回字符长度
	n, err = file.WriteString(time.Now().String() + " this is OpenFile string line2\n")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("write string done ", n)

	/**
	方案4 写入Json文件,将Map转换为Json写入文件
	*/
	dic := map[int]string{1: "qin", 2: "han", 3: "sui", 4: "tang", 5: "song", 6: "yuan", 7: "ming"}
	filePath = "./tmp.json"

	bb := new(bytes.Buffer)
	je := json.NewEncoder(bb)
	err = je.Encode(dic)
	if err != nil {
		log.Println(err)
	}
	//覆盖性的写入文件,如果不存在则创建,如果存在则覆盖
	err = os.WriteFile(filePath, bb.Bytes(), 0755)
	if err != nil {
		log.Println(err)
	}
	log.Println("write json done ")

	/**
	方案5写入数组
	*/
	var numbers = []int{1, 2, 3, 4, 5, 6}
	filePath = `./tmp.sli`

	b := new(bytes.Buffer)
	e := json.NewEncoder(b)

	err = e.Encode(numbers)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(filePath, b.Bytes(), 0755)
	if err != nil {
		log.Println(err)
	}

	/**
	方案6 写入二进制1
	*/
	filePath = `./tmp_bin.log`
	data := []byte{1, 2, 4, 8, 16, 32, 64, 128}

	err = ioutil.WriteFile(filePath, data, 0755)
	if err != nil {
		log.Println(err)
	}
	//写入二进制2
	outFile, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
	}
	n, err = outFile.Write(data)
	if err != nil {
		log.Println(err)
	}
	log.Println(n, "bytes write done")

}
