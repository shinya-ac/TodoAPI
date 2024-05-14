package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Host       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("config読み込み失敗：%v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Host:       cfg.Section("db").Key("host").String(),
		DBUser:     cfg.Section("db").Key("user").String(),
		DBPassword: cfg.Section("db").Key("password").String(),
		DBHost:     cfg.Section("db").Key("host").String(),
		DBPort:     cfg.Section("db").Key("port").String(),
		DBName:     cfg.Section("db").Key("name").String(),
	}
}
