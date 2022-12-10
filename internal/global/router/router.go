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

	productHandler "github.com/maheswaradevo/hacktiv8-finalproject4/internal/product/handler"
	productRepository "github.com/maheswaradevo/hacktiv8-finalproject4/internal/product/repository"
	productService "github.com/maheswaradevo/hacktiv8-finalproject4/internal/product/service"

	categoriesHandler "github.com/maheswaradevo/hacktiv8-finalproject4/internal/categories/handler"
	categoriesRepository "github.com/maheswaradevo/hacktiv8-finalproject4/internal/categories/repository"
	categoriesService "github.com/maheswaradevo/hacktiv8-finalproject4/internal/categories/service"
)

func Init(router *gin.Engine, db *sql.DB, logger *zap.Logger) {
	api := router.Group("api/v1")
	{
		InitPingModule(api)

		InitAuthModule(api, db, logger)

		InitProductModule(api, db)

		InitCategoriesModule(api, db)
	}
}

func InitPingModule(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	pingService := pingService.NewPingService()
	return pingHandlerPkg.NewPingHandler(routerGroup, pingService)
}

func InitAuthModule(routerGroup *gin.RouterGroup, db *sql.DB, logger *zap.Logger) *gin.RouterGroup {
	authRepository := authRepository.NewAuthRepository(db, logger)
	authService := authService.NewAuthService(authRepository, logger)
	return authHandlerPkg.NewUserHandler(routerGroup, authService, logger)
}

func InitProductModule(routerGroup *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	productRepository := productRepository.ProvideProductRepository(db)
	productService := productService.ProvideProductService(productRepository)
	return productHandler.NewProductHandler(routerGroup, productService)
}

func InitCategoriesModule(routerGroup *gin.RouterGroup, db *sql.DB) *gin.RouterGroup {
	categoriesRepository := categoriesRepository.ProvideCategoriesRepository(db)
	categoriesService := categoriesService.ProvideCategoriesService(categoriesRepository)
	return categoriesHandler.NewCategoriesHandler(routerGroup, categoriesService)
}
