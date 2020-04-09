package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_api_server/config"
	"log"
	"strconv"
)

var (
	DB *sql.DB
)

func InitMySQL() (err error) {
	dbConf := config.DBConfig
	dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.IP + ":" + strconv.Itoa(dbConf.Port) + ")/" + dbConf.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("open db failded, error: %v\n", err)
		return
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("open %s failed,error : %v\n", dsn, err)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	log.Println("open db success.")
	return
}

func Close() {
	DB.Close()
}
