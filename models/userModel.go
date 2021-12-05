package Models

import (
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
}

func NewUser(username string, password string) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return &User{}, err
	}

	return &User{
		Username: username,
		Password: string(bytes),
	}, nil
}
