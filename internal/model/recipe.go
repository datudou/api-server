package model

import (
	"time"
)

type Recipe struct {
	ID          uint       `json:"id" gorm:"primarykey"`
	Title       string     `json:"title" gorm:"type:varchar(100);not null"`
	MakingTime  string     `json:"making_time" gorm:"type:varchar(100);not null"`
	Serves      string     `json:"serves" gorm:"type:varchar(100);not null"`
	Ingredients string     `json:"ingredients" gorm:"type:varchar(300);not null"`
	Cost        string     `json:"cost" gorm:"not null"`
	DeletedAt   *time.Time `json:"-" gorm:"index"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
}
