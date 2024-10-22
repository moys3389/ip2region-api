package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *SearchHandler) HandleSearchPage(ctx *gin.Context) {
	ip := ctx.Query("ip")
	if ip == "" {
		ip = ctx.ClientIP()
	}

	if ip == "::1" {
		ip = "127.0.0.1"
	}

	result, err := h.searchService.Search(ip)
	if err != nil {
		ResponseFail(ctx, err)
		return
	}
	ctx.String(http.StatusOK, "当前IP:"+ip+" 来自于:"+result)
}
