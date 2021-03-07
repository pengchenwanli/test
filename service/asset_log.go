package service

import (
	"context"
	"test/model"
	"time"
)

type AssetLogService interface {
	GetAssetLogs(ctx context.Context, req *AssetLogReq) (*AssetLogRep, error)
}
type AssetLogReq struct {
	Pagination
	Id        int64     `json:"id" form:"id"`
	IdCard    string    `json:"id_card" form:"id_card"`
	Type      int64     `json:"type" form:"type"`
	SinceTime time.Time `json:"since_time" form:"id_card"`
	UntilTime time.Time `json:"until_time" form:"id_card"`
}
type AssetLogRep struct {
	Total int64
	Asset []*model.Asset
}
