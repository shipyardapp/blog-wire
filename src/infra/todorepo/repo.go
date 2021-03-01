package todorepo

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/domain/todo"
)

var _ todo.Repo = &Repo{}

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Get(id uuid.UUID) (*todo.Todo, error) {
	return nil, errors.New("uimplemented")
}

func (r *Repo) Add(todo todo.Todo) error {
	return errors.New("unimplemented")
}

func (r *Repo) Set(todo todo.Todo) error {
	return errors.New("unimplemented")
}

func (r *Repo) Remove(id uuid.UUID) error {
	return errors.New("unimplemented")
}
