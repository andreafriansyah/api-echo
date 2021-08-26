package db

import (
	config "api-echo/Config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + "@/" + conf.DB_NAME

	// "root:@/bolang-api"
	_, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Connection error")
	}

}

func CreateCon() *sql.DB {
	return db
}
