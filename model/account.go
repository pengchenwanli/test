package model

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	Id        int64          `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
	IdCard    string         `json:"id_card" gorm:"type:varchar(255);uniqueIndex" `
	Name      string         `json:"name"`
	Password  string         `json:"-"`
	Balance   float64        `json:"balance"`
}
