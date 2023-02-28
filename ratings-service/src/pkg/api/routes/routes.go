package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/ratings-service/src/pkg/api/handlers"
	"github.com/szymon676/ratings-service/src/pkg/database"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	// ! setup
	r := gin.Default()

	// ! setup routes
	r.GET("/ratings", handlers.HandleGetRatings)

	// ? run server
	r.Run(":4002")
}
