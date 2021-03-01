package todo

import "github.com/google/uuid"

type GetRemoveRepo interface {
	Get(id uuid.UUID) (*Todo, error)

	Remove(id uuid.UUID) error
}

type Repo interface {
	GetRemoveRepo

	Add(Todo) error

	Set(Todo) error
}
