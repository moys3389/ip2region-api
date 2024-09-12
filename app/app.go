package app

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moys3389/ip2region-api/config"
	"github.com/moys3389/ip2region-api/handler"
	"github.com/samber/do/v2"
)

type App struct {
	server        *gin.Engine
	searchHandler *handler.SearchHandler
	cfg           *config.Config
}

func (a *App) Start() error {
	if len(a.cfg.Cors) > 0 {
		cfg := cors.DefaultConfig()
		cfg.AllowOrigins = strings.Split(a.cfg.Cors, ",")
		a.server.Use(cors.New(cfg))
	}

	apiGroup := a.server.Group("api")
	{
		apiGroup.GET("search", a.searchHandler.HandleSearchByQuery)
		apiGroup.POST("search", a.searchHandler.HandleSearch)
		apiGroup.GET("batch-search", a.searchHandler.HandleBatchSearchByQuery)
		apiGroup.POST("batch-search", a.searchHandler.HandleBatchSearch)
		apiGroup.GET("version", func(ctx *gin.Context) { ctx.String(http.StatusOK, "version: "+a.cfg.Version) })
	}

	return a.server.Run()
}

func NewApp(i do.Injector) (*App, error) {
	return &App{
		server:        gin.Default(),
		searchHandler: do.MustInvoke[*handler.SearchHandler](i),
		cfg:           do.MustInvoke[*config.Config](i),
	}, nil
}

func init() {
	do.Provide(nil, NewApp)
}
