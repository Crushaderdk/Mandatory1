package main

import (
	"context"
	pb "github.com/Crushaderdk/Mandatory1/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address   = "localhost:8080"
	defaultID = "1"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCourseServiceClient(conn)

	// Contact the server and print out its response.
	id := defaultID
	if len(os.Args) > 1 {
		id = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCourseByID(ctx, &pb.GetCourseByIDRequest{Request: id})
	if err != nil {
		log.Fatalf("could not output course: %v", err)
	}
	log.Printf("Greeting: %s", r.GetCourse())
}
