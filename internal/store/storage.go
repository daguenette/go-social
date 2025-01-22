package store

import (
	"context"
	"database/sql"

	"github.com/daguenette/go-social/internal/schemas"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *schemas.Post) error
	}
	Users interface {
		Create(context.Context, *schemas.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
