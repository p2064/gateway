package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	addProto "github.com/p2064/adder/proto"
	chProto "github.com/p2064/changer/proto"
	crProto "github.com/p2064/creator/proto"
)

type Event struct {
	Date       time.Time `form:"date" json:"date" binding:"required"`
	MaxPlayers int       `form:"max_players" json:"max_players" binding:"required"`
	Place      string    `form:"place" json:"place" binding:"required"`
}

func CreatEvent(ctx *gin.Context, c crProto.CreatorServiceClient) {
	var requestBody crProto.CreateEventRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Error when bindind JSON: %s", err)
	}
	log.Printf("Got data JSON: %+v", requestBody)

	resp, err := c.CreateEvent(context.Background(), &requestBody)
	if err != nil {
		log.Printf("Error when calling CreatEvent: %s", err)
	}

	log.Printf("Response from server: %+v", resp)
	ctx.JSON(200, gin.H{
		"response": resp,
	})

}

func ChangeEvent(ctx *gin.Context, c chProto.ChangerServiceClient) {
	var requestBody chProto.ChangeEventRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Error when bindind JSON: %s", err)
	}
	log.Printf("Got data JSON: %+v", requestBody)

	resp, err := c.ChangeEvent(context.Background(), &requestBody)
	if err != nil {
		log.Printf("Error when calling ChangeEvent: %s", err)
	}

	log.Printf("Response from server: %+v", resp)
	ctx.JSON(200, gin.H{
		"response": resp,
	})

}

func AddToEvent(ctx *gin.Context, c addProto.AdderServiceClient) {
	var requestBody addProto.AddToEventRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Printf("Error when bindind JSON: %s", err)
	}
	log.Printf("Got data JSON: %+v", requestBody)

	resp, err := c.AddToEvent(context.Background(), &requestBody)
	if err != nil {
		log.Printf("Error when calling AddToEvent: %s", err)
	}

	log.Printf("Response from server: %+v", resp)
	ctx.JSON(200, gin.H{
		"response": resp,
	})

}
