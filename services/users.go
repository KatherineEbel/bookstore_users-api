package services

import (
	"github.com/KatherineEbel/bookstore_users-api/domain/users"
	"github.com/KatherineEbel/bookstore_users-api/utils/crypt"
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
	"github.com/KatherineEbel/bookstore_users-api/utils/errors"
)

type usersService struct{}

var (
	UsersService IUsersService = &usersService{}
)

type IUsersService interface {
	Insert(*users.User) (*users.User, *errors.RestError)
	Get(int64) (*users.User, *errors.RestError)
	Update(bool, *users.User) (*users.User, *errors.RestError)
	Delete(int64) *errors.RestError
	FindByStatus(string) (users.Users, *errors.RestError)
}

func (s *usersService) Insert(u *users.User) (*users.User, *errors.RestError) {
	u.Status = users.StatusActive
	u.DateCreated = dates.GetNowString(dates.APIDateLayout)
	hp, err := crypt.Encrypt(u.Password)
	if err != nil {
		return nil, errors.NewBadRequestError("Invalid password")
	}
	u.Password = hp
	if err := u.Save(); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *usersService) Get(id int64) (*users.User, *errors.RestError) {
	u := users.User{Id: id}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *usersService) Update(partial bool, user *users.User) (*users.User, *errors.RestError) {
	cur, err := s.Get(user.Id)
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
		cur.FirstName = user.FirstName
		cur.LastName = user.LastName
		cur.Email = user.Email
	}
	if err := cur.Update(); err != nil {
		return nil, err
	}
	return cur, nil
}

func (s *usersService) Delete(id int64) *errors.RestError {
	u, err := s.Get(id)
	if err != nil {
		return err
	}
	if err := u.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) FindByStatus(status string) (users.Users, *errors.RestError) {
	results, err := users.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return results, nil
}
