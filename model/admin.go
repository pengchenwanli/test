package model

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int64          `json:"id" gorm:"primaryKey"`
	Account   string         `json:"account" gorm:"type:varchar(255);uniqueIndex"`
	Password  string         `json:"-" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
