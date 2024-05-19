package task

type CreateTaskParams struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
