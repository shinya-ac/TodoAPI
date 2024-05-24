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
	query := `INSERT INTO tasks (id, title, content, status) VALUES(?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, task.Id, task.Title, task.Content, task.Status)
	if err != nil {
		logging.Logger.Error("SQL実行に失敗", "error", err)
		return err
	}
	return nil
}

func (r *taskRepository) Get(ctx context.Context, offset int, pageSize int, status *string) ([]*task.Task, error) {
	logging.Logger.Info("Get実行", "offset:", offset)

	query := "SELECT id, title, content, status FROM tasks"
	var args []interface{}

	if status != nil {
		logging.Logger.Info("status:", "", *status)
		query += " WHERE status = ?"
		args = append(args, *status)
	}

	query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		logging.Logger.Error("SQL実行に失敗", "error", err)
		return nil, err
	}
	defer rows.Close()
	var tasks []*task.Task
	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Content, &t.Status)
		if err != nil {
			logging.Logger.Error("行のスキャンに失敗", "error", err)
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func (r *taskRepository) FindById(ctx context.Context, id string) (*task.Task, error) {
	query := "SELECT id, title, content, status FROM tasks WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)
	var task task.Task
	err := row.Scan(&task.Id, &task.Title, &task.Content, &task.Status)
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
		INSERT INTO tasks (id, title, content, status)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			title = VALUES(title),
			content = VALUES(content),
			status = VALUES(status),
			updated_at = NOW()`
	_, err := r.db.ExecContext(ctx, query, task.Id, task.Title, task.Content, task.Status)
	return err
}
