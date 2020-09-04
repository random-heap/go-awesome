package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name string	`json:"name"`
	Age string	`json:"age"`
	Gender string	`json:"gender"`
	Class string	`json:"class"`
}

func main() {

	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message":"hello gin!",
		})
	})

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")
		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})


	router.POST("/form/post", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		gender := context.PostForm("gender")
		class := context.PostForm("class")

		student := Student{}
		student.Name = name
		student.Age = age
		student.Gender = gender
		student.Class = class

		context.JSON(http.StatusOK, gin.H{
			"code":"200",
			"message":"success",
			"data": student,
		})
	})

	router.POST("/json/post", func(context *gin.Context) {

		student := &Student{}
		if err := context.BindJSON(student); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":"500",
				"message":"error",
				"data": nil,
			})
		}

		context.JSON(http.StatusOK, gin.H{
			"code":"200",
			"message":"success",
			"data": student,
		})
	})

	router .Run()
}
