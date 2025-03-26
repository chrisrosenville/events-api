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

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
        return
    }

	_, err = models.GetEventByID(eventId)
	if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
        return
    }

	var updatedEvent models.Event;

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
        return
    }

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"message": "Event successfully updated", "event": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
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

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again later."})
        return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})
}