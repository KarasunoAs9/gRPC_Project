package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/KarasunoAs9/gRPC-Project/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	createBlog(c)
	readBlog(c)
	updateBlog(c, "675a9fb187808480ec299859")
	readBlog(c)
	listBlogs(c)
	deleteBlog(c, "675a9fb187808480ec299859")
}
