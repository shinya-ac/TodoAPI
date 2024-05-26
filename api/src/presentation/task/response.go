package task

import "github.com/shinya-ac/TodoAPI/domain/task"

type createTaskResponse struct {
	TaskId string `json:"id" example:"4082ed31-263c-40ec-9d41-e9d274c6bca8"`
}

type getTaskResponse struct {
	Tasks []*task.Task `json:"tasks"`
}

type updateTaskResponse struct {
	TaskId string `json:"id" example:"4082ed31-263c-40ec-9d41-e9d274c6bca8"`
}

type deleteTaskResponse struct {
	TaskId string `json:"id" example:"4082ed31-263c-40ec-9d41-e9d274c6bca8"`
}
