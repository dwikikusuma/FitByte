package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	UserID   uint   `gorm:"column:user_id"`
	FileName string `gorm:"column:file_name"`
	FileURL  string `gorm:"column:file_url"`
}
