package main

import (
	"log"
	"strconv"
)

//将字符串解析成指定的类型
func parse() {
	log.Println(strconv.ParseBool("true"))
	log.Println(strconv.ParseComplex("10", 64))
	log.Println(strconv.ParseFloat("0.001", 64))
	log.Println(strconv.ParseInt("100", 10, 64))
	log.Println(strconv.ParseUint("200", 10, 64))
	//将字符串转换为int类型
	log.Println(strconv.Atoi("300"))
}

//将其他类型格式化为字符串
func format() {
	log.Println(strconv.FormatBool(true))
	//将复数格式化为字符串
	//log.Println(strconv.FormatComplex())
	/**
	0.001
	'f' 设置格式
		'b' (-ddddp±ddd，二进制指数)
		'e'(-d.dddde±dd。十进制指数Dddde±dd)
		'E'(-d.ddddE±dd. 十进制指数ddddE±dd)
		'f' (-ddd.dddd,没有指数)
		'g' ('e'表示大指数，'f'表示其他)
		'G' ('E'表示大指数，'f'表示大指数)，
		'x'(-0xd.ddddp±ddd。(十六进制分数和二进制指数)，或 'X'(-0Xd.ddddP±ddd(十六进制分数和二进制指数)。
	6 小数点后位数
	64 bitSize的值(float32为32,float64为64)
	*/
	log.Println(strconv.FormatFloat(0.00123456789, 'f', 6, 64))
	log.Println(strconv.FormatInt(100, 10))
	log.Println(strconv.FormatUint(200, 10))
	//将整型格式化为字符串
	log.Println(strconv.Itoa(10))
}

////就是将字符的一些符号来还转换
func quote() {

	// Quote 将字符串 s 转换为“双引号”引起来的字符串
	// 其中的特殊字符将被转换为“转义字符”
	// “不可显示的字符”将被转换为“转义字符”
	//将 s 转换为双引号字符串
	log.Println(strconv.Quote("对接当个人生几何"))

	// QuoteRune 将 Unicode 字符转换为“单引号”引起来的字符串
	// “特殊字符”将被转换为“转义字符”
	//将 r 转换为单引号字符
	log.Println(strconv.QuoteRune('哈'))

	// QuoteToASCII 将字符串 s 转换为“双引号”引起来的 ASCII 字符串
	// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
	//非 ASCII 字符和不可打印字符会被转义
	log.Println(strconv.QuoteToASCII("哈哈哈"))

	// QuoteToGraphic返回一个双引号的Go字符串字面量表示s。
	//返回的字符串保留Unicode图形字符，由
	// IsGraphic，没有改变，使用Go转义序列(\t， \n， \xFF， \u0100)
	//非图形字符。
	//非图形字符会被转义
	log.Println(strconv.QuoteToGraphic("哈哈你好"))

	// QuoteRuneToASCII 将 Unicode 字符转换为“单引号”引起来的 ASCII 字符串
	// “非 ASCII 字符”和“特殊字符”将被转换为“转义字符”
	log.Println(strconv.QuoteRuneToASCII('嘿'))

	// QuoteRuneToGraphic返回一个单引号的Go字符字面表示
	//符文。如果符文不是Unicode图形字符，
	//根据IsGraphic的定义，返回的字符串将使用Go转义序列
	// (\t， \n， \xFF， \u0100)
	log.Println(strconv.QuoteRuneToGraphic('嘎'))

	// Unquote 将“带引号的字符串” s 转换为常规的字符串（不带引号和转义字符）
	// s 可以是“单引号”、“双引号”或“反引号”引起来的字符串（包括引号本身）
	// 如果 s 是单引号引起来的字符串，则返回该该字符串代表的字符
	log.Println(strconv.Unquote("\"你\t好\t啊,你\t是\t谁\""))

	// 将 s 中的第一个字符“取消转义”并解码
	//UnquoteChar 将带引号字符串（不包含首尾的引号）中的第一个字符“取消转义”并解码
	//
	// s    ：带引号字符串（不包含首尾的引号）
	// quote：字符串使用的“引号符”（用于对字符串中的引号符“取消转义”）
	//
	// value    ：解码后的字符
	// multibyte：value 是否为多字节字符
	// tail     ：字符串 s 解码后的剩余部分
	// error    ：返回 s 中是否存在语法错误
	//
	// 参数 quote 为“引号符”
	// 如果设置为单引号，则 s 中允许出现 \'、" 字符，不允许出现单独的 ' 字符
	// 如果设置为双引号，则 s 中允许出现 \"、' 字符，不允许出现单独的 " 字符
	// 如果设置为 0，则不允许出现 \' 或 \" 字符，但可以出现单独的 ' 或 " 字符
	s := `\"你\\好\\啊,你是\t哪位?\"`
	c, mb, sr, _ := strconv.UnquoteChar(s, '"')
	log.Printf("%-3c %v\n", c, mb)
	log.Println("################")
	for ; len(sr) > 0; c, mb, sr, _ = strconv.UnquoteChar(sr, '"') {
		log.Printf("%-3c %v\n", c, mb)
	}
}

func append1() {
	sli := make([]byte, 0)
	sli = strconv.AppendBool(sli, true)
	sli = strconv.AppendFloat(sli, 0.001, 'f', 6, 64)
	sli = strconv.AppendInt(sli, 111, 10)

	/**
	AppendQuote 将字符串 s 转换为“双引号”引起来的字符串，
	并将结果追加到 dst 的尾部，返回追加后的 []byte
	其中的特殊字符将被转换为“转义字符”
	*/
	sli = strconv.AppendQuote(sli, "对接当个人生几何")

	log.Println(string(sli))

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//将字符串解析成指定的类型
	parse()

	//将其他类型格式化为字符串
	format()

	//将其他类型append到[]byte切片并转换成[]byte
	append1()

	//就是将字符的一些符号来还转换
	quote()

	//判断 Unicode 字符 r 是否是一个可显示的字符
	log.Println(strconv.IsPrint('哈'))

	//判断rune字符 r 是否是 unicode 图形字符。图形字符包括字母、标记、数字、符号、标点、空白。
	log.Println(strconv.IsGraphic('1'))  //true
	log.Println(strconv.IsGraphic('a'))  //true
	log.Println(strconv.IsGraphic('好'))  //true
	log.Println(strconv.IsGraphic('█'))  //true
	log.Println(strconv.IsGraphic('😁')) //true
}
