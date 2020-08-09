package controller

import (
	"net/http"

	"github.com/jonathanlazaro1/go-clean-arch/usecase/interactor"
	"github.com/labstack/echo"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

// UserController has contracts about how to handle proccess API requests about user model
type UserController interface {
	GetUsers(c echo.Context) error
}

// NewUserController creates a new instance of userController
func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c echo.Context) error {
	u, err := uc.userInteractor.Get()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
