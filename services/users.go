package services

import (
	"github.com/KatherineEbel/bookstore_users-api/domain/users"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
)

func Insert(u *users.User) (*users.User, *errors.RestError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.Save(); err != nil {
		return nil, err
	}
	return u, nil
}

func Get(id int64) (*users.User, *errors.RestError) {
	u := users.User{Id: id}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}
