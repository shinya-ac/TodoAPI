package router

import (
	"database/sql"
	"fmt"
	"net/http"

	config "github.com/shinya-ac/TodoAPI/configs"
	"github.com/shinya-ac/TodoAPI/internal/adapter/handler"
	"github.com/shinya-ac/TodoAPI/internal/infrastructure/mysql"
	"github.com/shinya-ac/TodoAPI/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func NewRouter() *http.ServeMux {
	dbUser := config.Config.DBUser
	dbPassword := config.Config.DBPassword
	dbHost := config.Config.DBHost
	dbPort := config.Config.DBPort
	dbName := config.Config.DBName

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {

	}
	taskRepo := mysql.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", handler.HealthCheckHandler)
	mux.HandleFunc("/todo", taskHandler.HandleCreateTask)
	return mux
}
