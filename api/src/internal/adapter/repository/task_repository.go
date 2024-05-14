package repository

import (
	"context"

	"github.com/shinya-ac/TodoAPI/internal/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task domain.Task) (int64, error)
}
