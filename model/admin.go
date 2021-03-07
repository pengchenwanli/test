package model

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id       int64          `json:"id" gorm:"primaryKey"`
	Account  string         `json:"account" gorm:"type:varchar(255);uniqueIndex"`
	Password string         `json:"-" gorm:"type:varchar(255)"`
	CreateAt time.Time      `json:"create_at"`
	UpdateAt time.Time      `json:"update_at"`
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
