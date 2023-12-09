package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/buemura/temp-storage/internal/application"
	"github.com/buemura/temp-storage/internal/constants"
	"github.com/buemura/temp-storage/internal/domain/file"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	sService *application.SessionService
}

func NewFileController(sService *application.SessionService) *FileController {
	return &FileController{
		sService: sService,
	}
}

func (fs *FileController) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}

	sessionId := form.Value["sessionId"][0]
	files := form.File["files"]

	sess, err := fs.sService.GetSession(sessionId)
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

	result, err := fs.sService.UpdateSession(sess)
	if err != nil {
		c.JSON(http.StatusBadRequest, HttpResponse{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
