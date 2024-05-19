package repository

import (
	"context"
	"database/sql"

	"github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) task.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *task.Task) error {
	query := `INSERT INTO tasks (title, content) VALUES(?, ?)`

	_, err := r.db.ExecContext(ctx, query, task.Title, task.Content)
	if err != nil {
		logging.Logger.Error("SQL実行に失敗", "error", err)
		return err
	}
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	return nil
}
