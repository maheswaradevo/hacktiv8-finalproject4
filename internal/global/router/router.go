package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	pingHandlerPkg "github.com/maheswaradevo/hacktiv8-finalproject4/internal/ping/handler"
	pingService "github.com/maheswaradevo/hacktiv8-finalproject4/internal/ping/service"
)

func Init(router *gin.Engine, db *sql.DB) {
	api := router.Group("api/v1")
	{
		InitPingModule(api)
	}
}

func InitPingModule(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	pingService := pingService.NewPingService()
	return pingHandlerPkg.NewPingHandler(routerGroup, pingService)
}
