package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-" gorm:"not null"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	IsActive  bool   `json:"is_active" gorm:"default:false"`
}
