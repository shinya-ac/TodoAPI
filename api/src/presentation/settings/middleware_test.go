package settings_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	config "github.com/shinya-ac/TodoAPI/configs"
	errDomain "github.com/shinya-ac/TodoAPI/domain/error"
	"github.com/shinya-ac/TodoAPI/presentation/settings"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(settings.ErrorHandler())

	router.GET("/test", func(c *gin.Context) {
		c.Error(errDomain.NotFoundErr)
		c.Error(errors.New("internal error"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestApiKeyAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	config.Config = config.ConfigList{
		APIKey1: "valid-key-1",
		APIKey2: "valid-key-2",
		APIKey3: "valid-key-3",
	}

	router.Use(settings.ApiKeyAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	t.Run("有効なAPIキーで認証されることを確認", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Todo-API-Key", "valid-key-1")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "ok"}`, w.Body.String())
	})

	t.Run("無効なAPIキーで認証されないことを確認", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Todo-API-Key", "invalid-key")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "APIkeyが有効ではありません。"}`, w.Body.String())
	})

	t.Run("APIキーがない場合に認証されないことを確認", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"error": "APIkeyが有効ではありません。"}`, w.Body.String())
	})
}
