package lh_common

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

/**
文件锁
*/
type fileLock struct {
	file string   // 目录路径，例如 /home/XXX/go/src
	f    *os.File // 文件描述符
}

// 新建一个 fileLock
func NewFileLock(fileName string) *fileLock {
	return &fileLock{
		file: fileName,
	}
}

// 加锁操作
func (r *fileLock) Lock() error {
	f, err := os.Open(r.file) // 获取文件描述符
	if os.IsNotExist(err) {
		f, err = os.Create(r.file)
		if err != nil {
			return err
		}
	}
	_, err = f.WriteString("flock " + time.Now().String())
	if err != nil {
		return err
	}

	r.f = f
	//Flock 是建议性的锁，使用的时候需要指定 how 参数，否则容易出现多个 goroutine 共用文件的问题 how 参数指定 LOCK_NB 之后，goroutine 遇到已加锁的 Flock，不会阻塞，而是直接返回错误
	fdInt := int(f.Fd())
	err = syscall.Flock(fdInt, syscall.LOCK_EX|syscall.LOCK_NB) // 加上排他锁，当遇到文件加锁的情况直接返回 Error
	if err != nil {
		return fmt.Errorf("cannot flock file %s - %s", r.file, err)
	}
	return nil
}

// 解锁操作
func (r *fileLock) Unlock() error {
	defer r.f.Close()                                    // close 掉文件描述符
	return syscall.Flock(int(r.f.Fd()), syscall.LOCK_UN) // 释放 Flock 文件锁
}
