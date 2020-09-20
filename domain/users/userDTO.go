package users

import (
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	dateCreated string
	updatedAt   string
}

func (u *User) Joined() string {
	return dates.GetDateFromString(u.dateCreated)
}
