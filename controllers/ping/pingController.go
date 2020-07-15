package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(*gin.Context)
}

type pingController struct {
}

// Ping func for health check
func (ct *pingController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
