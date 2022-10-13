package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255)" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
}
