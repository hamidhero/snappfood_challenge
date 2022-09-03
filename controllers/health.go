package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct{}

// Status is the endpoint for checking if the app is running
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
