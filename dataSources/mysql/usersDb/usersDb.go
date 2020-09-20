package usersDb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	UsersDB   *sql.DB
	mySQLPass = os.Getenv("MYSQL_PASS")
	dbName    = os.Getenv("DB_NAME")
	mySQLUser = os.Getenv("MYSQL_USER")
	host      = os.Getenv("HOST")
)

func init() {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mySQLUser, mySQLPass, host, dbName)
	UsersDB, err = sql.Open("mysql", url)
	if err != nil {
		log.Fatalln(err)
	}
	UsersDB.SetMaxOpenConns(20)
	fmt.Println("mysql connected...")
}
