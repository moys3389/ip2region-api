package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type HandleBatchSearchReq struct {
	Ips []string `json:"ips"`
}

type BatchSearchResp struct {
	Ip        string `json:"ip"`
	Region    string `json:"region"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"is_success"`
}

func (h *SearchHandler) HandleBatchSearch(ctx *gin.Context) {
	var req HandleBatchSearchReq
	if err := ctx.ShouldBind(&req); err != nil {
		ResponseFail(ctx, err)
		return
	}
	if req.Ips == nil {
		ResponseErrorParam(ctx)
		return
	}
	ResponseSuccess(ctx, h.batchSearch(req.Ips))
}

func (h *SearchHandler) HandleBatchSearchByQuery(ctx *gin.Context) {
	ipsStr := ctx.Query("ips")
	if ipsStr == "" {
		ResponseErrorParam(ctx)
		return
	}
	ips := strings.Split(ipsStr, ",")
	ResponseSuccess(ctx, h.batchSearch(ips))
}

func (h *SearchHandler) batchSearch(ips []string) []BatchSearchResp {
	return lo.Map(ips, func(item string, _ int) BatchSearchResp {
		res, err := h.searchService.Search(item)
		if err != nil {
			return BatchSearchResp{
				Ip:        item,
				Message:   err.Error(),
				IsSuccess: false,
			}
		}
		return BatchSearchResp{
			Ip:        item,
			Region:    res,
			Message:   "success",
			IsSuccess: true,
		}
	})
}
