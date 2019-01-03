package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tr *Trigger

func doEvent(c *gin.Context) {
	var ev Event
	if c.Bind(&ev) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}
	log.Printf("get event:%+v", ev)
	err := tr.AddEvent(&ev)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "add event error"})
		return
	}

	// dump events
	tr.DumpEventList()

	// check top events
	evs, err := tr.GetEventsFront()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "get events front error"})
		return
	}
	log.Printf("evs %+v", evs)

	// do something with events
	// do your staff

	c.JSON(http.StatusCreated, gin.H{"status": "do event ok"})
}

func getEvents(c *gin.Context) {
	tr.DumpEventList()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Post event
	r.POST("/trigger/events", doEvent)

	// get events
	r.GET("/trigger/events", getEvents)

	// delete events
	//r.GET("/trigger/events", getEvents)

	return r
}

func main() {
	tr = NewTrigger()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
