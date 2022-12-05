package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	pingHandlerPkg "github.com/maheswaradevo/hacktiv8-finalproject4/internal/ping/handler"
	pingService "github.com/maheswaradevo/hacktiv8-finalproject4/internal/ping/service"
	"go.uber.org/zap"

	authHandlerPkg "github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth/handler"
	authRepository "github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth/repository"
	authService "github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth/service"
)

func Init(router *gin.Engine, db *sql.DB, logger *zap.Logger) {
	api := router.Group("api/v1")
	{
		InitPingModule(api)

		InitAuthModule(api, db, logger)
	}
}

func InitPingModule(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	pingService := pingService.NewPingService()
	return pingHandlerPkg.NewPingHandler(routerGroup, pingService)
}

func InitAuthModule(routerGroup *gin.RouterGroup, db *sql.DB, logger *zap.Logger) *gin.RouterGroup {
	authRepository := authRepository.NewAuthRepository(db, logger)
	authService := authService.NewAuthService(authRepository, logger)
	return authHandlerPkg.NewUserHandler(routerGroup, authService)
}
