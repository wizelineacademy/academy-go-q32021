package datastore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (conn *sql.DB) {
	Driver := "mysql"
	User := "userdb"
	Passwd := "WzGoBoot"
	Port := "23306"
	DB := "api_db"

	db, err := sql.Open(Driver, User+":"+Passwd+"@tcp(localhost:"+Port+")/"+DB)
	if err != nil {
		panic(err.Error())
	}
	return db
}
