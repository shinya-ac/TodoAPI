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
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		taskRoute(v1)
	}
}

func taskRoute(r *ginpkg.RouterGroup) {
	taskRepository := repository.NewTaskRepository(db.GetDB())
	cuc := taskApp.NewCreateTaskUseCase(taskRepository)
	guc := taskApp.NewGetTaskUseCase(taskRepository)
	uuc := taskApp.NewUpdateTaskUseCase(taskRepository)
	h := taskPre.NewHandler(cuc, guc, uuc)

	group := r.Group("/task")
	group.POST("/", h.CreateTask)
	group.GET("/", h.GetTasks)
	group.PUT("/:id", h.UpdateTasks)
}
