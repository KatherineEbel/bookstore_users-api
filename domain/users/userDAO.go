package users

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/KatherineEbel/bookstore_users-api/dataSources/mysql/usersDb"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
	"github.com/KatherineEbel/bookstore_users-api/utils/mysql"
)

const (
	insertQuery = "INSERT INTO users(firstName, lastName, email) VALUES(?,?,?)"
	getQuery    = "SELECT id, firstName, lastname, email, dateCreated FROM users WHERE id=?"
	updateQuery = "UPDATE users SET firstName=?, lastName=?, email=?, updatedAt=? WHERE id=?"
	deleteQuery = "DELETE FROM users WHERE id=?"
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
	stmt, err := prepareStatement(getQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.Id)
	if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.dateCreated); err != nil {
		return mysql.ParseError(err)
	}
	return nil
}

func (u *User) Save() *errors.RestError {
	stmt, err := prepareStatement(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, execErr := stmt.Exec(u.FirstName, u.LastName, u.Email)
	fmt.Println(execErr)
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

func (u *User) Update() *errors.RestError {
	stmt, err := prepareStatement(updateQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	now := time.Now().UTC()
	_, execErr := stmt.Exec(u.FirstName, u.LastName, u.Email, now, u.Id)
	if execErr != nil {
		fmt.Println(execErr)
		return mysql.ParseError(execErr)
	}
	return nil
}

func (u *User) Delete() *errors.RestError {
	stmt, err := prepareStatement(deleteQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, execErr := stmt.Exec(u.Id)
	if execErr != nil {
		fmt.Println(execErr)
		return mysql.ParseError(execErr)
	}
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
