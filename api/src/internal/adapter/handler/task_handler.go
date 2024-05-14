package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shinya-ac/TodoAPI/internal/domain"
	"github.com/shinya-ac/TodoAPI/internal/usecase"
)

type TaskHandler struct {
	Usecase *usecase.TaskUsecase
}

func NewTaskHandler(usecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{Usecase: usecase}
}

func (h *TaskHandler) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.Usecase.CreateTask(r.Context(), task)
	if err != nil {
		http.Error(w, "Todoの作成失敗：", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}
