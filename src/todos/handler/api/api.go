package api

import (
	"log"
	"net/http"

	"github.com/shipyardapp/blog-wire/src/config"
	"github.com/shipyardapp/blog-wire/src/domain/apm"
	"github.com/shipyardapp/blog-wire/src/domain/todo"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

type Config struct {
	DisplayURL string
	SigningKey []byte
}

func NewConfig(enver config.Enver) (Config, error) {
	c := Config{}
	var err error

	if c.DisplayURL, err = config.RequiredString(enver, "BLOGWIRE_DISPLAY_URL"); err != nil {
		return Config{}, err
	}
	if c.SigningKey, err = config.RequiredBase64Bytes(enver, "BLOGWIRE_API_SIGNING_KEY"); err != nil {
		return Config{}, err
	}

	return c, nil
}

type API struct {
	config Config

	apm apm.APM

	logger *log.Logger

	users user.Service

	todos todo.Service
}

func New(config Config, apm apm.APM, logger *log.Logger, users user.Service, todos todo.Service) *API {
	return &API{
		config: config,
		apm:    apm,
		logger: logger,
		users:  users,
		todos:  todos,
	}
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOTE
	// Not really going to do anything here.
	// Just imagine that in some sort of way, through middleware and other
	// constructs, that all of a's fields would be used.
}
