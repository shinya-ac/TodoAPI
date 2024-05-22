package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	Get(ctx context.Context, offset int, pageSize int) ([]*Task, error)
}
