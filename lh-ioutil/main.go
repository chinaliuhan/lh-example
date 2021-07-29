package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r1 := strings.NewReader("### this is ioutil one\n")

	// ReadAll 读取 r 中的所有数据，返回读取的数据和遇到的错误。
	// 如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
	// 所有数据，所以不会把 EOF 当做错误处理。
	all, err := ioutil.ReadAll(r1)
	if err != nil {
		return
	}
	log.Println(string(all))

	//ReadDir 读取由 dirname 命名的目录，并返回目录内容的 fs.FileInfo 列表，按文件名排序。如果读取目录时发生错误，ReadDir 不会随错误返回任何目录条目。
	//从 Go 1.16 开始，os.ReadDir 是一个更有效和正确的选择：它返回 fs.DirEntry 的列表而不是 fs.FileInfo，并且在读取目录中途出错的情况下返回部分结果。
	dir, err := ioutil.ReadDir("./tmp")
	if err != nil {
		return
	}

	for k, v := range dir {
		log.Println(k, v.Name())    // base name of the file
		log.Println(k, v.Size())    // length in bytes for regular files; system-dependent for others
		log.Println(k, v.Mode())    // file mode bits
		log.Println(k, v.ModTime()) // modification time
		log.Println(k, v.IsDir())   // abbreviation for Mode().IsDir()
		log.Println(k, v.Sys())     // underlying data source (can return nil)
	}

	//ReadFile 读取以 filename 命名的文件并返回内容。成功的调用返回 err == nil，而不是 err == EOF。因为 ReadFile 读取整个文件，它不会将 Read 中的 EOF 视为要报告的错误。
	//从 Go 1.16 开始，这个函数只调用 os.ReadFile。
	file, err := ioutil.ReadFile("./tmp.log")
	if err != nil {
		return
	}
	log.Println(string(file))

	//os.TempDir()可以在系统默认的目录生成一个临时目录,但是官方文档说了, 该目录既不保证存在，也没有访问权限。所以这里只做学习使用
	parentDir := os.TempDir()
	//TempDir 在目录 dir 中创建一个新的临时目录。目录名称是通过采用模式并在末尾应用随机字符串来生成的。如果模式包含“*”，则随机字符串替换最后一个“*”。TempDir 返回新目录的名称。如果 dir 是空字符串，则 TempDir 使用临时文件的默认目录（请参阅 os.TempDir）。多个程序同时调用 TempDir 不会选择同一个目录。当不再需要目录时，调用者有责任删除该目录。
	logsDir, err := ioutil.TempDir(parentDir, "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	os.RemoveAll(logsDir)
	log.Println(logsDir)

	//TempFile 在目录 dir 中创建一个新的临时文件，打开该文件进行读写，并返回生成的 *os.File。文件名是通过采用模式并在末尾添加一个随机字符串来生成的。如果模式包含“*”，则随机字符串替换最后一个“*”。如果 dir 是空字符串，TempFile 使用临时文件的默认目录（请参阅 os.TempDir）。多个程序同时调用 TempFile 不会选择同一个文件。调用者可以使用 f.Name() 来查找文件的路径名。调用者有责任在不再需要时删除文件。
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	os.RemoveAll(tmpfile.Name())

	//向文件写入数据,如果文件不存在则创建,如果存在则覆盖
	ioutil.WriteFile("./tmp.log", []byte("this is ioutil test file"), 0755)

	//NopCloser返回一个ReadCloser函数，它带有一个无操作的Close方法来包装所提供的Reader r。
	// NopCloser 将 r 包装为一个 ReadCloser 类型，但 Close 方法不做任何事情。
	//从 Go 1.16 开始，这个函数只调用 io.NopCloser。
	//ioutil.NopCloser()
}
