## TodoAPI

### 前提
1. `config.ini`が`TodoAPI/api/src/config.ini`に置かれていること
2. `.env`が`TodoAPI/mysql/.env`に置かれていること
3. Dockerツールがバックグラウンドで起動状態にあること

### Setup
1. makeコマンドでdbとapi環境の立ち上げを行う

    `make setup`

2. 以下コマンドで`"status": "ok"`が返るか確認

```
curl --location 'http://127.0.0.1:8080/v1/health'
```


3. Swaggerが立ち上がるので以下でアクセスできる

    `http://localhost:8080/v1/swagger/index.html`

### Test
以下を実行してエラーがないことを確認する

`make test`

### トラブルシューティング
```
ERROR: Network todo_api_network declared as external, but could not be found. Please create the network manually using `docker network create todo_api_network` and try again.
```
上記はDockerNetworkのエラー
以下でNetworkを手動で作成する

`docker network create todo_api_network`

イメージ削除

`docker rmi todoapi_db`

Docker Volume削除

`docker volume rm mysql_todo_api_volume `

### 参考：本ツールで利用している関連パッケージ
gin

`go get github.com/gin-gonic/gin`

mysqlドライバー

`go get github.com/go-sql-driver/mysql`

swagコマンド

`go install github.com/swaggo/swag/cmd/swag@latest`

参考：https://github.com/swaggo/gin-swagger

Swagger on Gin導入

```
go get github.com/swaggo/files
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/swag/cmd/swag
```


cors

`go get "github.com/gin-contrib/cors"`

uuid

`go get "github.com/google/uuid"`
