package config

import (
	"os"

	"gopkg.in/ini.v1"

	"github.com/shinya-ac/TodoAPI/pkg/logging"
)

type ConfigList struct {
	Host          string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	ServerAddress string
	ServerPort    string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		logging.Logger.Error("configの読み込みに失敗", "error", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Host:          cfg.Section("db").Key("host").String(),
		DBUser:        cfg.Section("db").Key("user").String(),
		DBPassword:    cfg.Section("db").Key("password").String(),
		DBHost:        cfg.Section("db").Key("host").String(),
		DBPort:        cfg.Section("db").Key("port").String(),
		DBName:        cfg.Section("db").Key("name").String(),
		ServerAddress: cfg.Section("server").Key("address").String(),
		ServerPort:    cfg.Section("server").Key("port").String(),
	}
}
