package fetcher

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

//指定时间,自动向channel中推送一个时间
//var rateLimit = time.Tick(time.Millisecond * 100)

//Fetch 通过URL获取页面
func Fetch(url string) ([]byte, error) {
	//<-rateLimit //限流

	//获取页面
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logrus.Error("http 响应状态错误", err)
		return nil, fmt.Errorf("http状态错误 %d", resp.StatusCode)
	}

	//读取页面信息并转码为utf-8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//determineEncoding golang.org/x/text 将GBK转换为UTF8
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	peek, err := r.Peek(1024)
	if err != nil {
		logrus.Error("编码获取失败", err)
		return unicode.UTF8
	}
	//golang.org/x/net/html 简则字符串的字符集,例如UTF8,GBK,等
	e, _, _ := charset.DetermineEncoding(peek, "")
	return e
}
