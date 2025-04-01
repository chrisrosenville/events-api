package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/events", getEvents)
	
	authenticated := api.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.GET("/event/:id", getEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)
	authenticated.POST("/event/:id/register", registerForEvent)
	authenticated.DELETE("/event/:id/register", cancelRegistration)

	api.POST("/signup", signup)
	api.POST("/login", login)
	
}