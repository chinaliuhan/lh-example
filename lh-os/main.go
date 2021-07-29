package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func fileStat() {

	//获取文件属性
	filePath := `./tmp`
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("文件大小是:", fileInfo.Size())
	log.Println("修改日期是:", fileInfo.ModTime())
	fileName := filepath.Base(filePath)
	extension := filepath.Ext(filePath)
	log.Println("文件名是:", fileName)
	log.Println("文件扩展是:", extension)
	log.Println("不带扩展名的文件是:", fileName[0:len(fileName)-len(extension)])
	log.Println("文件目录是:", filepath.Dir(filePath))

	//通过文件属性获取文件创建时间(Win下不可用,需要其他方式)
	stat := fileInfo.Sys().(*syscall.Stat_t)
	log.Println("Sys创建时间", time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec)))
	log.Println("Sys最后访问时间", time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec)))
	log.Println("Sys最后修改时间", time.Unix(int64(stat.Mtimespec.Sec), int64(stat.Mtimespec.Nsec)))
}

func copyDirFile() {
	//复制目录和目录下的文件
	path := `./tmp/1`
	newPath := `./tmp/1_copy`
	err := copyDir(path, newPath)
	if err != nil {
		log.Println(err)
	}

}

func delFileOrDir() {
	//删除文件或目录
	filePath := "./tmp.log"
	err := os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}

}
func forceDelDirOrFile() {
	//强制删除目录
	path := "./tmp/"
	err := os.RemoveAll(path)
	if err != nil {
		log.Println(err)
	}

}

func mkdFullDir() {
	//建立多级目录
	path := `./tmp/3`
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func mvFile() {
	//剪切文件
	file := `./tmp/file.log`
	newFile := `./tmp\new_file.log`
	err := os.Rename(file, newFile)
	if err != nil {
		log.Println(err)
	}
}

func fileExist() {
	//判断文件是否存在
	filePath := "./tmp"
	_, err := os.Stat(filePath)
	exist := !os.IsNotExist(err)
	if !exist {
		log.Println(err)
	} else {
		log.Println("文件存在")
	}

}
func dirFileList() {
	dir := `./tmp`
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		log.Println(f.Name())
	}

}
func env() {
	//设置环境变量
	os.Setenv("PROJECT_NAME", "lh-example")
	os.Setenv("NAME", "lh")

	//遍历所有环境变量
	for k, v := range os.Environ() {
		log.Println(k, v)
	}
	//获取环境变量
	log.Println(os.Getenv("NAME"))

}
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//从命令行的标准输入获取输入的字符串,命令行会阻塞在这里等待用户从命令行输入内容,并打印一行
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	log.Println(string(line))

	//向命令行的标准输出打印字符串
	fmt.Fprint(os.Stdout, "this is os\n")

	//设置环境变量
	env()

	//文件属性
	fileStat()

	//复制目录与其文件
	copyDirFile()

	//删除文件或目录
	delFileOrDir()

	//建立多级目录
	mkdFullDir()

	//剪切文件
	mvFile()

	//判断文件是否存在
	fileExist()

	//目录中的文件列表
	dirFileList()

	//强制删除目录
	forceDelDirOrFile()
}

//复制文件
func copyFile(src, newDir string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	newFile, err := os.Create(newDir)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, srcFile)
	return err
}

//复制目录
func copyDir(src, newDir string) (err error) {
	// 获取原目录属性
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	//创建新的目录
	err = os.MkdirAll(newDir, srcInfo.Mode())
	if err != nil {
		return err
	}

	openDir, _ := os.Open(src)
	fileInfo, err := openDir.Readdir(-1)

	for _, info := range fileInfo {

		srcPath := src + "/" + info.Name()
		newDirPath := newDir + "/" + info.Name()

		if info.IsDir() {
			err = copyDir(srcPath, newDirPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, newDirPath)
			if err != nil {
				return err
			}
		}
	}
	return
}
