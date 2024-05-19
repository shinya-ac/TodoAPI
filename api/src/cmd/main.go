package main

import (
	"context"

	config "github.com/shinya-ac/TodoAPI/configs"
	"github.com/shinya-ac/TodoAPI/infrastructure/mysql/db"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	"github.com/shinya-ac/TodoAPI/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logging.InitLogger()

	db.NewMainDB(config.Config)

	server.Run(ctx, config.Config)
}
