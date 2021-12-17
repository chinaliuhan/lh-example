package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"log"
	"sync"
	"time"
)

var db *bolt.DB         //bolt句柄
var userBucket = "user" //桶名称
var tmpBucket = "tmp_bucket"

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// u64tob converts a uint64 into an 8-byte slice.
func u64tob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btou64 converts an 8-byte slice into an uint64.
func btou64(b []byte) uint64 { return binary.BigEndian.Uint64(b) }

type person struct {
	Id       int
	Name     string
	Age      uint8
	CreateAt time.Time
}

func putData(p person) {

	//这里暂时用JSON做序列化, 如果想要性能更加强悍,可以使用 protobuf
	pJson, err := json.Marshal(p)
	if err != nil {
		return
	}

	//更新数据
	aaa := db
	err = aaa.Update(func(tx *bolt.Tx) error {
		//创建桶,如果桶不存在
		bucket, err := tx.CreateBucketIfNotExists([]byte(userBucket))
		if err != nil {
			return err
		}
		//向该桶中推入数据
		err = bucket.Put(itob(p.Id), pJson)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func viewData(key int) []byte {
	var buf []byte
	//获取数据
	db.View(func(tx *bolt.Tx) error {
		//读取桶
		bucket := tx.Bucket([]byte(userBucket))
		if bucket == nil {
			return errors.New(" not has this bucket")
		}
		//通过key读取桶内信息
		val := bucket.Get(itob(key))
		if val == nil {
			return errors.New(" not has this key")
		}
		buf = val
		return nil
	})
	return buf
}

func rangeView(bucketName string) error {
	//遍历整个桶
	err := db.View(func(tx *bolt.Tx) error {
		//读取桶
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return errors.New(" not has this bucket")
		}
		//遍历桶内所有的数据
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("遍历桶 key=%d, value=%s", btou64(k), v)
		}

		//通过key读取桶内信息
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func main() {
	/**
	go的持久化方案的一种, 通过boltdb 将内存中的数据存储到本地文件中

	BoltDB是性能非常强悍的纯Go语言实现的持久化解决方案而非数据库

	bolt存储是分组为桶，桶是一组键值对集合的名称，就像Go中的map。桶的名称、键以及值都是[]byte类型。桶可以包括其他桶，也可以通过[]byte类型名称作为key。
	*/

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//在当前目录下打开my.db数据文件。如果它不存在，它将被创建。
	var err error
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//读取整个桶的数据
	rangeView(userBucket)

	//读取数据
	log.Printf("启动读取 %s", viewData(1))

	//写入数据
	p0 := person{
		Id:       1,
		Name:     "ZhangSan",
		Age:      10,
		CreateAt: time.Now(),
	}
	putData(p0)

	//读取数据
	log.Printf("读取0 %s", viewData(1))

	//写入数据
	p1 := person{
		Id:       2,
		Name:     "LiSi",
		Age:      11,
		CreateAt: time.Now(),
	}
	putData(p1)

	//读取数据
	log.Printf("读取1 %s", viewData(2))

	//事务
	log.Println("开始执行事务")
	err = transaction()
	if err != nil {
		log.Fatalln(err)
	}

	//批量事务处理
	log.Println("开始批量事务")
	batch()
	//遍历批量事务的桶内所有的数据
	err = rangeView(tmpBucket)
	if err != nil {
		log.Fatalln(err)
	}

	//删除桶
	delBucket(tmpBucket)
}
func delBucket(bucketName string) {
	db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucketName))
	})
}

func batch() {
	//批量更新事务
	//并发批处理调用机会地合并到更大的事务中。Batch只有在有多个goroutine调用时才有用。
	//这样做的代价是，如果部分事务失败，Batch可以多次调用给定函数。函数必须是幂等的，并且副作用必须在DB.Batch()成功返回后才生效
	wg := &sync.WaitGroup{}
	errCh := make(chan error)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(db *bolt.DB, wg *sync.WaitGroup, i int) {
			errCh <- db.Batch(func(tx *bolt.Tx) error {
				//这里可以有多个update操作,通实现多次数据更新操作,所有的更新会被当做一个事务来处理，如果Update()内的操作返回nil，则事务会被提交，否则事务会回滚
				bucket, err := tx.CreateBucketIfNotExists([]byte(tmpBucket))
				if err != nil {
					return err
				}
				err = bucket.Put(u64tob(uint64(i)), []byte{})
				if err != nil {
					return err
				}
				if i == 5 {
					//模拟update 失败,整体被回滚
					return errors.New("batch 批量update失败")
				}
				log.Println("批量事务推入", i)
				return nil
			})

			wg.Done()
		}(db, wg, i)
	}
	for i := 0; i < 10; i++ {
		if err := <-errCh; err != nil {
			log.Println(err)
		}
	}
	wg.Wait()
}

func transaction() error {
	// 开启事务,这里的true代表是可写入的
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	//回滚事务
	defer tx.Rollback()

	// 不存在则创建桶
	_, err = tx.CreateBucketIfNotExists([]byte(tmpBucket))
	if err != nil {
		return err
	}
	err = tx.DeleteBucket([]byte(tmpBucket))
	if err != nil {
		return err
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
