package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email:    "test@mail.com",
		Password: "password",
	}
}
