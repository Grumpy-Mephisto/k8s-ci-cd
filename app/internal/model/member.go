package model

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
}
