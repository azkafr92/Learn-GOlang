package interfaces

import (
	"day-15/app/domain"
	"day-15/app/usecases"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	BookInteractor usecases.BookInteractor
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		BookInteractor: usecases.BookInteractor{
			BookRepository: &BookRepository{db},
		},
	}
}

func (uc *BookController) Index(c echo.Context) error {
	books, err := uc.BookInteractor.Index()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": books})
}

func (uc *BookController) Store(c echo.Context) error {
	var book domain.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	_, err := uc.BookInteractor.Store(&book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": book})
}

func (uc *BookController) Show(c echo.Context) error {
	var (
		err    error
		bookID int
	)
	bookID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	book, err := uc.BookInteractor.Show(uint(bookID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": book})
}

func (uc *BookController) Edit(c echo.Context) error {
	var (
		err    error
		bookID int
		data   domain.Book
	)
	
	bookID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err = c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	book, err := uc.BookInteractor.Edit(uint(bookID), data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": book})
}

func (uc *BookController) Destroy(c echo.Context) error {
	var (
		err    error
		bookID int
	)
	
	bookID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	book, err := uc.BookInteractor.Destroy(uint(bookID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": book})
}
