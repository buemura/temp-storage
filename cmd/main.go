package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/buemura/temp-storage/internal/constants"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			filename := filepath.Base(file.Filename)
			uploadDest := constants.UploadFileDest + filename
			c.SaveUploadedFile(file, uploadDest)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"success":  "true",
			"uploaded": len(files),
		})
	})

	router.Run(":8080")
}
