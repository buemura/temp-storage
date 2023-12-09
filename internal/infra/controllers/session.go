package controllers

import (
	"net/http"

	"github.com/buemura/temp-storage/internal/application"
	"github.com/buemura/temp-storage/internal/infra/cache/redis"
	"github.com/gin-gonic/gin"
	r "github.com/redis/go-redis/v9"
)

func GetSession(c *gin.Context) {
	cli := r.NewClient(&r.Options{
		Addr: "localhost:6379",
	})
	client := redis.NewRedisCacheStorage(cli)
	defer client.Close()

	sessionId := c.Param("sessionId")

	sService := application.NewSessionService(client)
	sess, err := sService.GetSession(sessionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, sess)
}

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
