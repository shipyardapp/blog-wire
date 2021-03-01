package userservice

import (
	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/config"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

var _ user.Service = &Service{}

type Config struct {
	AllowedEmailDomains []string
}

func NewConfig(enver config.Enver) (Config, error) {
	// TODO
	return Config{}, nil
}

type Service struct {
	config Config

	users user.Repo
}

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

	// TODO check email for valid domain.

	u := &user.User{
		ID:    id,
		Email: email,
	}
	return u, nil
}
