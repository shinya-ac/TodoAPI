package main

import (
	"fmt"
	"log/slog"
	"os"

	config "github.com/shinya-ac/TodoAPI/configs"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello slog", "name", "slog")
	fmt.Println(config.Config.Host)
}
