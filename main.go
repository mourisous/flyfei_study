package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var students = []Student{
	{
		ID:    1,
		Name:  "Zhang San",
		Age:   18,
		Grade: 1,
	},
	{
		ID:    2,
		Name:  "Li Si",
		Age:   19,
		Grade: 2,
	},
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.Printf("%s %s ip=%s status=%d duration=%v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "Bearer 123456" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		c.Next()
	}
}

func getStudentInfo(c *gin.Context) {
	time.Sleep(2 * time.Second)
	c.JSON(http.StatusOK, students[0])
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(loggingMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	studentGroup := router.Group("/students")
	studentGroup.Use(authMiddleware())
	{
		studentGroup.GET("/info", getStudentInfo)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
