package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"treeleaf/controllers"
	"treeleaf/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
var obtainedCode=""


func TestCreateShortURL(t *testing.T) {
	r := SetUpRouter()
	r.POST("/shorten", controllers.CreateShortURL)

	url:="https://google.com"

	validInput := models.InputSchema{URL: url}

	validInputJSON, _ := json.Marshal(validInput)
	
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(validInputJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]string

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error parsing JSON response: %v", err)
	}
	shortURL, ok := response["shortURL"]
	if ok{
		shortCode := extractShortCode(shortURL)
		obtainedCode=shortCode
		// fmt.Println(shortCode)
	}

}

// func TestGetURL(t *testing.T) {
// 	r := SetUpRouter()

// 	shortCode := obtainedCode
// 	longURL := "https://google.com"
// 	models.URLs[shortCode] = longURL

// 	req := httptest.NewRequest("GET", "/"+shortCode, nil)

// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	if w.Code != http.StatusMovedPermanently {
// 		t.Errorf("Expected status %d, got %d", http.StatusMovedPermanently, w.Code)
// 	}

// 	expectedLocation := longURL
// 	if gotLocation := w.Header().Get("Location"); gotLocation != expectedLocation {
// 		t.Errorf("Expected redirect location %s, got %s", expectedLocation, gotLocation)
// 	}
// }

func extractShortCode(shortURL string) string {
    parts := strings.Split(shortURL, "/")
    return parts[len(parts)-1]
}