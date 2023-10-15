package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/p2064/gateway/client"
)

func main() {
	router := gin.Default()
	svc := &client.Clients{
		Creator: client.InitCreatorClient(),
		Changer: client.InitChangerClient(),
		Adder:   client.InitAdderClient(),
	}

	apiGrop := router.Group("/api")

	apiGrop.POST("/create", svc.CreatEvent)
	apiGrop.POST("/change", svc.ChangeEvent)
	apiGrop.POST("/add", svc.AddToEvent)
	log.Print("Start gateway")
	router.Run(":8081")

}
