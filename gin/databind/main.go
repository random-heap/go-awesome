package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type People struct {
	ID string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	router := gin.Default()

	router.POST("/login", func(context *gin.Context) {
		var login Login
		if err := context.ShouldBindJSON(&login); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":    "500",
				"message": "failed",
				"data":    err.Error(),
			})
			return
		}

		if login.Name == "Jim" && login.Password == "123" {

			context.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "success",
				"data":    "用户登录成功",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "success",
			"data":    "用户登录失败",
		})
	})

	router.POST("/loginForm", func(context *gin.Context) {
		var form Login

		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":    "500",
				"message": "failed",
				"data":    err.Error(),
			})
			return
		}

		if form.Name == "Jim" && form.Password == "123" {

			context.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "success",
				"data":    "用户登录成功",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "success",
			"data":    "用户登录失败",
		})
	})

	router.GET("/person", func(context *gin.Context) {
		var person Person
		if context.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
		}
	})

	router.GET("/people/:name/:id", func(context *gin.Context) {
		var people People
		if err := context.ShouldBindUri(&people); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":    "500",
				"message": "failed",
				"data":    err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "success",
			"data":    people,
		})
	})

	router.Run(":8080")

}
