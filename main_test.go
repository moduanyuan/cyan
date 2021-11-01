package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var TestRouter *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	TestRouter = gin.Default()
}

func TestHandler(t *testing.T) {
	TestRouter.GET("/", Handler)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	TestRouter.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Error("Request Error!")
	}
}
