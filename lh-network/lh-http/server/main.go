package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

func ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	date := time.Now().String()
	w.Write([]byte(date))
	return
}
func ping1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	date := time.Now().String()
	fmt.Fprintf(w, date) //写入到w的时输出到客户端
	return
}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法

	//判断请求方式
	if r.Method == "GET" { //get请求,获取模板展示
		//获取当前时间戳
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//这个应该是读取静态页面
		t, _ := template.ParseFiles("upload.gtpl")
		//打印页面内容,同时输出变量到页面
		t.Execute(w, token)
	} else { //其他请求,代表提交数据
		//调用r.ParseMultipartForm,里面的参数maxMemory,调用ParsemultipartForm之后,上传的文件存储在maxMemory大小的内存里,如果文件大小超过了maxMemory,那么剩下的部分将存储在系统的临时文件里,我们通过r.FormFile获取上面的文件句柄,然后利用使用io.Copy来存储文件
		//获取其他非文件字段信息的时候就不需要调用r.ParseForm，因为在需要的时候Go自动会去调用。而且ParseMultipartForm调用一次之后，后面再次调用不会再有效果。
		r.ParseMultipartForm(32 << 20)
		//通过上传的文件名,获取上传的文件
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//保存文件,同时赋予文件666的权限
		//openFile和open的地方在于open只能用来读取文件,handler.Filename获取文件名
		//os.O_WRONLY | os.O_CREATE | O_EXCL           【如果已经存在，则失败】
		//os.O_WRONLY | os.O_CREATE                         【如果已经存在，会覆盖写，不会清空原来的文件，而是从头直接覆盖写】
		//os.O_WRONLY | os.O_CREATE | os.O_APPEND  【如果已经存在，则在尾部添加写】
		//这里这个handler.Filename,会获取到包括路径在内的文件名,所以我改成了时间
		f, err := os.OpenFile("./tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

var static = flag.String("static", "", "指定静态文件目录")

func main() {
	flag.Parse()
	if *static == "" {
		log.Println("请指定静态文件目录")
	}

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/ping1", ping1)
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}

}
