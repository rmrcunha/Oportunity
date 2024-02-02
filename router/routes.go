package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rmrcunha/Oportunity.git/handlers"
)

func initializeRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handlers.ShowOpeningHandler)
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.DELETE("/opening", handlers.DeleteOpeningHandler)
		v1.PUT("/opening", handlers.UpdatepeningHandler)
		v1.GET("/openings", handlers.ListOpeningsHandler)
	}
}
