package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/szymon676/workers-service/src/pkg/api/handlers"
	"github.com/szymon676/workers-service/src/pkg/database"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	// ! setup
	r := gin.Default()

	// ! setup routes
	r.POST("/workers", api.HandleCreateWorker)
	r.GET("/workers", api.HandleGetWorkers)

	r.Run(":4000")
}
