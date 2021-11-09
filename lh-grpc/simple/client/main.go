package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "lh-example/lh-grpc/proto/simple" // 引入编译生成的包
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8888"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	for i := 0; i < 100; i++ {

		// 初始化客户端, 调用方法
		c := pb.NewHelloClient(conn)
		reqBody := new(pb.HelloRequest)
		reqBody.Name = "gRPC 调用 SayHello"
		r, err := c.SayHello(context.Background(), reqBody)
		if err != nil {
			log.Fatalln("服务初始化失败", err)
		}
		log.Println("resp:", r.Message)

		// 初始化客户端, 调用方法
		client := pb.NewSearchServiceClient(conn)
		resp, err := client.Search(context.Background(), &pb.SearchRequest{
			Request: "gRPC 调用 Search",
		})
		if err != nil {
			log.Fatalln("服务初始化失败", err)
		}
		log.Println("resp:", resp.GetResponse())
	}

}
