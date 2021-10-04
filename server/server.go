package main

import (
	"context"
	pb "github.com/Crushaderdk/Mandatory1/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	port = ":8080"
)

type course struct {
	ID                        string  `json:"id"`
	Name                      string  `json:"name"`
	Workload                  float64 `json:"workload"`
	StudentSatisfactionRating float64 `json:"studentSatisfactionRating"`
}

var courses = []course{
	{ID: "1", Name: "Distributed Systems", Workload: 7.5, StudentSatisfactionRating: 0.8},
	{ID: "2", Name: "Introduction to Database Systems", Workload: 7.5, StudentSatisfactionRating: 0.79},
	{ID: "3", Name: "Analysis, Design and Software Architecture", Workload: 15, StudentSatisfactionRating: 0.7},
}

func getCourse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

type server struct {
	pb.UnimplementedCourseServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func postCourse(c *gin.Context) {
	var newCourse course
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

/*func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}*/

func (s *server) GetCourseByID(ctx context.Context, in *pb.GetCourseByIDRequest) (*pb.ReturnCourse, error) {
	log.Printf("Received: %v", in.GetRequest())
	return &pb.ReturnCourse{Course: pb.GetCourseByIDRequest{}}, nil
}
