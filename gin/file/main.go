package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(context *gin.Context) {
		// 单文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的路径
		context.SaveUploadedFile(file, "C:/data0/"+file.Filename)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.POST("/multiupload", func(context *gin.Context) {
		// 多文件
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			context.SaveUploadedFile(file, "C:/data0/"+file.Filename)
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}
