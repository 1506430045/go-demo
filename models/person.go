package models

import (
	db "demo/database"
	"fmt"
	"log"
)

type Person struct {
	Id        int    `json:"id" from:"id"`
	LastName  string `json:"last_name" from:"last_name"`
	FirstName string `json:"first_name" from:"first_name"`
}

//新增
func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person (first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

//查询单个
func (p *Person) GetPerson() {

	err := db.SqlDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", p.Id).Scan(&p.Id, &p.FirstName, &p.LastName)

	if err != nil {
		log.Println(err)
		return
	}
	return
}

//查询所有
func (p *Person) GetPersons() []Person {
	persons := make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()

	if err != nil {
		return persons
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return persons
	}
	return persons
}

//修改
func (p *Person) ModPerson() bool {
	stmt, err := db.SqlDB.Prepare("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?")
	if err != nil {
		log.Println(err)
		return false
	}
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		log.Println(err)
		return false
	}
	ra, err := rs.RowsAffected()
	if ra <= 0 || err != nil {
		log.Println(err)
		return false
	}
	return true
}

//删除
func (p Person) DelPerson() bool {
	stmt, err := db.SqlDB.Prepare("DELETE FROM person WHERE id = ?")
	if err != nil {
		log.Println(err)
		return false
	}
	rs, err := stmt.Exec(p.Id)

	fmt.Println(p.Id)

	if err != nil {
		log.Println(err)
		return false
	}

	ra, err := rs.RowsAffected()

	if ra <= 0 || err != nil {
		log.Println(err)
		return false
	}

	return true
}
