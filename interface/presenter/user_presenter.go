package presenter

import (
	"fmt"

	"github.com/jonathanlazaro1/go-clean-arch/domain/model"
)

type userPresenter struct{}

// UserPresenter has contracts about to present Users to the visualization layers
type UserPresenter interface {
	ResponseUsers(us []model.User) []model.User
}

// NewUserPresenter creates a new instance of userPresenter
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []model.User) []model.User {
	for _, u := range us {
		u.Name = fmt.Sprint("Mr.", u.Name)
	}
	return us
}
