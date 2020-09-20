package users

import (
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
)

var (
	db = make(map[int64]*User)
)

// Only place that should have access to the database is the DAO

func (u *User) Get() *errors.RestError {
	r := db[u.Id]
	if r == nil {
		return errors.NewNotFoundError("user not found")
	}
	u.FirstName = r.FirstName
	u.LastName = r.LastName
	u.Email = r.Email
	u.CreatedAt = r.CreatedAt
	return nil
}

func (u *User) Save() *errors.RestError {
	r := db[u.Id]
	if r != nil {
		if r.Email == u.Email {
			return errors.NewBadRequestError("email in use")
		}
		return errors.NewBadRequestError("invalid request")

	}
	u.CreatedAt = dates.GetNowString()
	db[u.Id] = u
	return nil
}
