package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/ratings-service/src/pkg/database"
)

func HandleGetRatings(c *gin.Context) {
	if user, err := database.GetRatings(); err != nil {
		c.String(500, "error getting ratings %c", err)
	} else {
		c.JSON(200, gin.H{
			"users": user,
		})
	}
}
