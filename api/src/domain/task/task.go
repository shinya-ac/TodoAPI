package task

import (
	"unicode/utf8"

	"github.com/google/uuid"

	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type Task struct {
	Id      uuid.UUID
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
	id, err := uuid.NewUUID()
	if err != nil {
		logging.Logger.Error("UUIDの生成に失敗", "error", err)
		return nil, err
	}
	return &Task{
		Id:      id,
		Title:   Title,
		Content: Content,
	}, nil
}

func (t *Task) GetId() uuid.UUID {
	return t.Id
}

const (
	titleLengthMin = 1
	titleLengthMax = 50

	contentLengthMin = 1
	contentLengthMax = 3000
)
