package main

import (
	"fmt"

	"github.com/KatherineEbel/bookstore_users-api/app"
	"github.com/KatherineEbel/bookstore_users-api/dataSources/mysql/usersDb"
)

func main() {
	app.StartApplication()
	fmt.Println(usersDb.UsersDB)
}
