package users

import (
	"database/sql"
	"fmt"
	"html"

	"github.com/KatherineEbel/bookstore_users-api/dataSources/mysql/usersDb"
	"github.com/KatherineEbel/bookstore_users-api/logger"
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
	"github.com/KatherineEbel/bookstore_users-api/utils/mysql"
)

const (
	insertQuery           = "INSERT INTO users(firstName, lastName, email, dateCreated, status, password) VALUES(?,?,?,?,?,?)"
	getQuery              = "SELECT id, firstName, lastname, email, dateCreated, status FROM users WHERE id=?"
	updateQuery           = "UPDATE users SET firstName=?, lastName=?, email=?, dateUpdated=?, status=?, WHERE id=?"
	deleteQuery           = "DELETE FROM users WHERE id=?"
	findUserByStatusQuery = "SELECT id, firstName, lastName, email, dateCreated, status FROM users WHERE status=?"
)

// Only place that should have access to the database is the DAO

func init() {
	if err := usersDb.Client.Ping(); err != nil {
		logger.Error("Database Ping err", err)
		panic(err)
	} else {
		logger.Info("Ping Success")
	}
}

func (u *User) Get() *errors.RestError {
	stmt, err := prepareStatement(getQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.Id)
	if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); err != nil {
		logger.Error("GET Scan err", err)
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
	result, execErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
	if execErr != nil {
		logger.Error("Save execErr", execErr)
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
		return err
	}
	defer stmt.Close()
	now := dates.GetNowString(dates.APIDateLayout)
	_, execErr := stmt.Exec(u.FirstName, u.LastName, u.Email, now, u.Status, u.Id)
	if execErr != nil {
		logger.Error("Update exec err", execErr)
		return mysql.ParseError(execErr)
	}
	return nil
}

func (u *User) Delete() *errors.RestError {
	stmt, err := prepareStatement(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, execErr := stmt.Exec(u.Id)
	if execErr != nil {
		logger.Error("DELETE exec err", execErr)
		return mysql.ParseError(execErr)
	}
	return nil
}

func FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := prepareStatement(findUserByStatusQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, queryErr := stmt.Query(html.EscapeString(status))
	if queryErr != nil {
		logger.Error("FindByStatus Query err", queryErr)
		return nil, mysql.ParseError(queryErr)
	}
	defer rows.Close()
	res := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Scan row err", err)
			return nil, mysql.ParseError(err)
		}
		res = append(res, user)
	}
	if len(res) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return res, nil
}

func prepareStatement(query string) (*sql.Stmt, *errors.RestError) {
	stmt, err := usersDb.Client.Prepare(query)
	if err != nil {
		logger.Error("prepareStatement err", err)
		return nil, errors.NewDatabaseError()
	}
	return stmt, nil
}
