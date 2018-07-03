package messages

import (
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Service interface {
	Post(author string, text string) error
	List() []Message

	AddWSClient(login string, conn *websocket.Conn)
}
