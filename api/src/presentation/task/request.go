package task

type CreateTaskParams struct {
	Title   string `json:"title" validate:"required" example:"読書"`
	Content string `json:"content" validate:"required" example:"「達人に学ぶクリーンアーキテクチャp100~105」までを読む"`
}

type UpdateTaskParams struct {
	Id      string `json:"id" example:"4082ed31-263c-40ec-9d41-e9d274c6bca8"`
	Title   string `json:"title" validate:"required" example:"輪読会"`
	Content string `json:"content" validate:"required" example:"「達人に学ぶクリーンアーキテクチャp200~300」までを読む"`
}
