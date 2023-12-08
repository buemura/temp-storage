package routes

import (
	"github.com/buemura/temp-storage/internal/infra/controllers"
	"github.com/gin-gonic/gin"
)

func SetupGinRoutes(router *gin.Engine) {
	router.POST("/session", controllers.CreateSession)
	router.POST("/upload", controllers.UploadFiles)
}
