package model

import "time"

const (
	AssetTypeCharge  = 1
	AssetTypeConsume = 2
	AssetTypeReward  = 3
)

type Asset struct {
	Id       int64 `json:"id" gorm:"primaryKey"`
	IdCard   string
	Type     int64
	Amount   float64
	CreateAt time.Time
}
