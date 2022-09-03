package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"snappfood_challenge/forms"
	"snappfood_challenge/internal"
)

type OrderController struct {}

var orderService = new(internal.OrderService)

// the only endpoint that gets the input order and
// put it on a redis channel
func (o OrderController)AddOrder(c *gin.Context)  {
	var req forms.AddOrder

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to handle data",
			"error": err})
		c.Abort()
		return
	}

	if err := orderService.PublishOrder(req); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "Error to publish data",
			"error": err})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data published!"})
	c.Abort()
	return
}