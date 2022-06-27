package infrastructure

import (
	"day-15/app/dto"
	"day-15/app/interfaces"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Dispatch(sqlHandler *gorm.DB) {
	e := echo.New()
	LogMiddlewares(e)
	e.Validator = &CustomValidator{validator: validator.New()}
	
	auth := e.Group("/auth")
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExp := time.Hour * 1
	jwtSigningMethod := jwt.SigningMethodHS256
	jwtAuthController := interfaces.NewJWTAuthController(sqlHandler, jwtSecret, jwtExp, jwtSigningMethod)
	auth.POST("/login", jwtAuthController.LoginWithEmailAndPassword)

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status":"OK"})
	})

	v1 := e.Group("/v1")

	userController := interfaces.NewUserController(sqlHandler)
	
	
	
	jwtConfig := middleware.JWTConfig{
		Claims: &dto.JWTClaims{},
		SigningKey: []byte(jwtSecret),
		Skipper: func(c echo.Context) bool {
			var (
				path = c.Request().URL.Path
				method = c.Request().Method
				usersPath = regexp.MustCompile(`\/v1\/users`)
				booksPath = regexp.MustCompile(`\/v1\/books`)
			)
			if (usersPath.MatchString(path)) && (method == "POST") {
				return true
			} else if (booksPath.MatchString(path)) && (method == "GET") {
				return true
			}
			return false
		},
	}
	
	v1Users := v1.Group("/users")
	v1Users.Use(middleware.JWTWithConfig(jwtConfig))
	v1Users.GET("", userController.Index)
	v1Users.POST("", userController.Store)
	v1Users.GET("/:id", userController.Show)
	v1Users.PUT("/:id", userController.Edit)
	v1Users.DELETE("/:id", userController.Destroy)

	bookController := interfaces.NewBookController(sqlHandler)
	
	v1Books := v1.Group("/books")
	v1Books.Use(middleware.JWTWithConfig(jwtConfig))
	v1Books.GET("", bookController.Index)
	v1Books.POST("", bookController.Store)
	v1Books.GET("/:id", bookController.Show)
	v1Books.PUT("/:id", bookController.Edit)
	v1Books.DELETE("/:id", bookController.Destroy)

	e.Logger.Fatal(e.Start(":8080"))
}
