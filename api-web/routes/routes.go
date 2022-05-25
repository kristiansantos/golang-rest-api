package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kristiansantos/go-api-rest/api-web/controllers"
)

func HandleRequests() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/planets", controllers.Index)
		v1.GET("/planets/:id", controllers.Show)
		v1.DELETE("/planets/:id", controllers.Delete)
		v1.PATCH("/planets/:id", controllers.Update)
		v1.POST("/planets", controllers.Create)
	}

	router.Run(":8080")
}
