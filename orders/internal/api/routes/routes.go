package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/orders-service/internal/api/handlers"
	"github.com/szymon676/orders-service/internal/database"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	r := gin.Default()

	r.POST("/orders", handlers.HandleCreateOrder)
}
