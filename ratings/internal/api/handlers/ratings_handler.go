package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/szymon676/ratings-service/internal/database"
	"github.com/szymon676/ratings-service/internal/models"
)

func HandleGetRatings(c *gin.Context) {
	if ratings, err := database.GetRatings(); err != nil {
		c.String(500, "error getting ratings %c", err)
	} else {
		c.JSON(200, gin.H{
			"ratings": ratings,
		})
	}
}

func HandleSaveRatings(c *gin.Context) {
	var rating *models.BindRating

	if err := c.BindJSON(&rating); err != nil {
		c.String(500, "Error binding ratings: %s", err.Error())
		return
	}

	if rating.User == "" || rating.Comment == "" || rating.Stars == 0 {
		c.String(400, "Please fill all fields")
		return
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, rating.Date)

	if err != nil {
		c.String(500, "Error parsing date: %s", err.Error())
		return
	}

	if err := database.SaveRating(rating.Stars, rating.User, rating.Comment, date); err != nil {
		c.String(500, "Error saving rating: %s", err.Error())
	} else {
		c.JSON(202, gin.H{
			"status": "success",
		})
	}
}

func HandleDeleteRating(c *gin.Context) {
	id := c.Param("id")
	database.DeleteRating(id)
	c.String(204, "ok")
}
