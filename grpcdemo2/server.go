package main

import (
	"flag"
	"fmt"
	pb "golangDemos/grpcdemo2/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type routeGuideServer struct {
	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pb.RouteNote
}

func serialize(point *pb.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	return s
}

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for {
		time.Sleep(time.Second * 1)

		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)

		s.mu.Lock()
		s.routeNotes[key] = append(s.routeNotes[key], in)
		// Note: this copy prevents blocking other clients while serving this one.
		// We don't need to do a deep copy, because elements in the slice are
		// insert-only and never modified.
		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		i := 0

		for index, note := range rn {

			i++
			fmt.Println("i=", i)

			time.Sleep(time.Second * 2)
			note.Message = note.Message + " === " + strconv.Itoa(index)

			fmt.Println("note:", note.Message, note.Location)
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
