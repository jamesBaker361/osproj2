package main

import (
	"flag"
	"fmt"
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"bufio"
	"strconv"
	"os"
	"strings"
	//"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/examples/data"
	//"google.golang.org/protobuf/proto"
	pb"project/grpc/proto"
)




type DispatcherServer struct {
	pb.UnimplementedDispatcherServiceServer 
}


func (s *DispatcherServer )  AcceptRequest(_ context.Context, disreq *pb.DispatcherRequest)  (*pb.DispatcherResponse,error) {
		return &pb.DispatcherResponse{JobId:1,NChunks:1,StartingIndex:1},nil
	}

func newDispatcherServer() *DispatcherServer {
	s:=&DispatcherServer{}
	return s
}

type ConsolidatorServer struct {
	pb.UnimplementedConsolidatorServiceServer
}

func (s * ConsolidatorServer) AcceptRequest(_ context.Context, conreq *pb.ConsolidatorRequest) (*pb.ConsolidatorResponse,error) {
	return &pb.ConsolidatorResponse{},nil
}

func newConsolidatorServer() *ConsolidatorServer {
	s:=&ConsolidatorServer{}
	return s
}

type FilesystemServer struct {
	pb.UnimplementedFilesystemServiceServer
}

func (s *FilesystemServer ) AcceptRequest(_ context.Context, fsreq *pb.FilesystemRequest) ( *pb.FilesystemResponse,error) {
	return &pb.FilesystemResponse{Data:[]byte("hello world")},nil 

}

func startDispatcherServer(d_port int,opts []grpc.ServerOption){
	d_lis, d_err := net.Listen("tcp", fmt.Sprintf("localhost:%d", d_port))
	if d_err != nil {
		log.Fatalf("failed to listen: %v", d_err)
	}

	d_grpcServer := grpc.NewServer(opts...)
	pb.RegisterDispatcherServiceServer(d_grpcServer, newDispatcherServer())
	d_grpcServer.Serve(d_lis)
}

func startConsolidatorServer(c_port int,opts []grpc.ServerOption) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", c_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterConsolidatorServiceServer(grpcServer, newConsolidatorServer())
	grpcServer.Serve(lis)
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

	file, err := os.Open(*configPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize a slice to store the port values as integers
	var ports []int

	// Read lines and extract the third value (port)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		// Read only the first three lines
		if lineCount > 3 {
			break
		}

		// Split the line into parts (assuming space/tab separation)
		parts := strings.Fields(scanner.Text())

		// Check if there are at least 3 parts
		if len(parts) >= 3 {
			// Convert the third item (port) from string to int
			d_port, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Error converting port:", err)
				continue
			}

			// Append the port to the slice
			ports = append(ports, d_port)
		} else {
			fmt.Printf("Line %d does not contain enough parts.\n", lineCount)
		}
	}

	// Check for errors while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print the array of port values
	fmt.Println("Ports:", ports)

	var d_port int = ports[0]
	var c_port int = ports[1]
	//var f_port int = ports[2]
	var opts []grpc.ServerOption

	//Dispatcher
	//var d_port int = 5001
	startDispatcherServer(d_port,opts)

	//Consolidator
	//var c_port int = 5002
	startConsolidatorServer(c_port,opts)
	


	
	

	
	//FileServer

}
