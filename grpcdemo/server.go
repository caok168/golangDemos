package main

import (
	"context"
	"log"
	"net"

	pb "golangDemos/grpcdemo/proto"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{} //服务对象

// SayHello 实现服务的接口 在proto中定义的所有服务都是接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) RetrievalRepo(ctx context.Context, in *pb.RetrievalRepoRequest) (*pb.RetrievalRepoResponse, error) {
	retrievalRepo := &pb.RetrievalRepoResponse{}
	retrievalRepo.RepoId = 1
	retrievalRepo.RepoName = "人像库"
	retrievalRepo.PersonUuid = "33333333333333333333333"
	retrievalRepo.IdentityNumber = "130525999999999999"
	retrievalRepo.Age = 25

	return retrievalRepo, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //起一个服务
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterRetrievalServer(s, &server{})

	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	//reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
