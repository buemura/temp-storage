package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/buemura/temp-storage/internal/domain/session"
	"github.com/buemura/temp-storage/internal/infra/cache/redis"
	"github.com/gin-gonic/gin"
	r "github.com/redis/go-redis/v9"
)

func CreateSession(c *gin.Context) {
	cli := r.NewClient(&r.Options{
		Addr: "localhost:6379",
	})
	client := redis.NewRedisCacheStorage(cli)
	defer client.Close()

	sess := session.NewSession(10, 10)
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

	c.JSON(http.StatusCreated, sess)
}
