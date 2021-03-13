package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"test/service"
)

func (h *Handler) NewAdmin(c *gin.Context) {
	var req service.NewAdminReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AdminService.NewAdmin(c, &req)
	if err != nil {
		log.Printf("[E] NewAdmin:#{err}")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, rep)
}

func (h *Handler) Login(c *gin.Context) {
	var req service.LoginAdminReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AdminService.LoginAdmin(c, &req)
	if err != nil {
		log.Printf("[E] NewAdmin:#{err}")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, rep)
}

func (h *Handler) Logout(c *gin.Context) {
	err := h.AdminService.LogoutAdmin(c)
	if err != nil {
		log.Printf("[E] Logout:#{err}")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func parseBearToken(authorization string) string {
	return strings.TrimPrefix(authorization, "Bearer") //strings.TrimPrefix删除开头一样的前缀，还有strings.TrimLeft()：删除之后剩余的字符若开头与前缀存在一样的字符则还会删除
}

func (h *Handler) SessionVerifier(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	tokenStr := parseBearToken(authorization)
	var req = service.SessionVerifyReq{AccessToken: tokenStr}
	err := h.AdminService.SessionVerify(c, &req)
	if err != nil {
		log.Printf("[E] NewAccount: %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	c.Next()
}

/*func (h *Handler) SessionVerifier(c *gin.Context) {
	authorization := c.GetHeader("Authorization") //获取请求头部信息
	tokenStr := parseBearToken(authorization)
	var req = service.SessionVerifyReq{AccessToken: tokenStr}
	err := h.AdminService.SessionVerify(c, &req)
	if err != nil {
		log.Printf("[E] SessionVerifier:%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort() //Abort表示终止，也就是说，执行Abort的时候会停止所有的后面的中间件函数的调用。
		return
	}
	c.Next() //作为中间件，简单理解next表示挂起，当处理完所有的中间件函数（包括本次请求）的时候才会停止，执行完一次完整的请求。
}*/
