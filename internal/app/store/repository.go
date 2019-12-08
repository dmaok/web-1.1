package store

import "github.com/dmaok/web-1.1/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
