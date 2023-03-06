package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/szymon676/products-service/internal/api/handlers"
	"github.com/szymon676/products-service/internal/database"
)

func init() {
	database.ConnectDB()
}

func SetupRoutes() {
	r := gin.Default()

	r.POST("/products", api.HandleCreateProduct)
	r.GET("/products", api.HandleGetProducts)
	r.PUT("/products/:id", api.HandleUpdateProduct)
	r.DELETE("/products/:id", api.HandleDeleteProduct)

	r.Run(":4001")
}
