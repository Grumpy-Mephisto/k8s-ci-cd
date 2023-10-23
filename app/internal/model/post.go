package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text     string `json:"text"`
	AuthorID uint   `json:"author_id"`
	Author   Member `gorm:"foreignKey:AuthorID"`
}
