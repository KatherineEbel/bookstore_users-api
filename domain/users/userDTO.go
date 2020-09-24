package users

import (
	"errors"

	"github.com/KatherineEbel/bookstore_users-api/utils/crypt"
	"github.com/KatherineEbel/bookstore_users-api/utils/dates"
)

const (
	StatusActive = "active"
)

type JoiningUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	PasswordHash string `json:"-"`
	DateCreated  string `json:"date_created"`
	dateUpdated  string
}

type Users []User

func NewUser(from *JoiningUser) (*User, error) {
	hp, err := crypt.Encrypt(from.Password)
	if err != nil {
		return nil, errors.New("error processing password")
	}
	return &User{
		FirstName:    from.FirstName,
		LastName:     from.LastName,
		Email:        from.Email,
		Status:       StatusActive,
		DateCreated:  dates.GetNowString(dates.APIDateLayout),
		PasswordHash: hp,
	}, nil
}
func (u *User) Joined() string {
	return dates.GetDateFromString(u.DateCreated)
}
