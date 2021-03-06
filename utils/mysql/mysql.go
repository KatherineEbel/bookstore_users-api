package mysql

import (
	"strings"

	"github.com/KatherineEbel/bookstore-utils-go/rest/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	// emailIndex = "users_email_uindex"
	notFound = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), notFound) {
			return errors.NewNotFoundError("no matching record found")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("value taken")
	default:
		return errors.NewInternalServerError("error processing request")
	}
}
