package api

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/products-service/internal/database"
	"github.com/szymon676/products-service/internal/models"
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

func HandleUpdateProduct(c *gin.Context) {
	var p models.BindProduct
	id := c.Param("id")

	if err := c.BindJSON(&p); err != nil {
		c.String(500, "error binding product")
	}

	if err := database.UpdateProduct(id, p.Title, p.Price, p.Description); err != nil {
		c.String(400, "error updating product: %c", err)
	} else {
		c.JSON(200, gin.H{
			"product": "updated",
		})
	}
}

func HandleDeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := database.DeleteProduct(id); err != nil {
		c.String(500, "error deleting product %c", err)
	} else {
		c.String(204, "deleted")
	}
}
