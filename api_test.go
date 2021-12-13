package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloMessageWhenServerIsOn(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/", getHelloMessage)

	request, errors := http.NewRequest(http.MethodGet, "/", nil)

	if errors != nil {
		t.Fatalf("Not able to create request : %v\n", errors)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetAllAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/albums", getAlbums)

	request, errors := http.NewRequest(http.MethodGet, "/albums", nil)

	if errors != nil {
		t.Fatalf("Not able to create request : %v\n", errors)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
