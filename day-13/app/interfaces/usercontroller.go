package interfaces

import (
	"day-13/app/domain"
	"day-13/app/usecases"
	"net/http"
	"strconv"

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

func (uc *UserController) Show(c echo.Context) error {
	var (
		err    error
		userID int
	)
	userID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := uc.UserInteractor.Show(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (uc *UserController) Edit(c echo.Context) error {
	var (
		err    error
		userID int
		data   domain.User
	)
	
	userID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err = c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := uc.UserInteractor.Edit(uint(userID), data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}

func (uc *UserController) Destroy(c echo.Context) error {
	var (
		err    error
		userID int
	)
	
	userID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := uc.UserInteractor.Destroy(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": user})
}
