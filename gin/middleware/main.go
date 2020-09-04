package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {

		t := time.Now()

		context.Set("name", "Jim")

		context.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := context.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())

	router.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string)

		log.Println(name)
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

