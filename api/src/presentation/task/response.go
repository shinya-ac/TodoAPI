package task

import "github.com/google/uuid"

type createTaskResponse struct {
	TaskId uuid.UUID `json:"id"`
}
