package internal

import (
	"fmt"

	"github.com/Yeuoly/billboards/internal/db"
	"github.com/Yeuoly/billboards/internal/router"
	"github.com/Yeuoly/billboards/internal/static"
	"github.com/Yeuoly/billboards/internal/utils/log"
	"github.com/gin-gonic/gin"
)

func initConfig() {
	// auto migrate database
	err := static.InitConfig("conf/config.yaml")
	if err != nil {
		log.Panic("failed to init config: %v", err)
	}
	log.Info("config init success")
}

func initDB() {
	config := static.GetBillboardsGlobalConfigurations()
	err := db.InitBillboardsDB(config.DB.Host, config.DB.Port, config.DB.User, config.DB.Pass)
	if err != nil {
		log.Panic("failed to init database: %v", err)
	}
	log.Info("database init success")
	db.RegisterModel()
	log.Info("database register model success")

	db.InitBillboardsRedis()
	log.Info("redis init success")

	err = db.InitModel()
	if err != nil {
		log.Panic("failed to init model: %v", err)
	}
	log.Info("model init success")
}

func initServer() {
	config := static.GetBillboardsGlobalConfigurations()
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	router.Setup(r, &config)

	r.Run(fmt.Sprintf(":%d", config.App.Port))
}

func initMiddleware() {
	db.InitMinio()
}

func Server() {
	initConfig()
	initDB()
	initMiddleware()
	initServer()
}
