package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type MySQL struct {
	Database string `ini:"database"`
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

var DBConfig = MySQL{}

type App struct {
	AppName   string `ini:"app_name"`
	PrefixUrl string `ini:"prefix_url"`
}

var AppConfig = App{}

func Setup() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	cfg.Section("mysql").MapTo(&DBConfig)
	cfg.Section("app").MapTo(&AppConfig)
}
