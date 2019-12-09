package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptPassword(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 32)),
	)
}
