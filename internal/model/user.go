package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"index:idx_userId;unique;type:varchar(255);not null" json:"user_id"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	NickName string `gorm:"type:varchar(255);not null" json:"nickname"`
	Comment  string `gorm:"type:varchar(255)" json:"comment"`
}

func (User) TableName() string {
	return "users"
}
