package task

import "context"

type GetTasksDto struct {
	ID      string
	Title   string
	Content string
}

type TaskQueryService interface {
	GetTasks(ctx context.Context) ([]*GetTasksDto, error)
}
