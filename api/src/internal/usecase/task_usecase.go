package usecase

import (
	"context"

	"github.com/shinya-ac/TodoAPI/internal/adapter/repository"
	"github.com/shinya-ac/TodoAPI/internal/domain"
)

type TaskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (tuc *TaskUsecase) CreateTask(ctx context.Context, task domain.Task) (int64, error) {
	return tuc.repo.Create(ctx, task)
}
