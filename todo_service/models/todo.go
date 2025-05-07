package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	UserID     uint   `json:"user_id"`
	ProjectID  uint   `json:"project_id"`
	CategoryID uint   `json:"category_id"`
}
