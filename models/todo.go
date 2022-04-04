package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        int    `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Status    string `gorm:"type:varchar(255);default:'not done';not null" json:"status" form:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
