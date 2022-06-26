package domain

type User struct {
	Email    string `json:"email" gorm:"varchar(125);not null;unique" form:"email"`
	Password string `json:"password" gorm:"varchar(125);not null" form:"password"`
	Common
}
