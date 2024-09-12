package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moys3389/ip2region-api/service"
	"github.com/samber/do/v2"
)

type SearchHandler struct {
	searchService *service.SearchService
}

type HandleSearchReq struct {
	Ip string `json:"ip"`
}

func (h *SearchHandler) HandleSearch(ctx *gin.Context) {
	var req HandleSearchReq
	if err := ctx.ShouldBind(&req); err != nil {
		ResponseFail(ctx, err)
		return
	}

	result, err := h.searchService.Search(req.Ip)
	if err != nil {
		ResponseFail(ctx, err)
		return
	}

	ResponseSuccess(ctx, gin.H{
		"ip":     req.Ip,
		"region": result,
	})
}

func (h *SearchHandler) HandleSearchByQuery(ctx *gin.Context) {
	ip := ctx.Query("ip")
	if ip == "" {
		ip = ctx.RemoteIP()
	}

	if ip == "::1" {
		ip = "127.0.0.1"
	}

	result, err := h.searchService.Search(ip)
	if err != nil {
		ResponseFail(ctx, err)
		return
	}

	ResponseSuccess(ctx, gin.H{
		"ip":     ip,
		"region": result,
	})
}

func NewSearchHandler(i do.Injector) (*SearchHandler, error) {
	return &SearchHandler{
		searchService: do.MustInvoke[*service.SearchService](i),
	}, nil
}

func init() {
	do.Provide(nil, NewSearchHandler)
}
