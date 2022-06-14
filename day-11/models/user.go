package models

type User struct {
	Name     string `json:"name" gorm:"varchar(255);not null" form:"name"`
	Email    string `json:"email" gorm:"varchar(125);not null;unique" form:"email"`
	Password string `json:"password" gorm:"varchar(125);not null" form:"password"`
	Common
}
