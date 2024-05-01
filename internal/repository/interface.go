package repository

import (
	"context"

	"github.com/vlad-marlo/pgx-example/internal/model"
)

type Interface interface {
	Store(ctx context.Context, todo model.TodoCreateRequest) (*model.TodoDTO, error)
	Get(ctx context.Context, id int) (*model.TodoDTO, error)
}
