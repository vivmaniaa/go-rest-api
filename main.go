package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	mDB "practice.com/rest-api/mongoDb"
)

func main() {
	mDB.ConnectMongoDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8081") // localhost:8081
}

func createEvent(context *gin.Context) {
	var event mDB.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to create the request."})
		return
	}
	event.UserId = 1
	err = mDB.InsertEvent(event)

	if err != nil {
		fmt.Println("Unable to insert event to the database!")
		return
	}

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, mDB.ListAll())
}
