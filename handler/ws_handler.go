package handler

import (
	"encoding/json"
	"log"
	"ws-client-master/entities"

	"github.com/randyardiansyah25/wsbase-handler"
)

type HandlerFunc func(client wsbase.WSClient, message wsbase.Message, body string)

type WSHandler interface {
	GetHandler() wsbase.MessageHandler
	AddHandler(action string, handler HandlerFunc)
}

func NewWsHandler(client wsbase.WSClient) WSHandler {
	return &wsHandlerImpl{
		client:   client,
		registry: make(map[string]HandlerFunc),
	}
}

type wsHandlerImpl struct {
	client   wsbase.WSClient
	registry map[string]HandlerFunc
}

func (ws *wsHandlerImpl) GetHandler() wsbase.MessageHandler {
	return func(msg wsbase.Message) {

		byteBody, er := json.Marshal(msg.Body)
		if er != nil {
			log.Println("ERROR :", er.Error())
			return
		}

		handler := ws.registry[msg.Action]
		if handler == nil {
			msg.Body = entities.BaseResponseString{
				BaseResponse: entities.BaseResponse{
					ResponseCode:    "1111",
					ResponseMessage: "Action request tidak dikenal",
				},
				ResponseData: "",
			}
			ws.client.SendMessage(msg)
			return
		}

		handler(ws.client, msg, string(byteBody))
	}
}

func (ws *wsHandlerImpl) AddHandler(action string, handler HandlerFunc) {
	ws.registry[action] = handler
}
