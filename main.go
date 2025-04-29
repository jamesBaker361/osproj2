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
	responseQueue chan *pb.DispatcherResponse
}


func (s *DispatcherServer )  AcceptRequest(_ context.Context, disreq *pb.DispatcherRequest)  (*pb.DispatcherResponse,error) {
		job := <-s.responseQueue // take the next available job (blocking if empty)
		return job, nil
	}

func newDispatcherServer(buffer_size int) *DispatcherServer {
	s:=&DispatcherServer{
		responseQueue: make(chan *pb.DispatcherResponse, buffer_size),
	}
	return s
}

type ConsolidatorServer struct {
	pb.UnimplementedConsolidatorServiceServer
	primeQueue chan int
	timeQueue chan int
}

func (s * ConsolidatorServer) AcceptRequest(_ context.Context, conreq *pb.ConsolidatorRequest) (*pb.ConsolidatorResponse,error) {
	s.primeQueue <-int(conreq.NPrimes)
	fmt.Printf("primes %d\n",conreq.NPrimes)
	return &pb.ConsolidatorResponse{},nil
}

func newConsolidatorServer(primeQueue chan int,timeQueue chan int) *ConsolidatorServer {
	s:=&ConsolidatorServer{primeQueue:primeQueue,timeQueue:timeQueue}
	return s
}

type FilesystemServer struct {
	pb.UnimplementedFilesystemServiceServer
	FileName string
}

func (s *FilesystemServer ) AcceptRequest(_ context.Context, fsreq *pb.FilesystemRequest) ( *pb.FilesystemResponse,error) {
	//return &pb.FilesystemResponse{Data:[]byte("hello world")},nil 
	fileName:=s.FileName
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = file.Seek(int64(fsreq.StartingIndex), 0)
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fsreq.NBytes)
	bytesRead, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}
	buffer=buffer[:bytesRead]
	return &pb.FilesystemResponse{Data:buffer},nil
}

func newFilesystemServer(FileName string) *FilesystemServer {
	s:=&FilesystemServer{FileName:FileName}
	return s
}

func startDispatcherServer(d_port int,opts []grpc.ServerOption,server *DispatcherServer){
	d_lis, d_err := net.Listen("tcp", fmt.Sprintf("localhost:%d", d_port))
	if d_err != nil {
		log.Fatalf("failed to listen: %v", d_err)
	}
	log.Printf("gRPC server listening on port %d...", d_port)

	d_grpcServer := grpc.NewServer(opts...)
	pb.RegisterDispatcherServiceServer(d_grpcServer, server)
	d_grpcServer.Serve(d_lis)
}

func startConsolidatorServer(c_port int,opts []grpc.ServerOption,primeQueue chan int,timeQueue chan int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", c_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	log.Printf("gRPC server listening on port %d...", c_port)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterConsolidatorServiceServer(grpcServer, newConsolidatorServer(primeQueue,timeQueue))
	grpcServer.Serve(lis)
}

func startFilesystemServer(f_port int,opts []grpc.ServerOption,FileName string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", f_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	log.Printf("gRPC server listening on port %d...", f_port)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFilesystemServiceServer(grpcServer, newFilesystemServer(FileName))
	grpcServer.Serve(lis)
}


func main() {
	// command-line flags
	_N := flag.String("N", "64KB", "Number of something (int)")
	_C := flag.String("C", "1KB", "Another parameter (int)")
	dataPath := flag.String("data", "", "Path to the data file")
	configPath := flag.String("config", "", "Path to the config file")

	// Parse the flags
	flag.Parse()


	//N in {1KB, 32KB, 64KB, 256KB, 1MB, 64MB}; C in {64B, 1KB, 4KB, 8KB}.
	N_dict := make(map[string]int)
	N_dict["1KB"] = 1024
	N_dict["2KB"] = 2 * 1024 //testing only: delete this!
	N_dict["32KB"] = 32 * 1024
	N_dict["64KB"] = 64 * 1024
	N_dict["256KB"] = 256 * 1024
	N_dict["1MB"] = 1024 * 1024
	N_dict["64MB"] = 64 * 1024 * 1024

	C_dict := make(map[string]int)
	C_dict["64B"] = 64
	C_dict["128B"] = 128 //testing only: delete this!
	C_dict["1KB"] = 1024
	C_dict["4KB"] = 4 * 1024
	C_dict["8KB"] = 8 * 1024

	N, existsN := N_dict[*_N]

	if !existsN {
		N = 64 * 1024
	}
	C, existsC := C_dict[*_C]

	if !existsC {
		C = 1024
	}

	fmt.Println("N:", N)
	fmt.Println("C:", C)
	fmt.Println("Data file path:", *dataPath)
	fmt.Println("Config file path:", *configPath)

	// Get file info
	fileInfo, err := os.Stat(*dataPath)
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}

	// Get file size in bytes
	fileSize := fileInfo.Size()

	// Print the file size
	fmt.Printf("File size of %s is %d bytes\n", *dataPath, fileSize)
	
	total_jobs :=int(fileSize) / N
	fmt.Printf("Total jobs: %d\n",total_jobs)

	file, err := os.Open(*configPath)
	if err != nil {
		fmt.Println("Error opening config file:", err)
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
	var f_port int = ports[2]
	var opts []grpc.ServerOption

	//Dispatcher
	dispatcher_server:=newDispatcherServer(total_jobs+1)
	//responseQueue=dispatcher_server.responseQueue
	for i:=0; i< total_jobs;i+=1 {
		response:=&pb.DispatcherResponse{JobId:int32(i),StartingIndex:int32(i * N),EndingIndex:int32((i+1) * N)}
		dispatcher_server.responseQueue<-response
	}

	go startDispatcherServer(d_port,opts,dispatcher_server)

	//Consolidator
	primeQueue:=make(chan int, total_jobs+1)
	timeQueue:=make(chan int,total_jobs+1)
	go startConsolidatorServer(c_port,opts,primeQueue,timeQueue)

	
	//FileServer
	go startFilesystemServer(f_port,opts,*dataPath)

	count:=0
	total_primes:=0
	for count<total_jobs{
		primes:=<-primeQueue
		total_primes+=primes
		count++
	}

	fmt.Printf("Total Primes %d\n",total_primes)
	

	select {}
}
