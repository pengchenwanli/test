package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	Id       int64          `json:"id" gorm:"primaryKey"`
	CreatAt  time.Time      `json:"created_at"`
	UpdateAt time.Time      `json:"update_at"`
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
	IdCard   string         `json:"id_card" gorm:"type:varchar(255);uniqueIndex" `
	Name     string         `json:"name"`
	Password string         `json:"-"`
	Balance  float64        `json:"balance"`
}
