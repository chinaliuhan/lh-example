package main

import (
	"log"
	"sync"
)

var mutex = &sync.Mutex{}
var rwMutex = &sync.RWMutex{}

var lMap = make(map[string]string, 10)
var rwMap = make(map[string]string, 10)

func lockWrite() {
	//排它锁,读写都不是共享的,读写都是排他的
	go func() {
		mutex.Lock()
		lMap["first"] = "a"
		mutex.Unlock()
	}()
	go func() {
		mutex.Lock()
		lMap["second"] = "b"
		mutex.Unlock()
	}()
	go func() {
		mutex.Lock()
		lMap["third"] = "c"
		mutex.Unlock()
	}()
}
func lockRead() {
	//排它锁,读写都不是共享的,读写都是排他的
	mutex.Lock()
	log.Println(lMap["first"])
	log.Println(lMap["second"])
	log.Println(lMap["third"])
	mutex.Unlock()
}

func rwLockWrite() {
	//读写锁最大的不同是,读是共享的,可以并发读
	go func() {
		rwMutex.Lock()
		rwMap["first"] = "rw a"
		rwMutex.Unlock()
	}()
	go func() {
		rwMutex.Lock()
		rwMap["second"] = "rw b"
		rwMutex.Unlock()
	}()
	go func() {
		rwMutex.Lock()
		rwMap["third"] = "rw c"
		rwMutex.Unlock()
	}()
}
func rwLockRead() {
	rwMutex.Lock()
	log.Println(rwMap["first"])
	log.Println(rwMap["second"])
	log.Println(rwMap["third"])
	log.Println(rwMap["four"])
	log.Println(rwMap["fifth"])
	rwMutex.Unlock()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//多个协程写入,单个协程读取
	lockWrite()
	lockRead()

	//多个协程写入,单个协程读取
	rwLockWrite()
	rwLockRead()
}
