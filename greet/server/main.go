package main

import (
	"log"
	"net"

	pb "github.com/KarasunoAs9/gRPC-Project/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed listen: %v\n", err)
	}

	log.Printf("Listening on addres: %v\n", addr)

	s := grpc.NewServer()
	
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed load server: %v\n", err)
	}

}