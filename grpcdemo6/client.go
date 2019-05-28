package main

import (
	"context"
	"fmt"
	pb "golangDemos/grpcdemo6/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("192.168.11.5:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：【%v】", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := pb.NewGreeterClient(conn)
	// 声明 context
	ctx := context.Background()

	request := &pb.HelloRequest{
		Name: "gdh",
	}

	// 创建双向流
	stream, err := client.SayHello(ctx, request)
	if err != nil {
		log.Printf("创建数据流失败：【%v】", err)
	}

	fmt.Println("Name:", stream.Name)
	fmt.Println("Bbox:", stream.Bbox)
	fmt.Println("Sbox:", stream.Sbox)

	streamList, err := client.LstSayHello(ctx, request)
	if err != nil {
		log.Printf("创建数据流111失败：【%v】", err)
	}

	for {
		response, err := streamList.Recv()
		if err == io.EOF {
			log.Println("收到服务器的结束信号")
			break
		}
		if err != nil {
			log.Println("接收数据出错：", err)
		}
		log.Printf("Stream-Name: %s", response.Name)
		log.Printf("Stream-Bbox: %s", response.Bbox)
		log.Printf("Stream-Sbox: %s", response.Sbox)

		time.Sleep(time.Second * 1)
	}
}

func RequesTest() {

}

func StreamTest() {

}

//protoc -I grpcdemo6/proto/ grpcdemo6/proto/lynxi_nesting.proto --go_out=plugins=grpc:grpcdemo6/proto/
