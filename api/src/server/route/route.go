package route

import (
	ginpkg "github.com/gin-gonic/gin"

	taskApp "github.com/shinya-ac/TodoAPI/application/task"
	"github.com/shinya-ac/TodoAPI/infrastructure/mysql/db"
	"github.com/shinya-ac/TodoAPI/infrastructure/mysql/repository"
	"github.com/shinya-ac/TodoAPI/presentation/health_handler"
	"github.com/shinya-ac/TodoAPI/presentation/settings"
	taskPre "github.com/shinya-ac/TodoAPI/presentation/task"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func InitRoute(api *ginpkg.Engine) {
	api.Use(settings.ErrorHandler())

	api.GET("/v1/health", health_handler.HealthCheck)
	api.GET("/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protectedV1 := api.Group("/v1")
	protectedV1.Use(settings.ApiKeyAuthMiddleware())
	{
		taskRoute(protectedV1)
	}
}

func taskRoute(r *ginpkg.RouterGroup) {
	taskRepository := repository.NewTaskRepository(db.GetDB())
	cuc := taskApp.NewCreateTaskUseCase(taskRepository)
	guc := taskApp.NewGetTaskUseCase(taskRepository)
	uuc := taskApp.NewUpdateTaskUseCase(taskRepository)
	duc := taskApp.NewDeleteTaskUseCase(taskRepository)
	h := taskPre.NewHandler(cuc, guc, uuc, duc)

	group := r.Group("/tasks")
	group.POST("/", h.CreateTasks)
	group.GET("/", h.GetTasks)
	group.PUT("/:id", h.UpdateTasks)
	group.DELETE("/:id", h.DeleteTasks)
}
