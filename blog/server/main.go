package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/KarasunoAs9/gRPC-Project/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed listen: %v\n", err)
	}

	log.Printf("Listening on addres: %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed load server: %v\n", err)
	}

}