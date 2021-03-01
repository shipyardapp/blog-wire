package todo

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	// The ID of the Todo.
	ID uuid.UUID

	// The ID of the User who owns the Todo.
	Owner uuid.UUID

	// The content of the Todo.
	Text string

	// When the Todo was most recently created or updated.
	UpdatedAt time.Time
}
