package main

import (
	"github.com/buemura/temp-storage/internal/infra/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.MaxMultipartMemory = 8 << 20 // 8 MiB
	routes.SetupGinRoutes(server)
	server.Run(":8080")
}
