package sqldb

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/shipyardapp/blog-wire/src/config"
)

type Config struct {
	DataSourceName string
}

func NewConfig(enver config.Enver) (Config, error) {
	// TODO

	return Config{}, nil
}

func New(config Config) (*sql.DB, error) {
	return sql.Open("postgres", config.DataSourceName)
}
