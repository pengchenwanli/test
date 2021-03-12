package services

import (
	"context"
	"errors"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"test/model"
	"test/pkg/bcrypt"
	"test/service"
)

type adminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) service.AdminService {
	return &adminService{
		db: db,
	}
}

func (s *adminService) NewAdmin(ctx context.Context, req *service.NewAdminReq) (*service.NewAdminRep, error) {
	password, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	admin := &model.Admin{
		Account:  req.Account,
		Password: password,
	}
	err = s.db.Create(admin).Error
	if err != nil {
		return nil, err
	}
	return &service.NewAdminRep{Admin: admin}, err
}

func getAdminById(db *gorm.DB, id int64) (*model.Admin, error) {
	var admin model.Admin
	err := db.Where("id", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func getAdminByAccount(db *gorm.DB, account string) (*model.Admin, error) {
	var admin model.Admin
	err := db.Where("account", account).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func getTokenByAccount(db *gorm.DB, account string) (*model.Token, error) {
	var token model.Token
	err := db.
		Where("admin_id = (SELECT id FROM admin WHERE account = ?)", account).
		First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *adminService) LoginAdmin(ctx context.Context, req *service.LoginAdminReq) (*service.LoginAdminRep, error) {
	// find if already had a token
	token, err := getTokenByAccount(s.db, req.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if token != nil {
		if token.IsExpired() {
			err := s.db.Where("id", token.Id).Delete(&model.Token{}).Error
			if err != nil {
				return nil, err
			}
		} else {
			return &service.LoginAdminRep{Token: token}, nil
		}
	}

	// gen a token

	// query from database
	admin, err := getAdminByAccount(s.db, req.Account)
	if err != nil {
		return nil, err
	}

	// validate password
	if !bcrypt.ComparePassword(admin.Password, req.Password) {
		// 1. 密码错误 concurrent password
		// 2. 账号或密码错误 invalid account or password
		return nil, errors.New("concurrent password")
	}

	// generate token and response
	token = &model.Token{
		AdminId:     admin.Id,
		AccessToken: xid.New().String(),
	}
	err = s.db.Create(token).Error
	if err != nil {
		return nil, err
	}

	return &service.LoginAdminRep{Token: token}, nil
}

func (s *adminService) LogoutAdmin(ctx context.Context) error {
	// read admin information from context
	c := GetContext(ctx)
	// delete token of this admin
	err := s.db.Delete(&model.Token{}, "admin_id = ?", c.Token.AdminId).Error
	return err
}

func getTokenByAccessToken(db *gorm.DB, accessToken string) (*model.Token, error) {
	var token model.Token
	err := db.Where("access_token", accessToken).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

var ErrInvalidToken = errors.New("invalid token")

func invalidTokenErr(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInvalidToken
	}
	return err
}

func (s *adminService) SessionVerify(ctx context.Context, req *service.SessionVerifyReq) error {
	// TODO: parse token
	token, err := getTokenByAccessToken(s.db, req.AccessToken)
	if err != nil {
		return invalidTokenErr(err)
	}

	if token.IsExpired() {
		err := s.db.Where("id", token.Id).Delete(&model.Token{}).Error
		if err != nil {
			return err
		}
		return ErrInvalidToken
	}

	// validate token & get admin information
	admin, err := getAdminById(s.db, token.AdminId)
	if err != nil {
		return err
	}

	ctx = WithContext(ctx, &Context{
		Token: token,
		Admin: admin,
	})

	return nil
}

/*package services

import (
	"context"
	"errors"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"test/model"
	"test/pkg/bcrypt"
	"test/service"
)

type adminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) service.AdminService {
	return &adminService{
		db: db,
	}
}

func (s *adminService) NewAdmin(ctx context.Context, req *service.NewAdminReq) (*service.NewAdminRep, error) {
	hashPasswd, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	account := &model.Admin{
		Account:  req.Account,
		Password: string(hashPasswd),
	}
	err = s.db.Create(account).Error
	if err != nil {
		return nil, err
	}
	return &service.NewAdminRep{Admin: account}, err
}

func getAccountById(db *gorm.DB, id int64) (*model.Admin, error) {
	var admin model.Admin
	err := db.Where("id", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
func getAdminByAccount(db *gorm.DB, account string) (*model.Admin, error) {
	var admin model.Admin
	err := db.Where("account", account).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
func getTokenByAccount(db *gorm.DB, account string) (*model.Token, error) {
	var token model.Token
	err := db.Where("admin_id=(SELECT id FROM admin WHERE account = ?)", account).
		First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}
func (s *adminService) LoginAdmin(ctx context.Context, req *service.LoginAdminReq) (rep *service.LoginAdminRep, err error) {

	token, err := getTokenByAccount(s.db, req.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { //error.ISIs 是严格判断相等，即两个 error 是否相等。//error.AS则是判断类型是否相同，并提取第一个符合目标类型的错误，用来统一处理某一类错误。
		return nil, err
	}
	if token != nil {
		if token.IsExpired() {
			err = s.db.Where("id", token.Id).Delete(&model.Token{}).Error
			if err != nil {
				return nil, err
			}
		} else {
			return &service.LoginAdminRep{Token: token}, nil
		}

	}
	//产生token

	//查询用户
	admin, err := getAdminByAccount(s.db, req.Account)
	if err != nil {
		return nil, err
	}
	//对比密码
	if !bcrypt.ComparePassword(admin.Password, req.Password) {
		return nil, errors.New("concurrent Password")
	}

	token = &model.Token{
		AdminId:     admin.Id,
		AccessToken: xid.New().String(),
	}
	err = s.db.Create(token).Error
	if err != nil {
		return nil, err
	}

	return &service.LoginAdminRep{Token: token}, nil
}
func (s *adminService) LogoutAdmin(ctx context.Context) error {
	c := GetContext(ctx)
	err := s.db.Where("admin_id=?", c.Admin.Id).Delete(&model.Token{}).Error
	return err
}
func getTokenByAccessToken(db *gorm.DB, accessToken string) (*model.Token, error) {
	var token model.Token
	err := db.Where("access_token", accessToken).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

var ErrInvalidToken = errors.New("invalid token")

func invalidTokenErr(err error) error {
	if errors.Is(err, ErrInvalidToken) {
		return ErrInvalidToken
	}
	return err
}

func (s *adminService) SessionVerify(ctx context.Context, req *service.SessionVerifyReq) (err error) {
	//解析Token
	token, err := getTokenByAccessToken(s.db, req.AccessToken)
	if err != nil {
		return invalidTokenErr(err)
	}
	if token.IsExpired() {
		err = s.db.Where("id", token.Id).Delete(&model.Token{}).Error
		if err != nil {
			return err
		}
		return ErrInvalidToken
	}
	//1.验证token  2.获取管理员得信息
	admin, err := getAccountById(s.db, token.AdminId)
	if err != nil {
		return err
	}
	ctx = WithContext(ctx, &Context{
		Token: token,
		Admin: admin,
	})
	return nil
}
*/
