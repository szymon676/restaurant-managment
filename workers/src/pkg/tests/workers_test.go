package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	api "github.com/szymon676/workers-service/src/pkg/api/handlers"
	"github.com/szymon676/workers-service/src/pkg/database"
	"github.com/szymon676/workers-service/src/pkg/models"
)

func init() {
	database.ConnectDB()
}

func TestCreateWorker(t *testing.T) {
	gin.SetMode(gin.TestMode)

	name := "Ania"
	email := "ania@gmail.com"
	number := "123-321-908"

	w := models.BindWorker{Name: name, Email: email, Number: number}

	bw, err := json.Marshal(w)
	if err != nil {
		t.Fatalf("Unable to marshal %c", err)
	}

	req, err := http.NewRequest("POST", "/workers", bytes.NewBuffer(bw))
	if err != nil {
		t.Fatalf("Unable to create worker: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	api.HandleCreateWorker(c)
	if rec.Code != 202 {
		t.Fatalf("wrong status expected: %d got: %d", 202, rec.Code)
	}
}

func TestGetWorkers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	database.SaveWorker("foo", "bar", "baz")

	req, err := http.NewRequest("GET", "/workers", nil)
	if err != nil {
		t.Fatalf("unable to get workers: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	api.HandleGetWorkers(c)
	if rec.Code != 200 {
		t.Fatalf("wrong status expected: %d got: %d", 200, rec.Code)
	}

	var response map[string][]models.Worker

	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("error unmarshaling response: %v", err)
	}

	workers := response["workers"]
	if len(workers) != 4 {
		t.Fatalf("expected 4 workers; got %d", len(workers))
	}
}
