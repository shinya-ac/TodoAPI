package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	config "github.com/shinya-ac/TodoAPI/configs"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello slog", "name", "slog")
	fmt.Println(config.Config.Host)

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバー起動エラー：", err)
		return
	}
}
