package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetSKUsBranch(t *testing.T) {
	// create a new request with a GET method and "/skus_branch" endpoint
	req, err := http.NewRequest("GET", "/skus_branch", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new recorder to record the response
	recorder := httptest.NewRecorder()

	// create a new instance of gin to handle the request
	r := gin.Default()

	// call the handler function with the new request and recorder
	r.GET("/skus_branch", getSKUsBranch)

	r.ServeHTTP(recorder, req)

	// check the status code of the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// check the content type of the response
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}

func TestGetSKUs(t *testing.T) {
	// create a new request with a GET method and "/skus" endpoint
	req, err := http.NewRequest("GET", "/skus", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new recorder to record the response
	recorder := httptest.NewRecorder()

	// create a new instance of gin to handle the request
	r := gin.Default()

	// call the handler function with the new request and recorder
	r.GET("/skus", getSKUs)

	r.ServeHTTP(recorder, req)

	// check the status code of the response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// check the content type of the response
	assert.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
}
