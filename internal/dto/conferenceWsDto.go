package dto

import (
	"io"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ConferenceWsDto struct {
	People map[*websocket.Conn]bool
	log    *zap.Logger
}

func NewConferenceWsDto(log *zap.Logger) *ConferenceWsDto {
	return &ConferenceWsDto{
		People: make(map[*websocket.Conn]bool),
		log:    log,
	}
}

func (c *ConferenceWsDto) AddConnection(ws *websocket.Conn) {
	c.People[ws] = true // TODO: Make this concurrent safe
}

func (c *ConferenceWsDto) HandleWs(ws *websocket.Conn) {
	c.readLoop(ws)
}

func (c *ConferenceWsDto) readLoop(ws *websocket.Conn) {
	var message Message

	for {
		err := ws.ReadJSON(&message)
		if nil != err {
			if io.EOF == err {
				break
			}
			c.log.Error("Error in reading bytes", zap.Error(err))
			continue
		}

		c.broadcast(message)
	}
}

func (c *ConferenceWsDto) broadcast(message Message) {
	i := 0
	c.log.Info("Broadcasting", zap.Int("Number", len(c.People)))
	for ws := range c.People {
		c.log.Info("broadcast", zap.Int("person", i))
		i += 1
		go func(ws *websocket.Conn) {
			ws.WriteJSON(message)
		}(ws)
	}
}
