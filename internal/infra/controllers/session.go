package controllers

import (
	"net/http"

	"github.com/buemura/temp-storage/internal/application"
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

	sService := application.NewSessionService(client)
	sess, err := sService.CreateSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, sess)
}
