package repository

import "github.com/jonathanlazaro1/go-clean-arch/domain/model"

// UserRepository has contracts about how to retrieve and modify Users
type UserRepository interface {
	FindAll() ([]model.User, error)
}
