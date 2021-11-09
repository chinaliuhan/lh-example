package main

import (
	"log"
	"net"

	pb "lh-example/lh-grpc/proto/simple" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8888"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService ...
var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	log.Println("SayHello收到消息: ", in.Name)
	resp.Message = "你好 " + in.Name + "."

	return resp, nil
}

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println("Search收到消息: ", r.GetRequest())

	return &pb.SearchResponse{Response: "你好" + r.GetRequest()}, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//开启TCP服务
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln("TCP监听失败", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)
	//注册SearchService
	pb.RegisterSearchServiceServer(s, &SearchService{})

	err = s.Serve(listen)
	if err != nil {
		log.Fatalln("GRPC监听失败", err)
	}

	log.Println("服务已启动 " + Address)
}
