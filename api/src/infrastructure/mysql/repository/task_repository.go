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
	logging.Logger.Info("Create実行", "task:", task)
	query := `INSERT INTO tasks (id, title, content) VALUES(?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, task.Id, task.Title, task.Content)
	if err != nil {
		logging.Logger.Error("SQL実行に失敗", "error", err)
		return err
	}
	return nil
}

func (r *taskRepository) Get(ctx context.Context, offset int, pageSize int) ([]*task.Task, error) {
	logging.Logger.Info("Get実行", "offset:", offset)

	query := "SELECT id, title, content FROM tasks ORDER BY created_at DESC LIMIT ? OFFSET ?"

	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		logging.Logger.Error("SQL実行に失敗", "error", err)
		return nil, err
	}
	var tasks []*task.Task
	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Content)
		if err != nil {
			logging.Logger.Error("行のスキャンに失敗", "error", err)
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}
