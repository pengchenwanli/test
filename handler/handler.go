package handler

import (
	"github.com/gin-gonic/gin"
	"test/service"
)

type Handler struct {
	engine *gin.Engine
	service.AccountService
	service.AdminService
	service.AssetLogService
}

func New( //accountService service.AccountService,
	adminService service.AdminService,
	//assetLogService service.AssetLogService,
) *Handler {
	e := gin.Default()

	handler := &Handler{
		engine: e,
		//AccountService:  accountService,
		AdminService: adminService,
		//AssetLogService: assetLogService,
	}
	handler.initRouter()
	return handler
}
func (h *Handler) Run(addr ...string) error {
	return h.engine.Run(addr...) //将路由器连接到http.Server并开始侦听和处理HTTP请求
}
