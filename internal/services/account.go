package services

import (
	"context"
	"gorm.io/gorm"
	"log"
	"test/model"
	"test/pkg/bcrypt"
	"test/service"
)

type accountService struct {
	price float64
	db    *gorm.DB
}

func NewAccountService(price float64, db *gorm.DB) service.AccountService {
	return &accountService{
		price: price,
		db:    db,
	}
}

func (s *accountService) GetAccountByIdCard(idCard string) (*model.Account, error) {
	account := &model.Account{}
	err := s.db.Where("id_card", idCard).First(account).Error
	return account, err
}
func (s *accountService) NewAccount(ctx context.Context, req *service.NewAccountReq) (*service.NewAccountRep, error) {
	hashPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	account := &model.Account{
		IdCard:   req.IdCard,
		Name:     req.Name,
		Password: string(hashPassword),
	}
	err = s.db.Create(account).Error
	if err != nil {
		log.Printf("[D] service.NewAccount(%v): %v", req, err)
		return nil, err
	}
	return &service.NewAccountRep{Account: account}, nil
}
func (s *accountService) GetAccounts(ctx context.Context, req *service.GetAccountsReq) (*service.GetAccountsRep, error) {
	rep := new(service.GetAccountsRep)
	db := s.db.Model(model.Account{}). //选择指定模型,指定表名
						Where(&model.Account{
			Id:     req.Id,
			Name:   req.Name,
			IdCard: req.IdCard,
		})
	err := db.Count(&rep.Total).Error //统计符合条件的记录个数
	if err != nil {
		return nil, err
	}
	err = db.Offset(req.Offset).
		Limit(req.Limit).
		Order("id DESC").
		Find(rep.Accounts).Error //db.Find 获取全部记录
	if err != nil {
		return nil, err
	}
	return rep, nil
}

func (s *accountService) ChargeAccount(ctx context.Context, req *service.ChargeAccountReq) (*service.ChargeAccountRep, error) {
	err := s.db.
		Model(model.Account{}).
		Where("id_card", req.IdCard).
		Update("balance", gorm.Expr("balance+?", req.Amount)).Error
	if err != nil {
		return nil, err
	}
	assetLog := &model.Asset{
		IdCard: req.IdCard,
		Type:   model.AssetTypeCharge,
		Amount: float64(req.Amount),
	}
	err = s.db.Create(assetLog).Error //gorm创建记录利用db.Create,创建数据行
	if err != nil {
		return nil, err
	}
	account, err := s.GetAccountByIdCard(req.IdCard)
	if err != nil {
		return nil, err
	}

	return &service.ChargeAccountRep{Account: account}, err
}
func (s *accountService) CalcAccount(ctx context.Context, req *service.CalcAccountReq) (*service.CalcAccountRep, error) {
	amount := s.price * req.Hour
	err := s.db.
		Model(model.Account{}).
		Where("id_card", req.IdCard).
		Update("balance", gorm.Expr("balance+?", -amount)).Error
	if err != nil {
		return nil, err
	}
	assetLog := &model.Asset{
		IdCard: req.IdCard,
		Type:   model.AssetTypeConsume,
		Amount: amount,
	}
	err = s.db.Create(assetLog).Error
	if err != nil {
		return nil, err
	}
	account, err := s.GetAccountByIdCard(req.IdCard)
	if err != nil {
		return nil, err
	}
	return &service.CalcAccountRep{Account: account}, nil
}
func (s *accountService) DeleteAccount(ctx context.Context, req *service.DeleteAccountReq) error {
	err := s.db.
		Where("id_card", req.IdCard).
		Delete(&model.Account{}).Error
	if err != nil {
		return err
	}
	return nil
}
