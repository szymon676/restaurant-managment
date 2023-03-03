package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	api "github.com/szymon676/products-service/internal/api/handlers"
	"github.com/szymon676/products-service/internal/database"
	"github.com/szymon676/products-service/internal/models"
)

func init() {
	database.ConnectDB()
}

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	title := "Test Product"
	price := 9.99
	description := "This is a test product"

	p := models.BindProduct{Title: title, Price: price, Description: description}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("error marshaling product: %v", err)
	}

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	api.HandleCreateProduct(c)

	if rec.Code != 202 {
		t.Fatalf("expected status code 202; got %d", rec.Code)
	}
}

func TestGetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	database.SaveProduct("Product 1", 10.0, "Test product 1")
	database.SaveProduct("Product 2", 20.0, "Test product 2")
	database.SaveProduct("Product 3", 30.0, "Test product 3")

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}

	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	api.HandleGetProducts(c)

	if rec.Code != 200 {
		t.Fatalf("expected status code 200; got %d", rec.Code)
	}

	var response map[string][]models.Product

	err = json.Unmarshal(rec.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("error unmarshaling response: %v", err)
	}

	products := response["products"]
	if len(products) != 5 {
		t.Fatalf("expected 5 products; got %d", len(products))
	}
}
