package api_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/testfixtures.v2"

	"github.com/shinya-ac/TodoAPI/infrastructure/mysql/db"
	dbTest "github.com/shinya-ac/TodoAPI/infrastructure/mysql/db/db_test"
	"github.com/shinya-ac/TodoAPI/pkg/logging"
	"github.com/shinya-ac/TodoAPI/presentation/settings"
	"github.com/shinya-ac/TodoAPI/server/route"
)

var (
	fixtures *testfixtures.Context
	api      *gin.Engine
)

func TestMain(m *testing.M) {
	logging.InitLogger()
	var err error

	resource, pool := dbTest.CreateContainer()
	defer dbTest.CloseContainer(resource, pool)

	dbCon := dbTest.ConnectDB(resource, pool)
	defer dbCon.Close()

	// テスト用DBセットアップ
	dbTest.SetupTestDB(dbCon, "../../infrastructure/mysql/db/schema/schema.sql")

	// テストデータの準備
	fixtures, err = testfixtures.NewFolder(
		dbCon,
		&testfixtures.MySQL{},
		"../../infrastructure/mysql/fixtures",
	)
	if err != nil {
		panic(err)
	}
	if err := fixtures.Load(); err != nil {
		panic(err)
	}

	db.SetDB(dbCon)

	api = settings.NewGinEngine()
	route.InitRoute(api)

	m.Run()
}
