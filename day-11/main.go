package main

import (
	"day11/config"
	"day11/controller"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// var users []User

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": "OK"})
	})
	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)

	e.Logger.Fatal(e.Start(":8080"))
}
