package task

import (
	"context"

	"github.com/google/uuid"

	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type CreateTaskUseCase struct {
	taskRepo taskDomain.TaskRepository
}

func NewCreateTaskUseCase(
	taskRepo taskDomain.TaskRepository,
) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		taskRepo: taskRepo,
	}
}

type CreateTaskUseCaseInputDto struct {
	Title   string
	Content string
}

type CreateTaskUseCaseOutputDto struct {
	Id uuid.UUID
}

func (uc *CreateTaskUseCase) Run(
	ctx context.Context,
	dto CreateTaskUseCaseInputDto,
) (*CreateTaskUseCaseOutputDto, error) {
	t, err := taskDomain.NewTask(dto.Title, dto.Content)
	if err != nil {
		logging.Logger.Error("サーバーエラー", "error", err)
		return nil, err
	}
	err = uc.taskRepo.Create(ctx, t)
	if err != nil {
		logging.Logger.Error("サーバーエラー", "error", err)
		return nil, err
	}
	return &CreateTaskUseCaseOutputDto{
		Id: t.GetId(),
	}, nil
}
