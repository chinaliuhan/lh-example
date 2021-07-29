package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//简单的GET请求,不能配置
func simpleGet() {

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println("get响应", string(body))
}

//简单的POST请求,不能配置
func simplePost() {
	resp, err := http.Post("https://www.baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=li&age=19"),
	)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}

//简单的POSTForm请求不能配置
func simplePostForm() {
	reqBody := url.Values{}
	reqBody.Set("name", "li")
	reqBody.Set("age", "19")
	resp, err := http.PostForm("http://www.baidu.com", reqBody)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}

//能进行除了header之外的配置,例如代理,cookie,超时时间,指定处理重定向的策略等
func newClient() {
	var (
		err        error
		httpClient *http.Client
		resp       *http.Response
	)

	httpClient = &http.Client{
		//Transport: &http.Transport{
		//	Proxy: func(r *http.Request) (*url.URL, error) {
		//		//这是个回调函数
		//		return url.Parse("socks5://127.0.0.1:7890")
		//	}},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       10 * time.Second,
	}

	resp, err = httpClient.Post("https://www.baidu.com", "application/json", strings.NewReader("name=li&age=19"))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

//可配置所有内容, 包括http的配置,请求头,超时,代理,指定处理重定向的策略等
func newClientNewReq() {
	//第一步
	client := &http.Client{}
	urlPath := "https://www.baidu.com"

	//第二步
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return
	}

	//可以在这里设置各种header等
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=li&age=19&addr=beijing")

	//设置代理
	client.Transport = &http.Transport{
		//Proxy: func(r *http.Request) (*url.URL, error) {
		//	//这是个回调函数
		//	return url.Parse("socks5://127.0.0.1:7890")
		//},
	}
	//设置请求超时时间
	client.Timeout = time.Second * 5

	//第三步
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Println("响应状态码", resp.StatusCode)

	log.Println("获取到的body", string(body))
}

//上传函数
//下面的例子详细展示了客户端如何向服务器上传一个文件的例子，客户端通过multipart.Write把文件的文本流写入一个缓存中，然后调用http的Post方法把缓存传到服务器。
func uploadFile(filename string, targetUrl string) {
	//创建控件
	bodyBuf := &bytes.Buffer{}
	//创建一个可写入资源,吐槽一下,multipart这个包,设计的太二笔了...
	bodyWrite := multipart.NewWriter(bodyBuf)

	//从文件中读取数据, 创建表单文件名
	fileWrite, err := bodyWrite.CreateFormFile("uploadfile", filename)
	if err != nil {
		log.Println(err)
	}

	//打开文件,.open的形式打开,只能用作读取
	fh, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer fh.Close()

	//ioCopy,从fh复制到fileWrite，直到到达EOF或发生错误,返回拷贝的字节喝遇到的第一个错误.
	_, err = io.Copy(fileWrite, fh)
	if err != nil {
		log.Println(err)
	}

	//返回http的请求需要的类型
	contentType := bodyWrite.FormDataContentType()
	bodyWrite.Close()
	//开始上传
	response, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(response.Status)
	log.Println(string(responseBody))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//简单请求,不能配置,只有 PostForm 可以选择Content-Type
	simpleGet()
	simplePost()
	simplePostForm()

	//可以进行除header之外的配置,如代理,超时时间,重定向规则等
	newClient()

	//可以进行任何配置
	newClientNewReq()

	//文件上传
	uploadFile("./tmp.sli", "http://localhost:8080/upload")
}
