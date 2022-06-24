package domain

type Book struct {
	Title    string `json:"title" gorm:"varchar(125);not null;" form:"title"`
	Common
}
