package model

import (
	"gorm.io/gorm"
)

// INSERT INTO recipes (
// 	id,
// 	title,
// 	making_time,
// 	serves,
// 	ingredients,
// 	cost,
// 	created_at,
// 	updated_at
//   )

//   CREATE TABLE IF NOT EXISTS recipes (
// 	id integer PRIMARY KEY AUTO_INCREMENT,
// 	-- name of recipe
// 	title varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
// 	-- time required to cook/bake the recipe
// 	making_time varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
// 	-- number of people the recipe will feed
// 	serves varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
// 	-- food items necessary to prepare the recipe
// 	ingredients varchar(300) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
// 	-- price of recipe
// 	cost integer NOT NULL,
// 	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at datetime on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
//   );

type Recipe struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null"`
	MakingTime  string `gorm:"type:varchar(100);not null"`
	Serves      string `gorm:"type:varchar(100);not null"`
	Ingredients string `gorm:"type:varchar(300);not null"`
	Cost        int    `gorm:"not null"`
}
