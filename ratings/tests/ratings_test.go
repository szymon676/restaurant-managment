package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/szymon676/ratings-service/internal/api/handlers"
	"github.com/szymon676/ratings-service/internal/database"
	"github.com/szymon676/ratings-service/internal/models"
)

func init() {
	database.ConnectDB()
}

func TestCreateRatings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	user := "Ania"
	stars := 7.5
	comment := "such a nice products"
	date := "2013-12-11"

	r := models.BindRating{User: user, Comment: comment, Stars: stars, Date: date}

	br, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Unable to marshal %c", err)
	}

	req, err := http.NewRequest("POST", "/ratings", bytes.NewBuffer(br))
	if err != nil {
		t.Fatalf("Unable to create rating: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	handlers.HandleSaveRatings(c)
	if rec.Code != 202 {
		t.Fatalf("wrong status expected: %d got: %d", 202, rec.Code)
	}
}

func TestGetRatings(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, err := http.NewRequest("GET", "/ratings", nil)
	if err != nil {
		t.Fatalf("Unable to create rating: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	handlers.HandleGetRatings(c)
	if rec.Code != 200 {
		t.Fatalf("wrong status expected: %d got: %d", 200, rec.Code)
	}

	if rec.Body == nil {
		t.Fatalf("expected few ratings got 0")
	}
}

func TestDeleteRating(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, err := http.NewRequest("DELETE", "/ratings/2", nil)
	if err != nil {
		t.Fatalf("Unable to delete rating: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	handlers.HandleDeleteRating(c)
	if rec.Code != 204 {
		t.Fatalf("wrong status expected: %d got: %d", 204, rec.Code)
	}

	if rec.Body.Len() > 0 {
		t.Fatalf("expected nothing")
	}
}
