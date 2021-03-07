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
	}
	c.JSON(http.StatusOK, rep)
}

func (h *Handler) Logout(c *gin.Context) {
	err := h.AdminService.LogoutAdmin(c)
	if err != nil {
		log.Printf("[E] Logout:#{err}")
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, nil)
}

func parseBearToken(authorization string) string {
	return strings.TrimPrefix(authorization, "Bearer") //strings.TrimPrefix删除开头一样的前缀，还有strings.TrimLeft()：删除之后剩余的字符若开头与前缀存在一样的字符则还会删除
}

func (h *Handler) SessionVerifier(c *gin.Context) {
	authorization := c.GetHeader("Authorization") //获取请求头部信息
	tokenStr := parseBearToken(authorization)
	var req = service.SessionVerifyReq{AccessToken: tokenStr}
	err := h.AdminService.SessionVerify(c, &req)
	if err != nil {
		log.Printf("[E] SessionVerifier:%v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	c.Next() //作为中间件
}
