package main

import (
	"context"

	config "github.com/shinya-ac/TodoAPI/configs"
	"github.com/shinya-ac/TodoAPI/infrastructure/mysql/db"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	"github.com/shinya-ac/TodoAPI/server"

	_ "github.com/shinya-ac/TodoAPI/docs"
)

// @title           Todo API
// @version         1.0
// @description     RESTful API for TodoApp
// @termsOfService  localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Todo-API-Key

func main() {
	logging.InitLogger()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		logging.Logger.Error("configの読み込みに失敗", "error", err)
	}

	db.NewMainDB(cfg)

	server.Run(ctx, cfg)
}
