package controllers

import (
	"net/http"

	"github.com/buemura/temp-storage/internal/application"
	"github.com/gin-gonic/gin"
)

type SessionController struct {
	sService *application.SessionService
}

func NewSessionController(sService *application.SessionService) *SessionController {
	return &SessionController{
		sService: sService,
	}
}

func (ss *SessionController) GetSession(c *gin.Context) {
	sessionId := c.Param("sessionId")
	sess, err := ss.sService.GetSession(sessionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, sess)
}

func (ss *SessionController) CreateSession(c *gin.Context) {
	sess, err := ss.sService.CreateSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, sess)
}
