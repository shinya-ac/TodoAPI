package task

import (
	"context"

	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type DeleteTaskUseCase struct {
	taskRepository taskDomain.TaskRepository
}

func NewDeleteTaskUseCase(
	taskRepository taskDomain.TaskRepository,
) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskRepository: taskRepository,
	}
}

type DeleteTaskUseCaseInputDto struct {
	Id string
}

type DeleteTaskUseCaseOutputDto struct {
	Id string
}

func (uc *DeleteTaskUseCase) Run(
	ctx context.Context,
	dto DeleteTaskUseCaseInputDto,
) (*DeleteTaskUseCaseOutputDto, error) {
	t, err := uc.taskRepository.FindById(ctx, dto.Id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, errDomain.NewError("IDに対応するTodoが見つかりません。")
	}

	err = uc.taskRepository.Delete(ctx, dto.Id)
	if err != nil {
		logging.Logger.Error("サーバーエラー", "error", err)
		return nil, err
	}

	return &DeleteTaskUseCaseOutputDto{
		Id: t.Id,
	}, nil
}
