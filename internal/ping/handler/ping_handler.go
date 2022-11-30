package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/ping"
)

type PingHandler struct {
	r  *gin.RouterGroup
	ps ping.Ping
}

func NewPingHandler(r *gin.RouterGroup, ps ping.Ping) *gin.RouterGroup {
	delivery := PingHandler{
		r:  r,
		ps: ps,
	}
	pingRouter := delivery.r.Group("/ping")
	pingRouter.Handle(http.MethodGet, "", delivery.Ping)

	return pingRouter
}

func (p *PingHandler) Ping(ctx *gin.Context) {
	res := p.ps.Ping()
	response := utils.NewSuccessResponseWriter(ctx.Writer, "SUKSES", http.StatusOK, res)
	ctx.JSON(http.StatusOK, response)
}
