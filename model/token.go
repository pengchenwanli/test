package model

import "time"

type Token struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	AdminId     int64     `json:"amin_id" gorm:"index"`
	AccessToken string    `json:"access_token" gorm:"index"`
	CreateAt    time.Time `json:"create_at" `
	UpdateAt    time.Time `json:"update_at"`
}

func (t *Token) IsExpired() bool {
	return t.CreateAt.Before(time.Now().Add(-time.Hour))
}
