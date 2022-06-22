package infrastructure

import (
	"day-12/app/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Dispatch(sqlHandler *gorm.DB) {
	e := echo.New()
	
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status":"OK"})
	})

	v1 := e.Group("/v1")

	userController := interfaces.NewUserController(sqlHandler)
	v1Users := v1.Group("/users")
	v1Users.GET("", userController.Index)
	v1Users.POST("", userController.Store)
	
	e.Start(":8080")
}
