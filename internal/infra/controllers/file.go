package controllers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/buemura/temp-storage/internal/constants"
	"github.com/buemura/temp-storage/internal/domain/file"
	"github.com/buemura/temp-storage/internal/domain/session"
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

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	sessionId := form.Value["sessionId"][0]
	files := form.File["files"]

	val := client.Get("sessionId:" + sessionId)
	if val == "" {
		c.JSON(http.StatusNotFound, HttpResponse{
			"error": session.SessionNotFoundError,
		})
		return
	}

	sess := session.Session{}
	json.Unmarshal([]byte(val), &sess)

	for _, f := range files {
		filename := sessionId + "_" + filepath.Base(f.Filename)
		uploadDest := constants.UploadFileDest + filename
		file := file.NewFile(filename, uploadDest)
		sess.AddFile(file)
		c.SaveUploadedFile(f, file.FileUrl)
	}

	jsonBytes, err := json.Marshal(sess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	sessStr := string(jsonBytes)

	err = client.Set("sessionId:"+sess.ID, sessStr, sess.TimeToLive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, HttpResponse{
		"success":  "true",
		"uploaded": len(files),
	})
}
