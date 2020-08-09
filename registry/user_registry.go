package registry

import (
	"github.com/jonathanlazaro1/go-clean-arch/interface/controller"
	ip "github.com/jonathanlazaro1/go-clean-arch/interface/presenter"
	ir "github.com/jonathanlazaro1/go-clean-arch/interface/repository"
	"github.com/jonathanlazaro1/go-clean-arch/usecase/interactor"
	up "github.com/jonathanlazaro1/go-clean-arch/usecase/presenter"
	ur "github.com/jonathanlazaro1/go-clean-arch/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
