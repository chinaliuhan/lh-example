package main

import (
	"google.golang.org/grpc"
	"io"
	pb "lh-example/lh-grpc/proto"
	"log"
	"net"
)

type StreamService struct{}

//List 服务端主动发送
func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		//发送消息
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

//Record 服务端接收
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		//接收消息
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

}

//Route 双向
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		//发送消息
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		//接收消息
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//实例化grpc服务
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})

	//监听端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln("grpc监听失败", err)
	}

	//启动grpc服务
	err = server.Serve(listen)
	if err != nil {
		log.Fatalln("grpc服务启动失败", err)
	}
}
