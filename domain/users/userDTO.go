package users

import (
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Password    string `json:"-"`
	DateCreated string
	dateUpdated string
}

func (u *User) Joined() string {
	return dates.GetDateFromString(u.DateCreated)
}
