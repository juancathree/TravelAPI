package domain

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string   `json:"_id" bson:"_id"`
	Name     string   `validate:"required" json:"name" bson:"name"`
	Email    string   `validate:"required,email" json:"email" bson:"email"`
	Password string   `validate:"required" json:"password" bson:"password"`
	Country  string   `validate:"required" json:"country" bson:"country"`
	Travels  []string `validate:"required,dive,unique" json:"travels" bson:"travels"`
}

func (u *User) SetPassword(password *string) error {
	// Encrypt password
	passwd, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password couldn't be encrypted")
	}

	u.Password = string(passwd)
	return nil
}
