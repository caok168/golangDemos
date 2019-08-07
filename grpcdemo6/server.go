package main

import (
	pb "golangDemos/grpcdemo6/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type Streamer struct {
}

func (s *Streamer) StreamSayHello(stream pb.Greeter_StreamSayHelloServer) error {

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			recv, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错：", err)
				return err
			}

			switch recv.Name {
			case "结束对话 ":
				log.Println("收到‘结束对话’指令")
				if err := stream.Send(&pb.HelloReply{Name: "收到结束指令"}); err != nil {
					return err
				}
			case "返回数据流 ":
				log.Println("收到‘返回数据流’指令")
				for i := 0; i < 10; i++ {
					if err := stream.Send(&pb.HelloReply{Name: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}

				//default:
				//	log.Printf("[收到消息] %s", recv.Input)
				//	if err := stream.Send(&pb.Response{Output: "服务端返回： " + recv.Input}); err != nil {
				//		return err
				//	}
			}

		}
	}
}

func main() {
	log.Println("启动服务端...")
	server := grpc.NewServer()

	pb.RegisterGreeterServer(server, &Streamer{})
	address, err := net.Listen("tcp", ":3333")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
