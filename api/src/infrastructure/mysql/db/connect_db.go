package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	config "github.com/shinya-ac/TodoAPI/configs"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

const maxRetries = 5
const delay = 5 * time.Second

var (
	once  sync.Once
	dbcon *sql.DB
)

func SetDB(d *sql.DB) {
	dbcon = d
}

func GetDB() *sql.DB {
	return dbcon
}

func NewMainDB(cnf config.ConfigList) {
	once.Do(func() {
		var err error
		logging.Logger.Info("DBHost:", "", cnf.DBHost)
		dbcon, err := connect(
			cnf.DBUser,
			cnf.DBPassword,
			cnf.DBHost,
			cnf.DBPort,
			cnf.DBName,
		)
		if err != nil {
			logging.Logger.Error("DBの初期化に失敗", "error", err)
			panic(err)
		}
		SetDB(dbcon)
	})
}

func connect(user string, password string, host string, port string, name string) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)

		db, err := sql.Open("mysql", connect)
		if err != nil {
			logging.Logger.Error("MySQLの接続に失敗", "error", err)
			return nil, fmt.Errorf("DBに接続できません。: %w", err)
		}

		err = db.Ping()
		if err == nil {
			logging.Logger.Info("DB接続が確立")
			return db, nil
		}

		logging.Logger.Warn("DBへの接続に失敗。再試行中...", "attempt", i+1, "error", err)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("DBへの接続に %d 回失敗しました", maxRetries)
}
