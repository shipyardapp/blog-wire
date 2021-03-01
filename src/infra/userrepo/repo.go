package userrepo

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

var _ user.Repo = &Repo{}

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Get(id uuid.UUID) (*user.User, error) {
	return nil, errors.New("unimplemented")
}

func (r *Repo) Add(user user.User) error {
	return errors.New("unimplemented")
}
