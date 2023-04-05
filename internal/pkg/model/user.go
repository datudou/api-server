package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"index:idx_userId;unique;type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	NickName string `gorm:"type:varchar(255);not null"`
	Comment  string `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
