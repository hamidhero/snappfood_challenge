package server

import (
	"github.com/gin-gonic/gin"
	"snappfood_challenge/controllers"
)

// NewRouter declares route addresses
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			orderGroup := v1.Group("order")
			{
				order := new(controllers.OrderController)
				orderGroup.POST("", order.AddOrder)
			}
		}
	}
	return router
}
