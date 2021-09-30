package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type course struct {
	ID                        int     `json:"id"`
	Name                      string  `json:"name"`
	Workload                  float64 `json:"workload"`
	StudentSatisfactionRating float64 `json:"studentSatisfactionRating"`
}

var courses = []course{
	{ID: 1, Name: "Distributed Systems", Workload: 7.5, StudentSatisfactionRating: 0.8},
	{ID: 1, Name: "Introduction to Database Systems", Workload: 7.5, StudentSatisfactionRating: 0.79},
	{ID: 1, Name: "Analysis, Design and Software Architecture", Workload: 15, StudentSatisfactionRating: 0.7},
}

func getCourse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourse)
	router.POST("/courses", postCourse)

	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body.
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
