package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/KarasunoAs9/gRPC-Project/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked: %v\n", in)

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, fmt.Errorf("internal error: %v", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("cannot convert to oid")
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, fmt.Errorf("error with reading blog: %v", err)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, fmt.Errorf("cannot find blog with this id")
	}

	return documentToBlog(data), nil

}

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, fmt.Errorf("error with id: %v", err)
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, fmt.Errorf("error with updating: %v", err)
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("cannot find blog with id")
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog was invoked: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, fmt.Errorf("errow with id: %v", err)
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, fmt.Errorf("errow with deleting blog: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListBlogs(_ *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	log.Printf("List blog was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return fmt.Errorf("error with find data: %v", err)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err = cur.Decode(data)

		if err != nil {
			return fmt.Errorf("error with decode blog: %v", err)
		}
		stream.Send(documentToBlog(data))
	}

	return nil
}
