package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	MYSQL_HOST     = "127.0.0.1"
	MYSQL_PORT     = 3306
	MYSQL_USER     = "root"
	MYSQL_PASSWORD = "123456"
	MYSQL_DB_NAME  = "test"
	MYSQL_CHARSET  = "utf8"
)

var SqlDB *sql.DB

func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB_NAME, MYSQL_CHARSET)
	SqlDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Println(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Println(err.Error())
	}
}

func Close(db *sql.DB) {
	_ = db.Close()
}
