package presenter

import (
	"github.com/jonathanlazaro1/go-clean-arch/domain/model"
)

// UserPresenter has contracts about how to present Users to the outer layers
type UserPresenter interface {
	ResponseUsers(u []model.User) []model.User
}
