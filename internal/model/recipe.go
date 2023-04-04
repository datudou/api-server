package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(100);not null"`
	MakingTime  string `json:"making_time" gorm:"type:varchar(100);not null"`
	Serves      string `json:"serves" gorm:"type:varchar(100);not null"`
	Ingredients string `json:"ingredients" gorm:"type:varchar(300);not null"`
	Cost        int    `json:"cost" gorm:"not null"`
}

func (Recipe) TableName() string {
	return "recipes"
}
