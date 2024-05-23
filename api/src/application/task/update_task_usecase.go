package task

import (
	"context"

	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type UpdateTaskUseCase struct {
	taskRepository taskDomain.TaskRepository
}

func NewUpdateTaskUseCase(
	taskRepository taskDomain.TaskRepository,
) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskRepository: taskRepository,
	}
}

type UpdateTaskUseCaseInputDto struct {
	Id      string
	Title   string
	Content string
}

type UpdateTaskUseCaseOutputDto struct {
	Id string
}

func (uc *UpdateTaskUseCase) Run(
	ctx context.Context,
	dto UpdateTaskUseCaseInputDto,
) (*UpdateTaskUseCaseOutputDto, error) {
	t, err := uc.taskRepository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, errDomain.NewError("IDに対応するTodoが見つかりません。")
	}

	if err := t.UpdateTask(dto.Title, dto.Content); err != nil {
		return nil, err
	}

	err = uc.taskRepository.Save(ctx, t)
	if err != nil {
		logging.Logger.Error("サーバーエラー", "error", err)
		return nil, err
	}

	return &UpdateTaskUseCaseOutputDto{
		Id: t.Id,
	}, nil
}
