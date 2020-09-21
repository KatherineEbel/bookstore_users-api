package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string
}

func (u Users) Marshal(isPublic bool) []interface{} {
	res := make([]interface{}, len(u))
	for i, u := range u {
		res[i] = u.Marshal(isPublic)
	}
	return res
}

func (u *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          u.Id,
			Status:      u.Status,
			DateCreated: u.DateCreated,
		}
	}
	j, _ := json.Marshal(u)
	var pu PrivateUser
	_ = json.Unmarshal(j, &pu)
	return pu
}
