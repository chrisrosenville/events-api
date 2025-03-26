package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully found all events", "events": events})
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
        return
    }

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
        return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully found event", "event": event})
}

func createEvent(ctx *gin.Context) {
	var event models.Event;
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
        return
    }

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
        return
    }

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event successfully created", "event": event})
}