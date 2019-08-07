package main

import (
	"context"
	pb "golangDemos/grpcdemo6/proto1"
	pb2 "golangDemos/grpcdemo6/proto2"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//StreamTest2()

	//StreamTest_New()

	//Exec_test()

	imagePath := "/home/ck/go/src/golangDemos/grpcdemo6/images/123.jpg"
	Test_Image(imagePath)
}

//func StreamTest2() {
//	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure())
//	if err != nil {
//		log.Printf("连接失败：【%v】", err)
//		return
//	}
//	defer conn.Close()
//
//	// 声明客户端
//	client := pb.NewGreeterClient(conn)
//	// 声明 context
//	ctx := context.Background()
//
//	//
//
//	stream, err := client.StreamSayHello(ctx)
//
//	if err != nil {
//		log.Printf("创建数据流失败：【%v】", err)
//	}
//
//	go func() {
//		log.Printf("请输入消息...")
//		in := bufio.NewReader(os.Stdin)
//
//		for {
//			inputStr, _ := in.ReadString(' ')
//			request := &pb.HelloRequest{Name: inputStr}
//			if err := stream.Send(request); err != nil {
//				return
//			}
//		}
//	}()
//
//	for {
//		response, err := stream.Recv()
//		if err == io.EOF {
//			log.Println("收到服务器的结束信号")
//			break
//		}
//		if err != nil {
//			log.Println("接收数据出错：", err)
//		}
//		log.Printf("[客户端收到]: %s", response.Name)
//	}
//
//}

func RequesTest() {

}

func StreamTest() {
	//conn, err := grpc.Dial("192.168.11.5:50051", grpc.WithInsecure())
	//if err != nil {
	//	log.Printf("连接失败：【%v】", err)
	//	return
	//}
	//defer conn.Close()
	//
	//// 声明客户端
	//client := pb.NewGreeterClient(conn)
	//// 声明 context
	//ctx := context.Background()
	//
	//request := &pb.HelloRequest{
	//	Name: "gdh",
	//}
	//
	//// 创建双向流
	//stream, err := client.SayHello(ctx, request)
	//if err != nil {
	//	log.Printf("创建数据流失败：【%v】", err)
	//}
	//
	//fmt.Println("Name:", stream.Name)
	//fmt.Println("Bbox:", stream.Bbox)
	//fmt.Println("Sbox:", stream.Sbox)
	//
	//streamList, err := client.LstSayHello(ctx, request)
	//if err != nil {
	//	log.Printf("创建数据流111失败：【%v】", err)
	//}
	//
	//for {
	//	response, err := streamList.Recv()
	//	if err == io.EOF {
	//		log.Println("收到服务器的结束信号")
	//		break
	//	}
	//	if err != nil {
	//		log.Println("接收数据出错：", err)
	//	}
	//	log.Printf("Stream-Name: %s", response.Name)
	//	log.Printf("Stream-Bbox: %s", response.Bbox)
	//	log.Printf("Stream-Sbox: %s", response.Sbox)
	//
	//	time.Sleep(time.Second * 1)
	//}
}

func StreamTest_New() {
	conn, err := grpc.Dial("192.168.9.31:50059", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：【%v】", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := pb.NewSmartClassClient(conn)
	// 声明 context
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	frameExtractFrequency := &pb.FrameExtractFrequency{
		Pose:   1,
		Detect: 2,
		Face:   3,
	}

	content := &pb.PipelineRequest_FrameExtractFrequency{
		FrameExtractFrequency: frameExtractFrequency,
	}

	request := &pb.PipelineRequest{
		PipelineName:           "111",
		DeviceDesc:             1,
		PipelineRequestContent: content,
	}

	// 创建双向流
	stream, err := client.StartPipeline(ctx, request)
	if err != nil {
		log.Printf("创建数据流失败：【%v】", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("收到服务器的结束信号")
			break
		}
		if err != nil {
			log.Println("接收数据出错：", err)
		}
		log.Printf("PipelineReplyContent: %v", response.PipelineReplyContent)

		if response.GetStatus().Code == 0 {
			cancel()
		}

		time.Sleep(time.Second * 1)
	}
}

func Exec_test() {
	conn, err := grpc.Dial("192.168.9.31:50059", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：【%v】", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := pb.NewSmartClassClient(conn)
	// 声明 context
	ctx := context.Background()

	smartClass := &pb.SmartClassInfo{
		ConfigFilePath: "/home/lthpc/Workspace/ClassVideos/VID_20190624_152017.mp4",
	}

	content := &pb.TaskRequest_SmartClass{
		SmartClass: smartClass,
	}

	taskRequest := &pb.TaskRequest{
		TaskId:             1,
		PipelineName:       "pipelineName",
		TaskRequestContent: content,
	}

	stream, err := client.Exec(ctx, taskRequest)
	if err != nil {
		log.Printf("创建数据流失败：【%v】", err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("收到服务器的结束信号")
			break
		}
		if err != nil {
			log.Println("接收数据出错：", err)
		}
		log.Println("taskId: ", response.TaskId)
		log.Println("TaskReplyContent: ", response.TaskReplyContent)

		time.Sleep(time.Second * 1)
	}

}

func Test_Image(fileName string) {

	imageByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("error:", err.Error())
		panic("File " + fileName + " not found")
	}

	conn, err := grpc.Dial("192.168.11.5:50052", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败：【%v】", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := pb2.NewFaceClient(conn)
	// 声明 context
	ctx := context.Background()

	request := &pb2.TaskRequest{
		TaskId:       1,
		PipelineName: "gdh",
		Uuid:         "sfdsfdsf",
		Rawinput:     imageByte,
	}

	result, err := client.FaceData(ctx, request)
	if err != nil {
		log.Println("出错：", err.Error())
		return
	}

	log.Println(result)

}

//protoc -I grpcdemo6/proto/ grpcdemo6/proto/lynxi_nesting.proto --go_out=plugins=grpc:grpcdemo6/proto/
