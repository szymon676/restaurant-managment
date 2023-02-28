package api

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/workers-service/src/pkg/database"
	"github.com/szymon676/workers-service/src/pkg/models"
)

func HandleGetWorkers(c *gin.Context) {
	users, err := database.GetWorkers()
	if err != nil {
		c.String(500, "smth went wrong: %c", err)
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func HandleCreateWorker(c *gin.Context) {
	var user models.BindWorker
	c.BindJSON(&user)

	database.SaveUser(user.Name, user.Email, user.Number)

	c.JSON(200, gin.H{
		"status": "success",
	})
}
