package websocket

import (
	"encoding/json"
	"log"
	"time"
)

type WsMessage struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Channel   string    `json:"channel"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (message *WsMessage) encode() []byte {
	json, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
	}

	return json
}
