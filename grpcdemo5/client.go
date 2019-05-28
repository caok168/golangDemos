package main

import (
	"bufio"
	"context"
	pb "golangDemos/grpcdemo5/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：【%v】", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := pb.NewChatClient(conn)
	// 声明 context
	ctx := context.Background()

	// 创建双向流
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Printf("创建数据流失败：【%v】", err)
	}

	go func() {
		log.Printf("请输入消息...")
		in := bufio.NewReader(os.Stdin)

		for {
			inputStr, _ := in.ReadString(' ')
			if err := stream.Send(&pb.Request{Input: inputStr}); err != nil {
				return
			}

		}
	}()

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("收到服务器的结束信号")
			break
		}
		if err != nil {
			log.Println("接收数据出错：", err)
		}
		log.Printf("[客户端收到]: %s", response.Output)
	}
}

//protoc -I grpcdemo/routeguideDemo/proto/ grpcdemo/routeguideDemo/proto/route_guide.proto --go_out=plugins=grpc:grpcdemo/routeguideDemo/proto/
