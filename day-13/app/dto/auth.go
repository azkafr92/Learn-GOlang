package dto

import "github.com/golang-jwt/jwt"

type LoginWithEmailAndPasswordRequestBody struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type JWTClaims struct {
	Authorized     bool   `json:"authorized"`
	UserID         uint   `json:"user_id"`
	Email          string `json:"email"`
	jwt.StandardClaims
}
