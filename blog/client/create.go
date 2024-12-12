package main

import (
	"context"
	"io"
	"log"

	pb "github.com/KarasunoAs9/gRPC-Project/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Alex",
		Title:    "Best Blog",
		Content:  "This is my first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog hase been created: %s\n", res)
	return res.Id
}

func readBlog(c pb.BlogServiceClient) *pb.Blog {
	id := &pb.BlogId{
		Id: "675a9fb187808480ec299859",
	}

	res, err := c.ReadBlog(context.Background(), id)
	if err != nil {
		log.Fatalf("Error with ReadBlog func: %v\n", err)
	}
	log.Printf("Your Blog: %v\n", res)
	return res
}

func updateBlog(c pb.BlogServiceClient, id string) {
	data := &pb.Blog{
		Id:       id,
		AuthorId: "Pasha",
		Title:    "Best Marketing",
		Content:  "Why marketiong is the best job ever",
	}

	_, err := c.UpdateBlog(context.Background(), data)
	if err != nil {
		log.Fatalf("Error with UpdateBlog func: %v\n", err)
	}

	log.Println("Blog was update sucsessfully")
}

func deleteBlog(c pb.BlogServiceClient, id string) {
	oid := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), oid)
	if err != nil {
		log.Fatalf("Error with deleting blog: %v\n", err)
	}

	log.Println("Blog was delete sucsessfully")
}

func listBlogs(c pb.BlogServiceClient) {

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error with litsBlogs func: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Blogs: %v\n", msg)
	}
}
