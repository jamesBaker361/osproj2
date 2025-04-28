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

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/examples/data"
	//"google.golang.org/protobuf/proto"
	pb"project/grpc/proto"
)

func sendDispatcherRequest(client pb.DispatcherServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &pb.DispatcherRequest{JobId: 1}
	response, err := client.AcceptRequest(ctx, request)
	if err != nil {
		log.Fatalf("sendDispatcherRequest failed: %v", err)
	} else {
		fmt.Printf("Received response: JobId=%d, NChunks=%d, StartingIndex=%d\n",
			response.JobId, response.NChunks, response.StartingIndex)
	}
}

func main() {
	// command-line flags
	C := flag.Int("C", 0, "Another parameter (int)")
	configPath := flag.String("config", "", "Path to the config file")

	// Parse the flags
	flag.Parse()

	fmt.Println("C:", *C)
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
	//var c_port int = ports[1]
	//var f_port int = ports[2]
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	serverAddr := fmt.Sprintf("localhost:%d", d_port)

	conn, err := grpc.NewClient(serverAddr, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client:=pb.NewDispatcherServiceClient(conn)
	sendDispatcherRequest(client)
}
