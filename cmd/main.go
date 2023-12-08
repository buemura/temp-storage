package main

import (
	"github.com/buemura/temp-storage/internal/infra/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/session", controllers.CreateSession)
	router.POST("/upload", controllers.UploadFiles)

	router.Run(":8080")
}
