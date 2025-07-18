package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          	uint 	`gorm:"primarykey;autoIncrement" json:"id"`
	Title       	string 	`gorm:"type:varchar(255);not null" json:"title"`
	Description 	string 	`gorm:"type:text" json:"description"`
	Completed   	bool 	`gorm:"default:false" json:"completed"`
	CreatedAt  		time.Time `json:"created_at"`
	CompletedAt 	*time.Time `json:"completed_at"`
	UpdatedAt 		*time.Time `json:"updated_at"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}