package health_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/shinya-ac/TodoAPI/presentation/settings"
)

func HealthCheck(ctx *gin.Context) {
	res := HealthResponse{
		Status: "ok",
	}
	settings.ReturnStatusOK(ctx, res)
}
