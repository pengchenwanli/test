package service

import (
	"context"
	"test/model"
)

type AccountService interface {
	NewAccount(ctx context.Context, req *NewAccountReq) (*NewAccountRep, error)
	GetAccounts(ctx context.Context, req *GetAccountsReq) (*GetAccountsRep, error)
	ChargeAccount(ctx context.Context, req *ChargeAccountReq) (*ChargeAccountRep, error)
	CalcAccount(ctx context.Context, req *CalcAccountReq) (*CalcAccountRep, error)
	DeleteAccount(ctx context.Context, req *DeleteAccountReq) error
}
type NewAccountReq struct {
	IdCard   string `json:"id_card" form:"id_card" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type NewAccountRep struct {
	Account *model.Account `json:"account"`
}
type GetAccountsReq struct {
	Pagination
	Id     int64  `json:"id"  form:"id"`
	IdCard string `json:"id_card" form:"id_card"`
	Name   string `json:"name" form:"name"`
}
type GetAccountsRep struct {
	Total    int64            `json:"total"`
	Accounts []*model.Account `json:"account"`
}
type ChargeAccountReq struct {
	IdCard string `json:"id_card" form:"id_card" binding:"required"`
	Amount int    `json:"amount" form:"amount" binding:"required"`
}
type ChargeAccountRep struct {
	Account *model.Account `json:"account"`
}
type CalcAccountReq struct {
	IdCard string  `json:"id_card" form:"id_card" binding:"required"`
	Hour   float64 `json:"hour" form:"hour" binding:"required"`
}
type CalcAccountRep struct {
	Account *model.Account `json:"account"`
}
type DeleteAccountReq struct {
	IdCard string `json:"id_card" form:"id_card" binding:"required"`
}
