package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/buemura/temp-storage/internal/application"
	"github.com/buemura/temp-storage/internal/constants"
	"github.com/buemura/temp-storage/internal/domain/file"
	"github.com/buemura/temp-storage/internal/infra/cache/redis"
	"github.com/gin-gonic/gin"
	r "github.com/redis/go-redis/v9"
)

func UploadFiles(c *gin.Context) {
	cli := r.NewClient(&r.Options{
		Addr: "localhost:6379",
	})
	client := redis.NewRedisCacheStorage(cli)
	defer client.Close()

	sService := application.NewSessionService(client)

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	sessionId := form.Value["sessionId"][0]
	files := form.File["files"]

	sess, err := sService.GetSession(sessionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	for _, f := range files {
		filename := sessionId + "_" + filepath.Base(f.Filename)
		uploadDest := constants.UploadFileDest + filename
		file := file.NewFile(filename, uploadDest)
		sess.AddFile(file)
		c.SaveUploadedFile(f, file.FileUrl)
	}

	result, err := sService.UpdateSession(sess)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
