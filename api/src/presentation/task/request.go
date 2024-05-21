package task

type CreateTaskParams struct {
	Title   string `json:"title" validate:"required" example:"読書"`
	Content string `json:"content" validate:"required" example:"「達人に学ぶクリーンアーキテクチャp100~105」までを読む"`
}
