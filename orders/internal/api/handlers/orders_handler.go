package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/orders-service/internal/database"
	"github.com/szymon676/orders-service/internal/models"
)

func HandleCreateOrder(c *gin.Context) {
	var order models.BindOrder

	if err := c.BindJSON(&order); err != nil {
		c.String(402, "error binding order %c", err)
	}

	if err := database.CreateOrder(order.Product, order.Table); err != nil {
		c.String(500, "error saving order to db")
	}

	c.String(202, "success")
}

func HandleGetOrders(c *gin.Context) {
	if result, err := database.GetOrders(); err != nil {
		c.String(500, "error getting orders")
	} else {
		c.JSON(200, gin.H{
			"orders": result,
		})
	}
}
