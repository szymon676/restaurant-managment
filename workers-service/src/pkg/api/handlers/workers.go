package api

import (
	"github.com/gin-gonic/gin"
	"github.com/szymon676/workers-service/src/pkg/database"
	"github.com/szymon676/workers-service/src/pkg/models"
)

func HandleGetWorkers(c *gin.Context) {
	workers, err := database.GetWorkers()
	if err != nil {
		c.String(500, "smth went wrong: %c", err)
	}

	c.JSON(200, gin.H{
		"workers": workers,
	})
}

func HandleCreateWorker(c *gin.Context) {
	var user models.BindWorker
	c.BindJSON(&user)

	database.SaveWorker(user.Name, user.Email, user.Number)

	c.JSON(202, gin.H{
		"status": "created",
	})
}
