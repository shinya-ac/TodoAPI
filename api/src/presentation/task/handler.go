package task

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/shinya-ac/TodoAPI/application/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	validator "github.com/shinya-ac/TodoAPI/pkg/validator"
	"github.com/shinya-ac/TodoAPI/presentation/settings"
)

type handler struct {
	createTaskUseCase *task.CreateTaskUseCase
	getTaskUseCase    *task.GetTaskUseCase
	updateTaskUseCase *task.UpdateTaskUseCase
}

func NewHandler(
	createTaskUseCase *task.CreateTaskUseCase,
	getTaskUseCase *task.GetTaskUseCase,
	updateTaskUseCase *task.UpdateTaskUseCase,
) handler {
	return handler{
		createTaskUseCase: createTaskUseCase,
		getTaskUseCase:    getTaskUseCase,
		updateTaskUseCase: updateTaskUseCase,
	}
}

// CreateTask godoc
// @Summary Taskを登録する
// @Tags Task
// @Accept json
// @Produce json
// @Param request body CreateTaskParams true "Task登録"
// @Success 201 {object} createTaskResponse
// @Router /v1/task [post]
func (h handler) CreateTask(ctx *gin.Context) {
	logging.Logger.Info("CreateTask実行開始")
	var params CreateTaskParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		logging.Logger.Error("paramsの形式が不正", "error", err)
		settings.ReturnBadRequest(ctx, err)
		return
	}
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		logging.Logger.Error("paramsの内容が不正", "error", err)
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	input := task.CreateTaskUseCaseInputDto{
		Title:   params.Title,
		Content: params.Content,
	}

	dto, err := h.createTaskUseCase.Run(ctx, input)
	if err != nil {
		logging.Logger.Error("usecaseの実行に失敗", "error", err)
		settings.ReturnError(ctx, err)
		return
	}

	response := createTaskResponse{
		TaskId: dto.Id,
	}
	settings.ReturnStatusCreated(ctx, response)
}

// GetTask godoc
// @Summary Task一覧を取得する
// @Tags Task
// @Produce json
// @Param page query int false "ページ数" default(1)
// @Param pageSize query int false "ページサイズ" default(100)
// @Success 200 {object} getTaskResponse
// @Router /v1/task [get]
func (h handler) GetTasks(ctx *gin.Context) {
	logging.Logger.Info("GetTasks実行開始")

	page := 1
	pageSize := 100

	if p := ctx.Query("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if ps := ctx.Query("pageSize"); ps != "" {
		pageSize, _ = strconv.Atoi(ps)
	}

	offset := (page - 1) * pageSize

	input := task.GetTaskUseCaseInputDto{
		Offset:   offset,
		PageSize: pageSize,
	}

	dto, err := h.getTaskUseCase.Run(ctx, input)
	if err != nil {
		logging.Logger.Error("usecaseの実行に失敗", "error", err)
		settings.ReturnError(ctx, err)
		return
	}

	response := getTaskResponse{
		Tasks: dto.Tasks,
	}
	settings.ReturnStatusOK(ctx, response)
}

// GetTask godoc
// @Summary Taskを更新する
// @Tags Task
// @Produce json
// @Param request body UpdateTaskParams true "Task更新"
// @Success 200 {object} updateTaskResponse
// @Router /v1/task [put]
func (h handler) UpdateTasks(ctx *gin.Context) {
	logging.Logger.Info("UpdateTasks実行開始")
	var params UpdateTaskParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		logging.Logger.Error("paramsの形式が不正", "error", err)
		settings.ReturnBadRequest(ctx, err)
		return
	}
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		logging.Logger.Error("paramsの内容が不正", "error", err)
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	input := task.UpdateTaskUseCaseInputDto{
		Id:      params.Id,
		Title:   params.Title,
		Content: params.Content,
	}

	dto, err := h.updateTaskUseCase.Run(ctx, input)
	if err != nil {
		logging.Logger.Error("usecaseの実行に失敗", "error", err)
		settings.ReturnError(ctx, err)
		return
	}

	response := updateTaskResponse{
		TaskId: dto.Id,
	}
	settings.ReturnStatusOK(ctx, response)
}
