package interfaces

import (
	"day-14/app/dto"
	"day-14/app/usecases"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type JWTAuthController struct {
	UserInteractor usecases.UserInteractor
	secret         string
	exp            time.Duration
	signingMethod  jwt.SigningMethod
}

func NewJWTAuthController(db *gorm.DB, secret string, exp time.Duration, signingMethod jwt.SigningMethod) *JWTAuthController {
	return &JWTAuthController{
		UserInteractor: usecases.UserInteractor{UserRepository: &UserRepository{db}}, secret: secret, exp: exp, signingMethod: signingMethod}
}

func (jc *JWTAuthController) createJWTToken(userID uint, email string) (string, error) {
	claims := dto.JWTClaims{
		Authorized: true,
		UserID:     userID,
		Email:      email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jc.exp).Unix(),
		},
	}
	token := jwt.NewWithClaims(jc.signingMethod, claims)
	return token.SignedString([]byte(jc.secret))
}

func (jc *JWTAuthController) LoginWithEmailAndPassword(c echo.Context) error {
	var data dto.LoginWithEmailAndPasswordRequestBody

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := c.Validate(data); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := jc.UserInteractor.ShowByEmailAndPassword(data.Email, data.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Wrong email or password, or user may have been deleted."})
	}

	token, err := jc.createJWTToken(uint(user.ID), user.Email)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, echo.Map{"data": token})
}
