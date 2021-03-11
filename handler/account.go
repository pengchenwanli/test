package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test/service"
)

func (h *Handler) NewAccount(c *gin.Context) {
	var req service.NewAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AccountService.NewAccount(c, &req)
	if err != nil {
		log.Printf("[E] NewAccount:#{err}")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, rep)
}

func (h *Handler) GetAccounts(c *gin.Context) {
	var req service.GetAccountsReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AccountService.GetAccounts(c, &req)
	if err != nil {
		log.Printf("[E] Getaccounts:#{err}")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, rep)
}
func (h *Handler) Charge(c *gin.Context) {
	var req service.ChargeAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AccountService.ChargeAccount(c, &req)
	if err != nil {
		log.Printf("[E]charge:#{err}")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, rep)
}
func (h *Handler) CalcAccount(c *gin.Context) {
	var req service.CalcAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AccountService.CalcAccount(c, &req)
	if err != nil {
		log.Printf("[E]CalAccount:#{err}")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, rep)
}
func (h *Handler) DeleteAccount(c *gin.Context) {
	var req service.DeleteAccountReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = h.AccountService.DeleteAccount(c, &req)
	if err != nil {
		log.Printf("[E]CalAccount:#{err}")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, nil)
}
