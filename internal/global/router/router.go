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

	transactionHandler "github.com/maheswaradevo/hacktiv8-finalproject4/internal/transaction/handler"
	transactionRepository "github.com/maheswaradevo/hacktiv8-finalproject4/internal/transaction/repository"
	transactionService "github.com/maheswaradevo/hacktiv8-finalproject4/internal/transaction/service"
)

func Init(router *gin.Engine, db *sql.DB, logger *zap.Logger) {
	api := router.Group("api/v1")
	{
		InitPingModule(api)

		InitAuthModule(api, db, logger)

		InitProductModule(api, db, logger)

		InitCategoriesModule(api, db, logger)

		InitTransactionModule(api, db, logger)
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

func InitProductModule(routerGroup *gin.RouterGroup, db *sql.DB, logger *zap.Logger) *gin.RouterGroup {
	productRepository := productRepository.ProvideProductRepository(db, logger)
	productService := productService.ProvideProductService(productRepository, logger)
	return productHandler.NewProductHandler(routerGroup, productService, logger)
}

func InitTransactionModule(routerGroup *gin.RouterGroup, db *sql.DB, logger *zap.Logger) *gin.RouterGroup {
	transactionRepository := transactionRepository.NewTransactionRepository(db, logger)
	productRepository := productRepository.ProvideProductRepository(db, logger)
	authRepository := authRepository.NewAuthRepository(db, logger)

	transactionService := transactionService.NewTransactionService(transactionRepository, productRepository, authRepository, logger)
	return transactionHandler.NewTransactionHandler(routerGroup, transactionService, logger)
}

func InitCategoriesModule(routerGroup *gin.RouterGroup, db *sql.DB, logger *zap.Logger) *gin.RouterGroup {
	transactionRepository := transactionRepository.NewTransactionRepository(db, logger)
	categoriesRepository := categoriesRepository.ProvideCategoriesRepository(db)
	categoriesService := categoriesService.ProvideCategoriesService(categoriesRepository, transactionRepository)
	return categoriesHandler.NewCategoriesHandler(routerGroup, categoriesService)
}
