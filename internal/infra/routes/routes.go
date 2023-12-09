package routes

import (
	"github.com/buemura/temp-storage/internal/infra/container"
	"github.com/buemura/temp-storage/internal/infra/controllers"
	"github.com/gin-gonic/gin"
)

func SetupGinRoutes(router *gin.Engine) {
	cache := container.LoadCacheStorage()
	sService := container.LoadSessionService(cache)

	sController := controllers.NewSessionController(sService)
	fController := controllers.NewFileController(sService)

	router.POST("/session", sController.CreateSession)
	router.GET("/session/:sessionId", sController.GetSession)
	router.POST("/upload", fController.UploadFiles)
}
