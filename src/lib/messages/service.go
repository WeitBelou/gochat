package messages

import "time"

type Message struct {
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Service interface {
	Post(author string, text string)
	List() []Message
}
