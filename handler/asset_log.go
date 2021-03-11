package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"test/service"
)

func (h *Handler) GetAssetLog(c *gin.Context) {
	var req service.AssetLogReq
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rep, err := h.AssetLogService.GetAssetLogs(c, &req)
	if err != nil {
		log.Printf("[E] GetAssetlogs:#{err}")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, rep)
}
