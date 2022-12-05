package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/router"

	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/database"
)

func main() {
	config.Init()
	cfg := config.GetConfig()

	r := gin.Default()
	db := database.GetDatabase()

	logger, _ := zap.NewProduction()

	router.Init(r, db, logger)

	address := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port)
	r.Run(address)
}
