package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	Get(ctx context.Context, offset int, pageSize int, status, searchWord *string) ([]*Task, error)
	Save(ctx context.Context, task *Task) error
	FindById(ctx context.Context, id string) (*Task, error)
	Delete(ctx context.Context, id string) error
}
