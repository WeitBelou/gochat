package messages

import (
	"sort"
	"sync"
	"time"
)

type InMemStorage struct {
	mu sync.Mutex

	messages []Message

	limit  int
	curPos int
}

func (s *InMemStorage) Post(author string, text string) {
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
		messages: make([]Message, 0, cfg.Limit),
		limit:    cfg.Limit,
	}
}
