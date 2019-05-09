package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	PG_HOST     = "127.0.0.1"
	PG_PORT     = 5432
	PG_USER     = "xiangqian"
	PG_PASSWORD = ""
	PG_DB_NAME  = "test"
)

var PgDB *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PG_HOST, PG_PORT, PG_USER, PG_PASSWORD, PG_DB_NAME)
	PgDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("222")
	err = PgDB.Ping()
	if err != nil {
		log.Println(err.Error())
	}
}
