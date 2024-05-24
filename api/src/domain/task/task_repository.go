package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	Get(ctx context.Context, offset int, pageSize int, status *string) ([]*Task, error)
	Save(ctx context.Context, task *Task) error
	FindById(ctx context.Context, id string) (*Task, error)
}
