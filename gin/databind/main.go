package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
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

	router.Run(":8080")

}
