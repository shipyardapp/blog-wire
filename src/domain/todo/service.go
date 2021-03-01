package todo

import (
	"github.com/google/uuid"
	"github.com/shipyardapp/blog-wire/src/domain/user"
)

type Service interface {
	GetRemoveRepo

	Add(createdBy user.User, text string) (*Todo, error)

	Set(id uuid.UUID, text string) (*Todo, error)
}
