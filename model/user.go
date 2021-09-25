package model

import "encoding/json"

type User struct {
	UserId  int    `json:"userId"`
	Name  	string `json:"name"`
	Email 	string `json:"email"`
	Phone 	string `json:"phone"`
}

func (u *User) IsValid() bool {
	if u.UserId == 0 {
		return false
	}
	if u.Email == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.Phone == "" {
		return false
	}
	return true
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
