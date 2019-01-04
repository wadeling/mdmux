package main

import (
	"fmt"
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
	// do your stuff

	c.JSON(http.StatusCreated, gin.H{"status": "do event ok"})
}

func delEvents(c *gin.Context) {
	evList, err := tr.PopTopEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "del events error"})
		return
	}

	// dump events
	tr.DumpEventList()

	log.Printf("evs %+v", evList)

	c.JSON(http.StatusCreated, gin.H{"status": " pop event ok"})
}

func getEvents(c *gin.Context) {
	//tr.DumpEventList()
	result := make(map[string]interface{})
	for i, v := range tr.EventList {
		for h, k := range v {
			if k != nil {
				tmp := make(map[string]interface{})
				tmp["src"] = k.Src
				tmp["uuid"] = k.UUid
				tmp["ip"] = k.Ip
				key := fmt.Sprintf("%d_%d", i, h)
				result[key] = tmp
			}
		}
	}

	c.JSON(http.StatusOK, result)

	log.Printf("get events ok")
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
	r.DELETE("/trigger/events", delEvents)

	return r
}

func main() {
	tr = NewTrigger()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
