package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/szymon676/workers-service/internal/api/handlers"
	"github.com/szymon676/workers-service/internal/database"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	r := gin.Default()

	r.POST("/workers", api.HandleCreateWorker)
	r.GET("/workers", api.HandleGetWorkers)
	r.PUT("/workers/:id", api.HandleUpdateWorker)
	r.DELETE("/workers/:id", api.HandleDeleteWorker)

	r.Run(":4000")
}
