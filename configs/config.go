package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Host string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("config読み込み失敗：%v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Host: cfg.Section("db").Key("host").String(),
	}
}
