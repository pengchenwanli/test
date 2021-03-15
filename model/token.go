package model

import "time"

type Token struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	AdminId     int64     `json:"admin_id" gorm:"index"`
	AccessToken string    `json:"access_token" gorm:"index"`
	CreatedAt   time.Time `json:"created_at" `
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Token) IsExpired() bool {
	return t.CreatedAt.Before(time.Now().Add(-time.Hour))
}
