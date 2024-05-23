package task

import (
	"unicode/utf8"

	"github.com/google/uuid"

	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type Task struct {
	Id      string
	Title   string
	Content string
}

func NewTask(
	Title string,
	Content string,
) (*Task, error) {
	if utf8.RuneCountInString(Title) < titleLengthMin || utf8.RuneCountInString(Title) > titleLengthMax {
		err := errDomain.NewError("タイトルの値が不正です。")
		logging.Logger.Error("タイトルの値が不正", "error", err)
		return nil, err
	}
	if utf8.RuneCountInString(Content) < contentLengthMin || utf8.RuneCountInString(Content) > contentLengthMax {
		err := errDomain.NewError("説明文の値が不正です。")
		logging.Logger.Error("説明文の値が不正", "error", err)
		return nil, err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		logging.Logger.Error("UUIDの生成に失敗", "error", err)
		return nil, err
	}
	return &Task{
		Id:      id.String(),
		Title:   Title,
		Content: Content,
	}, nil
}

func (t *Task) UpdateTask(newTitle, newContent string) error {
	if err := t.UpdateTitle(newTitle); err != nil {
		return err
	}
	if err := t.UpdateContent(newContent); err != nil {
		return err
	}
	return nil
}

func (t *Task) UpdateTitle(newTitle string) error {
	if utf8.RuneCountInString(newTitle) < titleLengthMin || utf8.RuneCountInString(newTitle) > titleLengthMax {
		err := errDomain.NewError("タイトルの値が不正です。")
		logging.Logger.Error("タイトルの値が不正", "error", err)
		return err
	}
	t.Title = newTitle
	return nil
}

func (t *Task) UpdateContent(newContent string) error {
	if utf8.RuneCountInString(newContent) < contentLengthMin || utf8.RuneCountInString(newContent) > contentLengthMax {
		err := errDomain.NewError("説明文の値が不正です。")
		logging.Logger.Error("説明文の値が不正", "error", err)
		return err
	}
	t.Content = newContent
	return nil
}

func (t *Task) GetId() string {
	return t.Id
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetContent() string {
	return t.Content
}

const (
	titleLengthMin = 1
	titleLengthMax = 50

	contentLengthMin = 1
	contentLengthMax = 3000
)
