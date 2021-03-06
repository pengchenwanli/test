package service

import (
	"context"
	"test/model"
)

type AdminService interface {
	NewAdmin(ctx context.Context, req *NewAdminReq) (*NewAdminRep, error)
	LoginAdmin(ctx context.Context, req *LoginAdminReq) (*LoginAdminRep, error)
	LogoutAdmin(ctx context.Context) error
	SessionVerify(ctx context.Context, req *SessionVerifyReq) error //验证器

}
type NewAdminReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type NewAdminRep struct {
	Admin *model.Admin `json:"admin"`
}
type LoginAdminReq struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type LoginAdminRep struct {
	Token *model.Token `json:"token"`
}
type SessionVerifyReq struct {
	AccessToken string `json:"access_token"`
}
