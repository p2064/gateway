package client

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	addProto "github.com/p2064/adder/proto"
	chProto "github.com/p2064/changer/proto"
	crProto "github.com/p2064/creator/proto"
	"github.com/p2064/gateway/handlers"
	"github.com/p2064/pkg/config"
	"google.golang.org/grpc"
)

type Clients struct {
	Creator crProto.CreatorServiceClient
	Changer chProto.ChangerServiceClient
	Adder   addProto.AdderServiceClient
}

func init() {
	if config.Status != config.GOOD {
		log.Fatal("failed to get config")
	}
}

func InitCreatorClient() crProto.CreatorServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(fmt.Sprintf("creator:9000"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return crProto.NewCreatorServiceClient(cc)
}

func InitChangerClient() chProto.ChangerServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial("changer:9001", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return chProto.NewChangerServiceClient(cc)
}

func InitAdderClient() addProto.AdderServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(fmt.Sprintf("adder:9002"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return addProto.NewAdderServiceClient(cc)
}

func (svc *Clients) CreatEvent(ctx *gin.Context) {
	handlers.CreatEvent(ctx, svc.Creator)
}

func (svc *Clients) ChangeEvent(ctx *gin.Context) {
	handlers.ChangeEvent(ctx, svc.Changer)
}

func (svc *Clients) AddToEvent(ctx *gin.Context) {
	handlers.AddToEvent(ctx, svc.Adder)
}
