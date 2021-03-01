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
	c := Config{}
	var err error

	if c.DataSourceName, err = config.RequiredString(enver, "BLOGWIRE_DB_DATA_SOURCE_NAME"); err != nil {
		return Config{}, err
	}

	return c, nil
}

func New(config Config) (*sql.DB, error) {
	return sql.Open("postgres", config.DataSourceName)
}
