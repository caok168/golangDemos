package main

import (
	"context"
	"flag"
	"fmt"
	pb "golangDemos/grpcdemo3/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = 50051
)

type faceTaskServer struct {
}

func newServer() *faceTaskServer {
	s := &faceTaskServer{}
	return s
}

func (ft *faceTaskServer) FaceExtract(ctx context.Context, request *pb.FaceExtractRequest) (*pb.FaceExtractReply, error) {

	response := &pb.FaceExtractReply{}
	response.TaskId = 2

	result := &pb.FaceExtractResult{}
	result.Type = 3
	result.Uuid = "cccccccccccccccccccccc"

	status := &pb.Status{}
	status.Code = 100
	status.State = "OK"

	response.StatusOrResult = &pb.FaceExtractReply_Status{Status: status}

	return response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEngineServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
