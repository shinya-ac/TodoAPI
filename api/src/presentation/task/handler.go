package task

import (
	"github.com/gin-gonic/gin"

	"github.com/shinya-ac/TodoAPI/application/task"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	validator "github.com/shinya-ac/TodoAPI/pkg/validator"
	"github.com/shinya-ac/TodoAPI/presentation/settings"
)

type handler struct {
	createTaskUseCase *task.CreateTaskUseCase
}

func NewHandler(
	createTaskUseCase *task.CreateTaskUseCase,
) handler {
	return handler{
		createTaskUseCase: createTaskUseCase,
	}
}

func (h handler) CreateTask(ctx *gin.Context) {
	logging.Logger.Info("CreateTaskエンドポイント実行開始")
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
