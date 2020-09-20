package users

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/KatherineEbel/bookstore_users-api/dataSources/mysql/usersDb"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
	"github.com/KatherineEbel/bookstore_users-api/utils/mysql"
)

const (
	queryInsertUser = "INSERT INTO users(firstName, lastName, email) VALUES(?,?,?)"
	queryGetUser    = "SELECT id, firstName, lastname, email FROM users WHERE id=?"
)

// Only place that should have access to the database is the DAO

func init() {
	if err := usersDb.Client.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Ping Success!!")
	}
}

func (u *User) Get() *errors.RestError {
	stmt, err := prepareStatement(queryGetUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.Id)
	if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email); err != nil {
		return mysql.ParseError(err)
	}
	return nil
}

func (u *User) Save() *errors.RestError {
	stmt, err := prepareStatement(queryInsertUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, execErr := stmt.Exec(u.FirstName, u.LastName, u.Email)
	if execErr != nil {
		return mysql.ParseError(execErr)
	}
	uId, sErr := result.LastInsertId()
	if sErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error saving user %s", sErr.Error()))
	}
	u.Id = uId
	return nil
}

func prepareStatement(query string) (*sql.Stmt, *errors.RestError) {
	stmt, err := usersDb.Client.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return stmt, nil
}
