package models

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	UserName string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}