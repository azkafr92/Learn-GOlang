package controller

import (
	"day11/config"
	"day11/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func returnBadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": message,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnBadRequest(c, "invalid user id")
	}
	
	// if len(users) < id {
	// 	return returnBadRequest(c, "user id not found")
	// }
	// user := users[id-1]

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return returnBadRequest(c, "user not found")
	}

	return c.JSON(
		http.StatusOK, map[string]interface{}{
			"message": "success get user",
			"user":    user,
		},
	)
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnBadRequest(c, "invalid user id")
	}
	// if len(users) < id {
	// 	return returnBadRequest(c, "user id not found")
	// }
	// users[id-1] = User{}
	
	if err := config.DB.Model(&models.User{}).Delete("id = ?", id).Error; err != nil {
		return returnBadRequest(c, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return returnBadRequest(c, "invalid user id")
	}

	// if len(users) < id {
	// 	return returnBadRequest(c, "user id not found")
	// }

	var update models.User
	c.Bind(&update)

	// user := users[id-1]

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return returnBadRequest(c, "user not found")
	}

	if update.Email != "" {
		user.Email = update.Email
		// users[id-1].Email = update.Email
	}
	if update.Name != "" {
		user.Name = update.Name
		// users[id-1].Name = update.Name
	}
	if update.Password != "" {
		user.Password = update.Password
		// users[id-1].Password = update.Password
	}

	config.DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func CreateUsersController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	if err:= config.DB.Save(&user).Error; err != nil {
		return returnBadRequest(c, err.Error())
	}

	// user.ID = len(users) + 1
	// users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}
