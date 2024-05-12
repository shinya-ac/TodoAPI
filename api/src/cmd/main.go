package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	config "github.com/shinya-ac/TodoAPI/configs"
	"github.com/shinya-ac/TodoAPI/internal/infrastructure/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello slog", "name", "slog")
	fmt.Println(config.Config.Host)

	mux := router.NewRouter()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Error("サーバー起動エラー：", err)
	}
}
