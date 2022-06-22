package interfaces

import (
	"day-12/app/domain"
	"day-12/app/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		UserInteractor: usecases.UserInteractor{
			UserRepository: &UserRepository{db},
		},
	}
}

func (uc *UserController) Index(c echo.Context) error {
	users, err := uc.UserInteractor.Index()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": users})
}

func (uc *UserController) Store(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	_, err := uc.UserInteractor.Store(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}
