package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.GET("/blog", BlogController.GetBlogs)
}

func TestGetBlogs(t *testing.T) {
	req, err := http.NewRequest("GET", "/blog", nil)
	if err != nil {
		t.Fail()
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)

}
