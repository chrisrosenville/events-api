package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/events", getEvents)
	api.POST("/events", createEvent)
	api.GET("/event/:id", getEvent)
	api.PUT("/event/:id", updateEvent)
	api.DELETE("/event/:id", deleteEvent)

	api.POST("/signup", signup)
	
}