package main

import (
	"context"
	"fmt"
	pb "golangDemos/grpcdemo4/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterFaceRecognitionServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to server ", err)
	}

	log.Println("complete")
}

type server struct{}

var index uint32 // test ONLY

func (s *server) FaceExtract(ctx context.Context, request *pb.FaceExtractRequest) (*pb.FaceExtractReply, error) {
	var r pb.FaceExtractReply

	println(index)

	r = pb.FaceExtractReply{
		TaskId: request.TaskId,
	}

	r.StatusOrResult = &pb.FaceExtractReply_Status{&pb.Status{State: "stopped", Code: 2000}}

	//r.StatusOrResult = &pb.FaceExtractReply_Result{&pb.FaceExtractResult{Uuid: "lasejflasjdfl", Attrs: map[string]string{"age": "12"}}}
	fmt.Printf("%v\n", r)

	time.Sleep(time.Millisecond * 15)

	return &r, nil
}

//
//func (s *server) FaceCompare(ctx context.Context, request *pb.FaceCompareRequest) (*pb.FaceCompareReply, error) {
//
//	return nil, nil
//}
//
//func (s *server) VideoRecognize(request *pb.VideoRecognizeRequest, serv pb.FaceRecognition_VideoRecognizeServer) error {
//	for i := 0; i < 10; i++ {
//		r := pb.VideoRecognizeReply{
//			TaskId: request.TaskId,
//		}
//
//		if index%3 == 0 {
//			r.StatusOrResultOrBoxes = &pb.VideoRecognizeReply_Status{&pb.Status{State: "stopped", Code: 2000}}
//		} else if index%3 == 1 {
//			r.StatusOrResultOrBoxes = &pb.VideoRecognizeReply_Result{&pb.VideoRecognizeResult{FaceUuid: "lsdfjlsf", TimestampOrCaptureTime: &pb.VideoRecognizeResult_Timestamp{Timestamp: 2323}}}
//		} else {
//			boxes := pb.BoxInfos{}
//			boxes.BoxInfos = append(boxes.BoxInfos, &pb.BoxInfos_BoxInfo{X0: 0, Y0: 0, X1: 100, Y1: 100})
//			boxes.BoxInfos = append(boxes.BoxInfos, &pb.BoxInfos_BoxInfo{X0: 110, Y0: 110, X1: 200, Y1: 200})
//
//			r.StatusOrResultOrBoxes = &pb.VideoRecognizeReply_BoxInfos{BoxInfos: &boxes}
//		}
//
//		index++
//
//		serv.Send(&r)
//
//		time.Sleep(time.Second)
//	}
//
//	return nil
//}
