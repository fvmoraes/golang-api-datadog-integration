package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Job       string         `json:"job"`
	Age       uint           `json:"age"`
	Photo     string         `json:"photo"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
