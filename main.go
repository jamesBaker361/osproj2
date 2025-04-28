package main

import (
	"flag"
	"fmt"
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/examples/data"
	//"google.golang.org/protobuf/proto"
	pb"project/grpc/proto"
)

type DispatcherServer struct {
	pb.UnimplementedDispatcherServiceServer 
}

type ConsolidatorServer struct {
	pb.UnimplementedConsolidatorServiceServer
}

type FilesystemServer struct {
	pb.UnimplementedFilesystemServiceServer
}

func (s *DispatcherServer )  AcceptRequest(_ context.Context, disreq *pb.DispatcherRequest)  (*pb.DispatcherResponse,error) {
		return &pb.DispatcherResponse{JobId:1,NChunks:1,StartingIndex:1},nil
	}

func newDispatcherServer() *DispatcherServer {
	s:=&DispatcherServer{}
	return s
}


func main() {
	// command-line flags
	N := flag.Int("N", 0, "Number of something (int)")
	C := flag.Int("C", 0, "Another parameter (int)")
	dataPath := flag.String("data", "", "Path to the data file")
	configPath := flag.String("config", "", "Path to the config file")

	// Parse the flags
	flag.Parse()

	fmt.Println("N:", *N)
	fmt.Println("C:", *C)
	fmt.Println("Data file path:", *dataPath)
	fmt.Println("Config file path:", *configPath)

	//Consolidator



	
	//Dispatcher
	var port int = 100
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDispatcherServiceServer(grpcServer, newDispatcherServer())
	grpcServer.Serve(lis)

	
	//FileServer

}
