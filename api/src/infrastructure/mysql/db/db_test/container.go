package dbTest

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ory/dockertest"
)

var (
	username = "root"
	password = "password"
	hostname = "localhost"
	dbName   = "todo_api_test"
	port     int
)

func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Dockerに接続できませんでした。: %s", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + password,
			"MYSQL_DATABASE=" + dbName,
		},
		Mounts: []string{},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("リソースが起動しません。: %s", err)
	}

	return resource, pool
}

func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("リソースを解放できません。: %s", err)
	}
}

func ConnectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	var db *sql.DB
	if err := pool.Retry(func() error {
		time.Sleep(time.Second * 3)
		var err error
		port, err = strconv.Atoi(resource.GetPort("3306/tcp"))
		if err != nil {
			return err
		}
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", username, password, hostname, resource.GetPort("3306/tcp"), dbName))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Databaseに接続できません。: %s", err)
	}
	return db
}

func SetupTestDB(db *sql.DB, schemaFilePath string) {
	schema, err := os.ReadFile(schemaFilePath)
	if err != nil {
		log.Fatalf("スキーマファイルを読み取れません。: %s", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		log.Fatalf("スキーマを実行できません: %s", err)
	}
}
