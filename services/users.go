package services

import (
	"github.com/KatherineEbel/bookstore_users-api/domain/users"
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
)

func Insert(u *users.User) (*users.User, *errors.RestError) {
	u.Status = users.StatusActive
	u.DateCreated = dates.GetNowString(dates.APIDateLayout)
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

func Update(partial bool, user *users.User) (*users.User, *errors.RestError) {
	cur, err := Get(user.Id)
	if err != nil {
		return nil, err
	}
	if partial {
		if user.FirstName != "" {
			cur.FirstName = user.FirstName
		}
		if user.LastName != "" {
			cur.LastName = user.LastName
		}
		if user.Email != "" {
			cur.Email = user.Email
		}
	} else {
		// if err := user.Validate(); err != nil {
		// 	return nil, err
		// }
		cur.FirstName = user.FirstName
		cur.LastName = user.LastName
		cur.Email = user.Email
	}
	if err := cur.Update(); err != nil {
		return nil, err
	}
	return cur, nil
}

func Delete(id int64) *errors.RestError {
	u, err := Get(id)
	if err != nil {
		return err
	}
	if err := u.Delete(); err != nil {
		return err
	}
	return nil
}

func FindByStatus(status string) ([]*users.User, *errors.RestError) {
	results, err := users.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return results, nil
}
