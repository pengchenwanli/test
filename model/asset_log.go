package model

import "time"

const (
	AssetTypeCharge  = 1
	AssetTypeConsume = 2
	AssetTypeReward  = 3
)

type Asset struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	IDCard    string    `json:"id_card"`
	Type      int64     `json:"type"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
