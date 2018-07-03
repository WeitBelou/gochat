package messages

import (
	"sort"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type InMemStorage struct {
	mu sync.Mutex

	clients map[string]*websocket.Conn

	messages []Message

	limit  int
	curPos int
}

func (s *InMemStorage) AddWSClient(login string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.clients[login]; exists {
		s.clients[login].Close()
	}
	s.clients[login] = conn
}

func (s *InMemStorage) Post(author string, text string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	message := Message{
		Author:    author,
		Text:      text,
		CreatedAt: time.Now(),
	}

	if len(s.messages) < s.limit {
		s.messages = append(s.messages, message)
		s.curPos++
	} else if s.curPos < s.limit {
		s.messages[s.curPos] = message
		s.curPos++
	} else {
		s.curPos -= s.limit
		s.messages[s.curPos] = message
	}

	for login, c := range s.clients {
		err := c.WriteJSON(message)
		if err != nil {
			return errors.Wrapf(err, "failed to write message to client: %s", login)
		}
	}

	return nil
}

func (s *InMemStorage) List() []Message {
	messages := make([]Message, len(s.messages))
	copy(messages, s.messages)

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.Unix() < messages[j].CreatedAt.Unix()
	})

	return messages
}

func New(cfg Config) *InMemStorage {
	return &InMemStorage{
		clients:  make(map[string]*websocket.Conn),
		messages: make([]Message, 0, cfg.Limit),
		limit:    cfg.Limit,
	}
}
