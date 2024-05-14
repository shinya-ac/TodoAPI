package mysql

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"github.com/shinya-ac/TodoAPI/internal/adapter/repository"
	"github.com/shinya-ac/TodoAPI/internal/domain"
)

type TaskRepositoryService struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &TaskRepositoryService{db: db}
}

func (repo *TaskRepositoryService) Create(ctx context.Context, task domain.Task) (int64, error) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	result, err := repo.db.ExecContext(ctx, "INSERT INTO tasks (title, content) VALUES(?, ?)", task.Title, task.Content)
	if err != nil {
		logger.Error("インサートエラー：", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("前回インサートID取得エラー：", err)
		return 0, err
	}
	return id, nil
}
