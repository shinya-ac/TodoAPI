package health_handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shinya-ac/TodoAPI/presentation/health_handler"
	"github.com/stretchr/testify/assert"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/health", health_handler.HealthCheck)

	t.Run("健康チェックエンドポイントがステータス200を返し、ステータスが'ok'であることを確認", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response HealthResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response.Status)
	})
}
