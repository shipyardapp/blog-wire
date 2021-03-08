package userservice

import (
	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/config"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

type Config struct {
	AllowedEmailDomains []string
}

func NewConfig(enver config.Enver) Config {
	c := Config{}
	c.AllowedEmailDomains = config.StringSlice(enver, "BLOGWIRE_ALLOWED_EMAIL_DOMAINS")
	return c
}

// Service is the todo.Service implementation.
// We have this type in a separate package so that we can have the interface
// and this type without any name games.
type Service struct {
	config Config

	users user.Repo
}

// New creates and returns a new Service.
func New(config Config, users user.Repo) *Service {
	return &Service{
		config: config,
		users:  users,
	}
}

func (s *Service) Get(id uuid.UUID) (*user.User, error) {
	return s.users.Get(id)
}

func (s *Service) Add(email string) (*user.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// If this were real, check email for valid domain.

	u := &user.User{
		ID:    id,
		Email: email,
	}
	return u, nil
}
