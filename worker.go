package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"context"
	"time"
	"math"
	"bytes"
	"encoding/binary"
	"io"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/examples/data"
	//"google.golang.org/protobuf/proto"
	pb"project/grpc/proto"
)

var wg sync.WaitGroup

func isPrime(n uint64) bool {
	if n <= 1 {
		return false
	} else if n == 2 || n == 3 || n%2 == 0 {
		return true
	}

	for i := uint64(3); i <= uint64(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getPrimes(numbers []uint64,prime_channel chan int) {
	defer wg.Done()
	result := 0
	for _, n := range numbers {
		if isPrime(n) {
			result++
		}

	}
	//fmt.Printf("prime results %d\n",result)
	prime_channel<-result
}

func readAllUvarints(buf []byte) ([]uint64, error) {
	var numbers []uint64
	reader := bytes.NewReader(buf) // Create a reader from the byte slice
	for {
		var num uint64
		err := binary.Read(reader, binary.LittleEndian, &num)
		if err == io.EOF {
			break // Stop when reaching the end of the buffer
		}
		if err != nil {
			return nil, err // Return error if reading fails
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func sendDispatcherRequest(client pb.DispatcherServiceClient) (*pb.DispatcherResponse, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &pb.DispatcherRequest{JobId: 1}
	response, err := client.AcceptRequest(ctx, request)
	if err != nil {
		log.Fatalf("sendDispatcherRequest failed: %v", err)
		return &pb.DispatcherResponse{},err
	} else {
		return response,err
	}
}

func sendFilesystemRequest(client pb.FilesystemServiceClient,startingIndex int32, nBytes int32 ) (*pb.FilesystemResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request:=&pb.FilesystemRequest{NBytes:nBytes,StartingIndex:startingIndex }
	response,err:=client.AcceptRequest(ctx,request)
	if err != nil {
		log.Fatalf("sendDispatcherRequest failed: %v", err)
		return &pb.FilesystemResponse{},err
	} else {
		return response,err
	}

}

func sendConsolidatorRequest(client pb.ConsolidatorServiceClient, nPrimes int,start int32){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	end:= int32(time.Now().Unix())
	request:=&pb.ConsolidatorRequest{NPrimes:int32(nPrimes),TimeElapsed:end-start}
	_,err:=client.AcceptRequest(ctx,request)
	if err != nil {
		log.Fatalf("sendConsolidatorRequest failed: %v", err)
	} 
}

func main() {
	// command-line flags
	_C := flag.String("C", "1KB", "Chunk size")
	configPath := flag.String("config", "", "Path to the config file")
	//N in {1KB, 32KB, 64KB, 256KB, 1MB, 64MB}; C in {64B, 1KB, 4KB, 8KB}.
	// Parse the flags
	flag.Parse()

	C_dict := make(map[string]int)
	C_dict["64B"] = 64
	C_dict["128B"] = 128 //testing only: delete this!
	C_dict["1KB"] = 1024
	C_dict["4KB"] = 4 * 1024
	C_dict["8KB"] = 8 * 1024


	C, existsC := C_dict[*_C]

	if !existsC {
	fmt.Println("no C??")
		C = 1024
	}

	

	fmt.Println("C:", C)
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
	var f_port int = ports[2]
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	dispatcherServerAddr := fmt.Sprintf("localhost:%d", d_port)

	conn, err := grpc.NewClient(dispatcherServerAddr, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	start := int32(time.Now().Unix())
	client:=pb.NewDispatcherServiceClient(conn)
	response,err:=sendDispatcherRequest(client)
	if err!=nil{
		log.Fatalf("sendDispatcherRequest failed: %v", err)
	}else{
		fmt.Printf("Received response: JobId=%d,  StartingIndex=%d, EndingIndex=%d\n",
			response.JobId, response.StartingIndex,response.EndingIndex)
		filesystemServerAddr := fmt.Sprintf("localhost:%d", f_port)
		f_conn, f_err := grpc.NewClient(filesystemServerAddr, opts...)
		if f_err != nil {
			log.Fatalf("fail to grpc.NewClient(filesystemServerAddr): %v", f_err)
		}
		defer f_conn.Close()
		f_client:=pb.NewFilesystemServiceClient(f_conn)

		
		prime_channel := make(chan int, 1+((response.EndingIndex-response.StartingIndex)/int32(C)))
		for j:=response.StartingIndex; j<  response.EndingIndex; j+=int32(C){
			f_response,f_err:=sendFilesystemRequest(f_client,j,int32(C))
			if f_err != nil {
				log.Fatalf("fail to fs reques: %v",f_err)
			}
			//fmt.Printf("Received fs response: data0=%d,  \n",f_response.Data[0])
			numbers,byte_read_err:=readAllUvarints(f_response.Data)
			if byte_read_err != nil {
				log.Fatalf("byte read error: %v",byte_read_err)
			}
			wg.Add(1)
			go getPrimes(numbers, prime_channel)
			
		}
		wg.Wait()
		close(prime_channel)
		prime_count:=0
		for prime:=range(prime_channel) {
			prime_count+=prime
		}
		fmt.Printf("total primes %d\n",prime_count)
		consolidatorServerAddr:=fmt.Sprintf("localhost:%d", c_port)
		c_conn, c_err := grpc.NewClient(consolidatorServerAddr, opts...)

		if c_err != nil {
			log.Fatalf("fail to connect grpc.NewClient(consolidatorServerAddr: %v", c_err)
		}
		defer c_conn.Close()
		c_client:=pb.NewConsolidatorServiceClient(c_conn)
		sendConsolidatorRequest(c_client,prime_count,start)

	}
}
