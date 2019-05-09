package models

import (
	db "demo/database"
	"log"
)

const (
	PG_HOST     = "127.0.0.1"
	PG_PORT     = 5432
	PG_USER     = "xiangqian"
	PG_PASSWORD = ""
	PG_DB_NAME  = "test"
)

type Employee struct {
	Id      int    `json:"id" from:"id"`
	Name    string `json:"name" from:"name"`
	Age     int    `json:"age" from:"age"`
	Address string `json:"address" from:"address"`
	Salary  string `json:"salary" from:"salary"`
}

func (e *Employee) GetEmployee() {
	//fmt.Println(db.PgDB.Ping())
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PG_HOST, PG_PORT, PG_USER, PG_PASSWORD, PG_DB_NAME)
	//PgDB, _ := sql.Open("postgres", psqlInfo)
	query := "SELECT id, Address, age FROM employee WHERE id=$1"
	err := db.PgDB.QueryRow(query, e.Id).Scan(&e.Id, &e.Address, &e.Age)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
