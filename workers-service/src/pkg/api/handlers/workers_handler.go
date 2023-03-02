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

func HandleUpdateWorker(c *gin.Context) {
	var worker models.BindWorker
	id := c.Param("id")

	if err := c.BindJSON(&worker); err != nil {
		c.String(500, "error binding worker: %c", err)
	}

	if err := database.UpdateWorker(id, worker.Name, worker.Email, worker.Number); err != nil {
		c.String(400, "error updating worker: %c", err)
	} else {
		c.JSON(200, gin.H{
			"user": "updated",
		})
	}
}

func HandleDeleteWorker(c *gin.Context) {
	id := c.Param("id")

	if err := database.DeleteUser(id); err != nil {
		c.String(500, "error deleting worker: %c", err)
	} else {
		c.String(204, "deleted successfully")
	}
}
