package task

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/shinya-ac/TodoAPI/application/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	"github.com/shinya-ac/TodoAPI/pkg/utils"
	validator "github.com/shinya-ac/TodoAPI/pkg/validator"
	"github.com/shinya-ac/TodoAPI/presentation/settings"
)

type handler struct {
	createTaskUseCase *task.CreateTaskUseCase
	getTaskUseCase    *task.GetTaskUseCase
	updateTaskUseCase *task.UpdateTaskUseCase
	deleteTaskUseCase *task.DeleteTaskUseCase
}

func NewHandler(
	createTaskUseCase *task.CreateTaskUseCase,
	getTaskUseCase *task.GetTaskUseCase,
	updateTaskUseCase *task.UpdateTaskUseCase,
	deleteTaskUseCase *task.DeleteTaskUseCase,
) handler {
	return handler{
		createTaskUseCase: createTaskUseCase,
		getTaskUseCase:    getTaskUseCase,
		updateTaskUseCase: updateTaskUseCase,
		deleteTaskUseCase: deleteTaskUseCase,
	}
}

// CreateTasks godoc
// @Summary Taskを登録する
// @Description パラメーターから新規Todoを作成する
// @Tags Task
// @Accept json
// @Produce json
// @Param request body CreateTaskParams true "Task登録"
// @Success 201 {object} createTaskResponse
// @Router /v1/tasks [post]
// @Security ApiKeyAuth
func (h handler) CreateTasks(ctx *gin.Context) {
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

// GetTasks godoc
// @Summary Task一覧を取得する
// @Description 検索ワードとstatusの指定でフィルターをかけてTodoの一覧を取得する
// @Tags Task
// @Produce json
// @Param page query int false "ページ数" default(1)
// @Param pageSize query int false "ページサイズ" default(100)
// @Param status query string false "Todoのステータス(Pending, InProgress, Completed)" default(Pending)
// @Param searchWord query string false "Todoのタイトルまたは説明文から検索する検索ワード"
// @Success 200 {object} getTaskResponse
// @Router /v1/tasks [get]
// @Security ApiKeyAuth
func (h handler) GetTasks(ctx *gin.Context) {
	logging.Logger.Info("GetTasks実行開始")

	page := 1
	pageSize := 100
	var status *string
	var searchWord *string

	if p := ctx.Query("page"); p != "" {
		page, _ = strconv.Atoi(p)
	}
	if ps := ctx.Query("pageSize"); ps != "" {
		pageSize, _ = strconv.Atoi(ps)
	}
	if st := ctx.Query("status"); st != "" {
		status = &st
		v := validator.GetValidator()
		err := v.Var(*status, "status")
		if err != nil {
			logging.Logger.Error("statusが有効ではない", "error", err)
			settings.ReturnBadRequest(ctx, err)
			return
		}
	}
	if sw := ctx.Query("searchWord"); sw != "" {
		normalizedSearchWord := fmt.Sprintf("%%%s%%", utils.NormalizeString(sw))
		searchWord = &normalizedSearchWord
	}

	offset := (page - 1) * pageSize

	input := task.GetTaskUseCaseInputDto{
		Offset:     offset,
		PageSize:   pageSize,
		Status:     status,
		SearchWord: searchWord,
	}

	dto, err := h.getTaskUseCase.Run(ctx, input)
	if err != nil {
		logging.Logger.Error("usecaseの実行に失敗", "error", err)
		settings.ReturnError(ctx, err)
		return
	}

	taskCount := len(dto.Tasks)
	taskCountStr := strconv.Itoa(taskCount)

	response := getTaskResponse{
		Tasks:      dto.Tasks,
		TotalTasks: taskCountStr,
	}
	settings.ReturnStatusOK(ctx, response)
}

// UpdateTasks godoc
// @Summary Taskを更新する
// @Description 指定されたIDに合致するTodoのタイトルやステータスを更新する
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "更新するTodoを指定するid"
// @Param request body UpdateTaskParams true "Task更新"
// @Success 200 {object} updateTaskResponse
// @Router /v1/tasks/{id} [put]
// @Security ApiKeyAuth
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

	id := ctx.Param("id")

	input := task.UpdateTaskUseCaseInputDto{
		Id:      id,
		Title:   params.Title,
		Content: params.Content,
		Status:  params.Status,
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

// DeleteTask godoc
// @Summary Taskを削除する
// @Description 指定されたIDに合致するTodoを削除する
// @Tags Task
// @Produce json
// @Param id path string true "削除するTodoを指定するid"
// @Success 200 {object} deleteTaskResponse
// @Router /v1/tasks/{id} [delete]
// @Security ApiKeyAuth
func (h handler) DeleteTasks(ctx *gin.Context) {
	logging.Logger.Info("DeleteTasks実行開始")

	id := ctx.Param("id")

	input := task.DeleteTaskUseCaseInputDto{
		Id: id,
	}

	dto, err := h.deleteTaskUseCase.Run(ctx, input)
	if err != nil {
		logging.Logger.Error("usecaseの実行に失敗", "error", err)
		settings.ReturnError(ctx, err)
		return
	}

	response := deleteTaskResponse{
		TaskId: dto.Id,
	}
	settings.ReturnStatusOK(ctx, response)
}
