package api

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/products-service/src/pkg/database"
	"github.com/szymon676/products-service/src/pkg/models"
)

func HandleCreateProduct(c *gin.Context) {
	var p models.BindProduct

	if err := c.BindJSON(&p); err != nil {
		c.String(400, "error binding product")
	}

	database.SaveProduct(p.Title, p.Price, p.Description)
	c.String(202, "product created")
}

func HandleGetProducts(c *gin.Context) {
	products, err := database.GetProducts()
	if err != nil {
		c.String(500, "error getting products")
	}
	c.JSON(200, gin.H{
		"products": products,
	})
}
