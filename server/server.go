package main

import (
	"context"
	"fmt"
	pb "github.com/Crushaderdk/Mandatory1/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
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

/*
func postCourse(c *gin.Context) {
	var newCourse course
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

*/
func (s *server) postCourse(in *pb.PostCourseRequest) (*pb.PostCourseReply, error) {
	log.Printf("Received: %s, %s, %s, %s", in.Id, in.Name, in.Workload, in.Studentsatisfactionrating)
	workload, _ := strconv.ParseFloat(in.Workload, 64)
	studentSatisfactionRating, _ := strconv.ParseFloat(in.Studentsatisfactionrating, 64)
	var newCourse = course{ID: in.Id, Name: in.Name, Workload: workload, StudentSatisfactionRating: studentSatisfactionRating}
	courses = append(courses, newCourse)
	return &pb.PostCourseReply{}, nil
}

func (s *server) GetCourseByID(ctx context.Context, in *pb.GetCourseByIDRequest) (*pb.ReturnCourse, error) {
	log.Printf("Received: %v", in.GetRequest())

	for _, a := range courses {
		if a.ID == in.Request {
			returnString := fmt.Sprintf("Course: %s, Workload: %.1f, Student Satisfaction Rating: %.1f", a.Name, a.Workload, a.StudentSatisfactionRating)
			return &pb.ReturnCourse{Course: returnString}, nil
		}
	}

	return &pb.ReturnCourse{}, nil
}
