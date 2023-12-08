package main

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/buemura/temp-storage/internal/constants"
	"github.com/buemura/temp-storage/internal/domain/file"
	"github.com/buemura/temp-storage/internal/domain/session"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type HttpResponse map[string]interface{}

// var client *redis.Client

func main() {
	// client = connectRedis()

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/session", createSession)
	router.POST("/upload", uploadFiles)
	router.Run(":8080")
}

func createSession(c *gin.Context) {
	client := connectRedis()
	defer client.Close()
	ctx := context.Background()

	sess := session.NewSession(10, 10)
	jsonBytes, err := json.Marshal(sess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	sessStr := string(jsonBytes)

	err = client.Set(ctx, "sessionId:"+sess.ID, sessStr, time.Duration(sess.TimeToLive)*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, sess)
}

func uploadFiles(c *gin.Context) {
	client := connectRedis()
	defer client.Close()
	ctx := context.Background()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	sessionId := form.Value["sessionId"][0]
	files := form.File["files"]

	val, err := client.Get(ctx, "sessionId:"+sessionId).Result()
	if err != nil {
		if val == "" {
			c.JSON(http.StatusNotFound, HttpResponse{
				"error": session.SessionNotFoundError,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
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

	err = client.Set(ctx, "sessionId:"+sess.ID, sessStr, time.Duration(sess.TimeToLive)*time.Minute).Err()
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

func connectRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
