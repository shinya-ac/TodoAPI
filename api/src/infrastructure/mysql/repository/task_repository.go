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

func (r *taskRepository) FindById(ctx context.Context, id string) (*task.Task, error) {
	query := "SELECT id, title, content FROM tasks WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	var task task.Task
	err := row.Scan(&task.Id, &task.Title, &task.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Save(ctx context.Context, task *task.Task) error {
	query := `
		INSERT INTO tasks (id, title, content)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			content = VALUES(content),
			updated_at = VALUES(updated_at)`
	_, err := r.db.ExecContext(ctx, query, task.Id, task.Title, task.Content)
	return err
}
