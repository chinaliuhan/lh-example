package main

import (
	"encoding/json"
	"os"
)
import "log"

// 下面我们将使用这两个结构体来演示自定义类型的编码和解码。
type pager1 struct {
	Page   int
	Fruits []string
}
type pager2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}
type pager3 struct {
	Page   int      `json:"page"`
	Count  int      `json:"-"`
	Fruits []string `json:"fruits"`
}

func baseType2Json() {
	// 基本类型到json的结果
	bolB, _ := json.Marshal(true)
	log.Println(string(bolB))

	intB, _ := json.Marshal(1)
	log.Println(string(intB))

	fltB, _ := json.Marshal(1.11)
	log.Println(string(fltB))

	strB, _ := json.Marshal("aaa")
	log.Println(string(strB))
}

func mapSlice2Json() {
	//将slice编码为json数组
	slcD := []string{"apple", "peach", "orange"}
	slcB, _ := json.Marshal(slcD)
	log.Println(string(slcB))

	//将map编码为json对象
	mapD := map[string]int{"apple": 5, "orange": 7}
	mapB, _ := json.Marshal(mapD)
	log.Println(string(mapB))

}

func custom2Json() {
	// JSON 包可以自动的编码你的自定义类型。编码仅输出可导出的字段，并且默认使用他们的名字作为 JSON 数据的键。
	res1D := &pager1{
		Page:   1,
		Fruits: []string{"apple", "peach", "orange"}}
	res1B, _ := json.Marshal(res1D)
	log.Println(string(res1B))
}

func customTag2Json() {
	// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。在上面 `pager2` 的定义可以作为这个标签这个的一个例子。
	res2D := pager2{
		Page:   1,
		Fruits: []string{"apple", "orange", "pear"}}
	res2B, _ := json.Marshal(res2D)
	log.Println(string(res2B))
}
func customTag2JsonSkip() {
	// 你可以给结构字段声明标签来自定义编码的 JSON 数据键名称。在上面 `pager2` 的定义可以作为这个标签这个的一个例子。
	res2D := pager3{
		Page:   1,
		Fruits: []string{"apple", "orange", "pear"},
		Count:  111, //这里虽然指定了count的值,但是因为pager3中的tag标记为-,所以即使这里给值了,转码为json时也不会有该字段
	}
	res2B, _ := json.Marshal(res2D)
	log.Println(string(res2B))
}

func json2Map() {

	// 现在来看看解码 JSON 数据为 Go 值的过程。这里是一个普通数据结构的解码例子。
	byt := []byte(`{"num":1111,"strs":["a","b"]}`)

	// 我们需要提供一个 JSON 包可以存放解码数据的变量。这里 的 `map[string]interface{}` 将保存一个 string 为键,值为任意值的map。
	var dat map[string]interface{}

	// 这里就是实际的解码和相关的错误检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Fatalln(err)
	}
	log.Println(dat)

	// 为了使用解码 map 中的值，我们需要将他们进行适当的类型转换。例如这里我们将 `num` 的值转换成 `float64`类型。
	num := dat["num"].(float64)
	log.Println(num)

	// 访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	log.Println(str1)
}

func json2Struct() {
	// 我们也可以解码 JSON 值到自定义类型。这个功能的好处就是可以为我们的程序带来额外的类型安全加强，并且消除在 访问数据时的类型断言。
	str := `{"page": 1, "fruits": ["apple", "orange"]}`
	res := &pager2{}
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	log.Println(res.Fruits[0])
}

func map2Json2Stdout() {
	// 在上面的例子中，我们经常使用 byte 和 string 作为使用 标准输出时数据和 JSON 表示之间的中间值。我们也可以和 `os.Stdout` 一样，直接将 JSON 编码直接输出至 `os.Writer` 流中，或者作为 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "orange": 7}
	err := enc.Encode(d)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//基本类型到json的结果
	//baseType2Json()

	//切片和 map 编码成 JSON 数组和对象
	//mapSlice2Json()

	//自定义一个struct解析为json
	//custom2Json()

	//自定义一个struct,通过tag解析为json
	//customTag2Json()

	//自定义一个struct,通过tag解析为json,同时跳过不想解析的字段
	//customTag2JsonSkip()

	//将json解析为map
	//json2Map()

	//将json解析为struct
	//json2Struct()

	//将map解析为json并打印到标准输出
	map2Json2Stdout()
}
