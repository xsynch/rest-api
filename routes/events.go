package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"udemy.com/rest-api/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
		return
	}
	context.JSON(http.StatusOK, events)
	
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect ID specified"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event."})
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		fmt.Println(err)
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created successfully", "event": event})
}

func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return 
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not get the event id."})
		return 
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})		
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update the event id."})
		return 

	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})

}

func deleteEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return 
	}

	event ,err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not get the event id."})
		return 
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event."})
		return 
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}