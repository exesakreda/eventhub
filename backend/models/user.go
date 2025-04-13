package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `gorm:"unique" json:"username"`
	Password   string `gorm:"column:password_hash" json:"-"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationCredentials struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
