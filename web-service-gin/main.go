package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	router := gin.Default()
	router.GET("/courses", getCourse)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourse)

	router.Run("localhost:8080")
}

// postCourses adds an album from JSON received in the request body.
func postCourse(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to
	// newCourse.
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new album to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

// getCourseByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of courses, looking for
	// an album whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}
