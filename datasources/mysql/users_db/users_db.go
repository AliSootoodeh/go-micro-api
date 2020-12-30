package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	//UsersDB ...
	UsersDB *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1", "users_db")
	var err error
	UsersDB, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = UsersDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
