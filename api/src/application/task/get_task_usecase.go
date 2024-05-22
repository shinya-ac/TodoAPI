package task

import (
	"context"

	"github.com/shinya-ac/TodoAPI/domain/task"
	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type GetTaskUseCase struct {
	taskRepository taskDomain.TaskRepository
}

func NewGetTaskUseCase(
	taskRepository taskDomain.TaskRepository,
) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskRepository: taskRepository,
	}
}

type GetTaskUseCaseOutputDto struct {
	Tasks []*task.Task
}

func (uc *GetTaskUseCase) Run(
	ctx context.Context,
	offset int,
	pageSize int,
) (*GetTaskUseCaseOutputDto, error) {
	tasks, err := uc.taskRepository.Get(ctx, offset, pageSize)
	if err != nil {
		logging.Logger.Error("サーバーエラー", "error", err)
		return nil, err
	}

	return &GetTaskUseCaseOutputDto{tasks}, nil
}
