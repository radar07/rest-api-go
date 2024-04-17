package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/radar07/rest-api-go/models"
)

func getEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(http.StatusOK, events)
}

func postEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request.",
		})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}

func main() {
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello, world!",
		})
	})

	server.GET("/events", getEvents)
	server.POST("/events", postEvent)

	server.Run(":8080")
}
