package interactor

import (
	"github.com/jonathanlazaro1/go-clean-arch/domain/model"
	"github.com/jonathanlazaro1/go-clean-arch/usecase/presenter"
	"github.com/jonathanlazaro1/go-clean-arch/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

// UserInteractor has contracts about how to interact with User model data from and to inner layers
type UserInteractor interface {
	Get() ([]model.User, error)
}

// NewUserInteractor creates a new instance of UserInteractor
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}

}

func (us userInteractor) Get() ([]model.User, error) {
	u, err := us.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUsers(u), nil
}
