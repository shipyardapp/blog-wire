package todoservice

import (
	"time"

	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/config"
	"github.com/shipyardapp/blog-wire/src/domain/todo"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

type Config struct {
	BadWords []string
}

func NewConfig(enver config.Enver) Config {
	c := Config{}
	c.BadWords = config.StringSlice(enver, "BLOGWIRE_TODO_BAD_WORDS")
	return c
}

// Service is the todo.Service implementation.
// We have this type in a separate package so that we can have the interface
// and this type without any name games.
type Service struct {
	config Config

	todos todo.Repo
}

// New creates and returns a new Service.
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

	// If this were real, check for text contains bad words and return an error.

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

	// If this were real, check for text contains bad words and return an error.

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
