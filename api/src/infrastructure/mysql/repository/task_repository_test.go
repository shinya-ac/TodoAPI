package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/shinya-ac/TodoAPI/domain/task"
	taskDomain "github.com/shinya-ac/TodoAPI/domain/task"
	repository "github.com/shinya-ac/TodoAPI/infrastructure/mysql/repository"
	"github.com/shinya-ac/TodoAPI/pkg/logging"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (task.TaskRepository, sqlmock.Sqlmock, *sql.DB, context.Context) {
	logging.InitLogger()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := repository.NewTaskRepository(db)
	ctx := context.Background()

	return repo, mock, db, ctx
}

func TestCreate(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()

	task := &taskDomain.Task{
		Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
		Title:   "Todoのテストを行う",
		Content: "Todo機能のテストをGo言語で行う",
		Status:  "Pending",
	}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs(task.Id, task.Title, task.Content, task.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(ctx, task)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateWithNilTask(t *testing.T) {
	repo, _, db, ctx := setup(t)
	defer db.Close()

	err := repo.Create(ctx, nil)
	assert.Error(t, err)
	assert.Equal(t, "Taskがnilです。", err.Error())
}

func TestGet(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()

	status := "Pending"
	rows := sqlmock.NewRows([]string{"id", "title", "content", "status"}).
		AddRow("46039334-6ffc-4fe3-ab59-f40a7b73b611", "Todoのテストを行う", "Todo機能のテストをGo言語で行う", "Pending").
		AddRow("46039334-6ffc-4fe3-ab59-f40a7b73b612", "Todoのテストを行う2", "Todo機能のテストをGo言語で行う2", "Completed")

	query := "SELECT id, title, content, status FROM tasks WHERE status = \\? ORDER BY created_at DESC LIMIT \\? OFFSET \\?"
	mock.ExpectQuery(query).
		WithArgs(status, 2, 0).
		WillReturnRows(rows)

	tasks, err := repo.Get(ctx, 0, 2, &status)
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "46039334-6ffc-4fe3-ab59-f40a7b73b611", tasks[0].Id)
	assert.Equal(t, "Todoのテストを行う", tasks[0].Title)
	assert.Equal(t, "Pending", tasks[0].Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindById(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()

	query := "SELECT id, title, content, status FROM tasks WHERE id = \\?"
	mock.ExpectQuery(query).
		WithArgs("46039334-6ffc-4fe3-ab59-f40a7b73b611").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "status"}).AddRow("46039334-6ffc-4fe3-ab59-f40a7b73b611", "Todoのテストを行う", "Todo機能のテストをGo言語で行う", "Pending"))

	result, err := repo.FindById(ctx, "46039334-6ffc-4fe3-ab59-f40a7b73b611")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "46039334-6ffc-4fe3-ab59-f40a7b73b611", result.Id)
	assert.Equal(t, "Todoのテストを行う", result.Title)
	assert.Equal(t, "Pending", result.Status)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSave(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()

	task := &taskDomain.Task{
		Id:      "46039334-6ffc-4fe3-ab59-f40a7b73b611",
		Title:   "Todoのテストを行う",
		Content: "Todo機能のテストをGo言語で行う",
		Status:  "Completed",
	}

	query := `
		INSERT INTO tasks \(id, title, content, status\)
		VALUES \(\?, \?, \?, \?\)
		ON DUPLICATE KEY UPDATE
			title = VALUES\(title\),
			content = VALUES\(content\),
			status = VALUES\(status\),
			updated_at = NOW\(\)`
	mock.ExpectExec(query).
		WithArgs(task.Id, task.Title, task.Content, task.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Save(ctx, task)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetWithDBError(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()
	status := "Pending"

	query := "SELECT id, title, content, status FROM tasks WHERE status = \\? ORDER BY created_at DESC LIMIT \\? OFFSET \\?"
	mock.ExpectQuery(query).
		WithArgs(status, 2, 0).
		WillReturnError(errors.New("DB error"))

	tasks, err := repo.Get(ctx, 0, 2, &status)
	assert.Error(t, err)
	assert.Nil(t, tasks)
	assert.Equal(t, "DB error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByIdWithDBError(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()

	query := "SELECT id, title, content, status FROM tasks WHERE id = \\?"
	mock.ExpectQuery(query).
		WithArgs("1").
		WillReturnError(errors.New("DB error"))

	result, err := repo.FindById(ctx, "1")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "DB error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveWithDBError(t *testing.T) {
	repo, mock, db, ctx := setup(t)
	defer db.Close()
	task := &taskDomain.Task{
		Id:      "1",
		Title:   "Updated Task",
		Content: "Updated Content",
		Status:  "Completed",
	}

	query := `
		INSERT INTO tasks \(id, title, content, status\)
		VALUES \(\?, \?, \?, \?\)
		ON DUPLICATE KEY UPDATE
			title = VALUES\(title\),
			content = VALUES\(content\),
			status = VALUES\(status\),
			updated_at = NOW\(\)`
	mock.ExpectExec(query).
		WithArgs(task.Id, task.Title, task.Content, task.Status).
		WillReturnError(errors.New("DB error"))

	err := repo.Save(ctx, task)
	assert.Error(t, err)
	assert.Equal(t, "DB error", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}
