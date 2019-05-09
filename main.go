package main

import db "demo/database"

func main() {
	//gin框架
	defer db.Close(db.SqlDB)
	//defer db.Close(db.PgDB)
	router := initRouter()
	_ = router.Run(":8018")
}
