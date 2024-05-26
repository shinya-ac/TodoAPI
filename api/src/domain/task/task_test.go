package task_test

import (
	"testing"

	"github.com/shinya-ac/TodoAPI/domain/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	logging.InitLogger()
	t.Run("正常にタスクを作成できることを確認", func(t *testing.T) {
		title := "Todoのテストを行う1"
		content := "Todo機能のテストをGo言語で行う1"
		tk, err := task.NewTask(title, content)
		assert.NoError(t, err)
		assert.NotNil(t, tk)
		assert.Equal(t, title, tk.GetTitle())
		assert.Equal(t, content, tk.GetContent())
		assert.Equal(t, "Pending", tk.Status)
	})

	t.Run("タイトルが不正な場合にエラーを返すことを確認", func(t *testing.T) {
		title := ""
		content := "Todo機能のテストをGo言語で行う1"
		tk, err := task.NewTask(title, content)
		assert.Error(t, err)
		assert.Nil(t, tk)
		assert.Equal(t, "タイトルの値が不正です。", err.Error())
	})

	t.Run("コンテンツが不正な場合にエラーを返すことを確認", func(t *testing.T) {
		title := "Todoのテストを行う1"
		content := ""
		tk, err := task.NewTask(title, content)
		assert.Error(t, err)
		assert.Nil(t, tk)
		assert.Equal(t, "説明文の値が不正です。", err.Error())
	})
}

func TestUpdateTask(t *testing.T) {
	tk := &task.Task{
		Id:      "test-id",
		Title:   "Todoのテストを行う1",
		Content: "Todo機能のテストをGo言語で行う1",
		Status:  "Pending",
	}

	t.Run("タイトル、コンテンツ、ステータスを更新できることを確認", func(t *testing.T) {
		newTitle := "Todoのテストを行う1-2"
		newContent := "Todo機能のテストをGo言語で行う1-2"
		newStatus := "InProgress"
		err := tk.UpdateTask(&newTitle, &newContent, &newStatus)
		assert.NoError(t, err)
		assert.Equal(t, newTitle, tk.GetTitle())
		assert.Equal(t, newContent, tk.GetContent())
		assert.Equal(t, newStatus, tk.Status)
	})

	t.Run("無効なタイトルを渡した場合にエラーを返すことを確認", func(t *testing.T) {
		newTitle := ""
		err := tk.UpdateTask(&newTitle, nil, nil)
		assert.Error(t, err)
		assert.Equal(t, "タイトルの値が不正です。", err.Error())
	})

	t.Run("無効なコンテンツを渡した場合にエラーを返すことを確認", func(t *testing.T) {
		newContent := ""
		err := tk.UpdateTask(nil, &newContent, nil)
		assert.Error(t, err)
		assert.Equal(t, "説明文の値が不正です。", err.Error())
	})

	t.Run("無効なステータスを渡した場合にエラーを返すことを確認", func(t *testing.T) {
		newStatus := "InvalidStatus"
		err := tk.UpdateTask(nil, nil, &newStatus)
		assert.Error(t, err)
		assert.Equal(t, "タスクの状態の値が不正です。", err.Error())
	})
}

func TestUpdateTitle(t *testing.T) {
	tk := &task.Task{
		Id:      "test-id",
		Title:   "Todoのテストを行う1",
		Content: "Todo機能のテストをGo言語で行う1",
		Status:  "Pending",
	}

	t.Run("タイトルを更新できることを確認", func(t *testing.T) {
		newTitle := "Todoのテストを行う1-2"
		err := tk.UpdateTitle(newTitle)
		assert.NoError(t, err)
		assert.Equal(t, newTitle, tk.GetTitle())
	})

	t.Run("無効なタイトルを渡した場合にエラーを返すことを確認", func(t *testing.T) {
		newTitle := ""
		err := tk.UpdateTitle(newTitle)
		assert.Error(t, err)
		assert.Equal(t, "タイトルの値が不正です。", err.Error())
	})
}

func TestUpdateContent(t *testing.T) {
	tk := &task.Task{
		Id:      "test-id",
		Title:   "Todoのテストを行う1",
		Content: "Todo機能のテストをGo言語で行う1",
		Status:  "Pending",
	}

	t.Run("コンテンツを更新できることを確認", func(t *testing.T) {
		newContent := "Todo機能のテストをGo言語で行う1-2"
		err := tk.UpdateContent(newContent)
		assert.NoError(t, err)
		assert.Equal(t, newContent, tk.GetContent())
	})

	t.Run("無効なコンテンツを渡した場合にエラーを返すことを確認", func(t *testing.T) {
		newContent := ""
		err := tk.UpdateContent(newContent)
		assert.Error(t, err)
		assert.Equal(t, "説明文の値が不正です。", err.Error())
	})
}

func TestUpdateStatus(t *testing.T) {
	tk := &task.Task{
		Id:      "test-id",
		Title:   "Todoのテストを行う1",
		Content: "Todo機能のテストをGo言語で行う1",
		Status:  "Pending",
	}

	t.Run("ステータスを更新できることを確認", func(t *testing.T) {
		newStatus := "Completed"
		err := tk.UpdateStatus(newStatus)
		assert.NoError(t, err)
		assert.Equal(t, newStatus, tk.Status)
	})

	t.Run("無効なステータスを渡した場合にエラーを返すことを確認", func(t *testing.T) {
		newStatus := "InvalidStatus"
		err := tk.UpdateStatus(newStatus)
		assert.Error(t, err)
		assert.Equal(t, "タスクの状態の値が不正です。", err.Error())
	})
}
