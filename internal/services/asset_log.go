package services

import (
	"context"
	"test/model"
	"test/service"

	"gorm.io/gorm"
)

type assetLogService struct {
	db *gorm.DB
}

func NewAssetLogService(db *gorm.DB) service.AssetLogService {
	return &assetLogService{
		db: db,
	}
}

func (s *assetLogService) GetAssetLogs(ctx context.Context, req *service.AssetLogReq) (*service.AssetLogRep, error) {
	db := s.db.
		Model(model.Asset{}).
		Where(&model.Asset{
			Id:     req.Id,
			IdCard: req.IdCard,
			Type:   req.Type,
		})
	if !req.SinceTime.IsZero() {
		db = db.Where("created_at >= ?", req.SinceTime)
	}
	if !req.UntilTime.IsZero() {
		db = db.Where("created_at <= ?", req.UntilTime)
	}

	rep := new(service.AssetLogRep)

	err := db.Count(&rep.Total).Error
	if err != nil {
		return nil, err
	}

	// SELECT * FROM asset_log WHERE id = ? AND id_cart = ? AND type = ? AND created_at >= ? LIMIT offset, limit;
	err = db.Offset(req.Offset).Limit(req.Limit).Order("id desc").Find(&rep.Asset).Error
	if err != nil {
		return nil, err
	}

	return rep, nil
}
