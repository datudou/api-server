package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null"`
	MakingTime  string `gorm:"type:varchar(100);not null"`
	Serves      string `gorm:"type:varchar(100);not null"`
	Ingredients string `gorm:"type:varchar(300);not null"`
	Cost        int    `gorm:"not null"`
}

func (Recipe) TableName() string {
	return "recipes"
}
