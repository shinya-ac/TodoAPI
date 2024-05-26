package repository

import (
	"context"
	"database/sql"
	"strings"

	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
	"github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) task.TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(ctx context.Context, task *task.Task) error {
	if task == nil {
		logging.Logger.Error("Taskがnil")
		err := errDomain.NewError("Taskがnilです。")
		return err
	}
	logging.Logger.Info("Create実行", "task:", *task)
	query := `INSERT INTO tasks (id, title, content, status) VALUES(?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, task.Id, task.Title, task.Content, task.Status)
	if err != nil {
		logging.Logger.Error("SQL実行に失敗", "error", err)
		return err
	}
	return nil
}

func (r *TaskRepository) Get(ctx context.Context, offset int, pageSize int, status, searchWord *string) ([]*task.Task, error) {
	logging.Logger.Info("Get実行", "offset:", offset)

	query := "SELECT id, title, content, status FROM tasks"
	var args []interface{}
	var conditions []string

	if status != nil {
		conditions = append(conditions, "status = ?")
		args = append(args, *status)
	}

	if searchWord != nil && *searchWord != "" {
		conditions = append(conditions, "(LOWER(title) LIKE ? OR LOWER(content) LIKE ?)")
		args = append(args, "%"+*searchWord+"%", "%"+*searchWord+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
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

func (r *TaskRepository) FindById(ctx context.Context, id string) (*task.Task, error) {
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

func (r *TaskRepository) Save(ctx context.Context, task *task.Task) error {
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
