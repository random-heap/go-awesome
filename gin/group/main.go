package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/hello1", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message":"/v1/hello1!",
			})
		})
		v1.GET("/hello2", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/v1/hello2!",
			})
		})
		v1.GET("/hello3", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/v1/hello3!",
			})
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/hello1", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message":"/v2/hello1!",
			})
		})
		v2.GET("/hello2", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/v2/hello2!",
			})
		})
		v2.GET("/hello3", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "/v2/hello3!",
			})
		})
	}

	router.Run(":8080")
}
