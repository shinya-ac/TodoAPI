package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
}
