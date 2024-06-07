package main

import (
	"log"
	"time"
	"ws-client-master/handler"
	"ws-client-master/registry"

	"github.com/randyardiansyah25/wsbase-handler"
)

func main() {

	WSClient := wsbase.NewWSClient("localhost:6680", "/connect/0001", false)
	WSClient.SetReconnectPeriod(5 * time.Second)

	handler := handler.NewWsHandler(WSClient)
	registry.RegisterHandler(handler)
	WSClient.SetMessageHandler(handler.GetHandler())
	er := WSClient.Start()
	if er != nil {
		log.Println("cannot start ws client, error :", er)
	}
}
