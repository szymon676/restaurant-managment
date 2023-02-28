package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/szymon676/products-service/src/pkg/api/handlers"
	"github.com/szymon676/products-service/src/pkg/database"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	// ! setup
	r := gin.Default()
	// ! setup routes
	r.POST("/products", api.HandleCreateProduct)
	r.GET("/products", api.HandleGetProducts)

	// ? run server
	r.Run(":4001")
}
