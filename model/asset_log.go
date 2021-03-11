package model

import "time"

const (
	AssetTypeCharge  = 1
	AssetTypeConsume = 2
	AssetTypeReward  = 3
)

type Asset struct {
	Id       int64     `json:"id" gorm:"primaryKey"`
	IdCard   string    `json:"id_card"`
	Type     int64     `json:"type"`
	Amount   float64   `json:"amount"`
	CreateAt time.Time `json:"created_at"`
}
