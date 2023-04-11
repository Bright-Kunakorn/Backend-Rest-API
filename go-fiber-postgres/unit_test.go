package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func getSKUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}
func getSKUsID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUsID called"})
}

func getSKUsBranch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUs called"})
}
func getSKUsBranchID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getSKUsID called"})
}

func TestGetSKUsBranch(t *testing.T) {
	req, err := http.NewRequest("GET", "/skus_branch", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/skus_branch", getSKUsBranch)
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}
func TestGetSKUsBranchID(t *testing.T) {
	req, err := http.NewRequest("GET", "/skus_branch/017b1c0a-df2e-4f24-be7c-4ae9cadd24ef", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/skus_branch/017b1c0a-df2e-4f24-be7c-4ae9cadd24ef", getSKUsBranch)
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetSKUs(t *testing.T) {
	req, err := http.NewRequest("GET", "/skus", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/skus", getSKUs)
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetSKUsID(t *testing.T) {
	req, err := http.NewRequest("GET", "/skus/017b1c0a-df2e-4f24-be7c-4ae9cadd24ef", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/skus/017b1c0a-df2e-4f24-be7c-4ae9cadd24ef", getSKUsID)
	r.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}
