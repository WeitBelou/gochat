package messages

type InMemStorage struct {
	messages []Message
}

func (s *InMemStorage) Post(author string, text string) error {
	panic("implement me")
}

func (s *InMemStorage) List() ([]Message, error) {
	panic("implement me")
}

func New(cfg Config) *InMemStorage {
	return &InMemStorage{
		messages: make([]Message, cfg.Limit),
	}
}
