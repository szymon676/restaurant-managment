package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/szymon676/orders-service/internal/api/handlers"
	"github.com/szymon676/orders-service/internal/database"
	"github.com/szymon676/orders-service/internal/models"
)

func init() {
	database.ConnectDB()
}

func TestSaveOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	table := 2
	product := "product"

	order := models.BindOrder{
		Table:   table,
		Product: product,
	}

	bo, err := json.Marshal(order)
	if err != nil {
		t.Fatalf("Unable to marshal %c", err)
	}

	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(bo))
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	handlers.HandleCreateOrder(c)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("wrong status code returned: %d wanted %d", rec.Code, http.StatusAccepted)
	}
}

func TestGetOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req, err := http.NewRequest("GET", "/orders", nil)

	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	handlers.HandleGetOrders(c)

	if rec.Code != http.StatusOK {
		t.Fatalf("wrong status code returned: %d wanted %d", rec.Code, http.StatusOK)
	}
}
