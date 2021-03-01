package user

import "github.com/google/uuid"

type GetRepo interface {
	Get(id uuid.UUID) (*User, error)
}

type Repo interface {
	GetRepo

	Add(user User) error
}
