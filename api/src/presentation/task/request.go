package task

type CreateTaskParams struct {
	Title   string `json:"title" validate:"required" example:"読書"`
	Content string `json:"content" validate:"required" example:"「達人に学ぶクリーンアーキテクチャp100~105」までを読む"`
}

type UpdateTaskParams struct {
	Title   *string `json:"title" validate:"omitempty,min=1,max=100" example:"輪読会"`
	Content *string `json:"content" validate:"omitempty,min=1,max=1000" example:"「達人に学ぶクリーンアーキテクチャp200~300」までを読む"`
	Status  *string `json:"status" validate:"omitempty,status" example:"Completed"`
}
