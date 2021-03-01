package todoservice

import (
	"time"

	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/config"
	"github.com/shipyardapp/blog-wire/src/domain/todo"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

var _ todo.Service = &Service{}

type Config struct {
	BadWords []string
}

func NewConfig(enver config.Enver) (Config, error) {
	return Config{}, nil
}

type Service struct {
	config Config

	todos todo.Repo
}

func New(config Config, todos todo.Repo) *Service {
	return &Service{
		config: config,
		todos:  todos,
	}
}

func (s *Service) Get(id uuid.UUID) (*todo.Todo, error) {
	return s.todos.Get(id)
}

func (s *Service) Add(createdBy user.User, text string) (*todo.Todo, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// TODO check for text contains bad words

	todo := &todo.Todo{
		ID:        id,
		Owner:     createdBy.ID,
		Text:      text,
		UpdatedAt: time.Now(),
	}
	return todo, nil
}

func (s *Service) Set(id uuid.UUID, text string) (*todo.Todo, error) {
	todo, err := s.todos.Get(id)
	if err != nil {
		return nil, err
	}

	// TODO check for text contains bad words

	todo.Text = text
	todo.UpdatedAt = time.Now()

	if err := s.todos.Set(*todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Service) Remove(id uuid.UUID) error {
	return s.todos.Remove(id)
}
