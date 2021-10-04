package main

import (
	"context"
	"fmt"
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
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if os.Args[1] == "getById" {
			id = os.Args[2]

			r, err := c.GetCourseByID(ctx, &pb.GetCourseByIDRequest{Request: id})
			if err != nil {
				log.Fatalf("could not output course: %v", err)
			}
			log.Printf(r.GetCourse())

		} else if os.Args[1] == "post" {
			r, err := c.PostCourse(ctx, &pb.PostCourseRequest{Id: os.Args[2], Name: os.Args[3], Workload: os.Args[4], Studentsatisfactionrating: os.Args[5]})
			if err != nil {
				log.Fatalf("400 wrong input: %v", err)
			}
			r.Reply = fmt.Sprintf("Course posted:\n Course ID: %s, Name: %s, Workload: %s, Student Satisfaction Rating: %s", os.Args[2], os.Args[3], os.Args[4], os.Args[5])
			log.Printf(r.Reply)
		} else if os.Args[1] == "getAll" {
			r, err := c.GetCourses(ctx, &pb.GetCoursesRequest{})
			if err != nil {
				log.Fatalf("could not output course: %v", err)
			}
			log.Printf(r.GetCourses())
		} else {
			log.Printf("400: bad request")
		}
	}
}
