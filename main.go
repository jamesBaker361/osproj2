package main

import (
	"flag"
	"fmt"
	pb"grpc/proto"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
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
	//FileServer

}
