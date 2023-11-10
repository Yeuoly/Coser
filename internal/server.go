package internal

import (
	"fmt"

	"github.com/Yeuoly/coshub/internal/db"
	"github.com/Yeuoly/coshub/internal/router"
	"github.com/Yeuoly/coshub/internal/static"
	"github.com/Yeuoly/coshub/internal/utils/log"
	"github.com/Yeuoly/coshub/internal/utils/oss"
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
	config := static.GetCoshubGlobalConfigurations()
	err := db.InitCoshubDB(config.DB.Host, config.DB.Port, config.DB.User, config.DB.Pass)
	if err != nil {
		log.Panic("failed to init database: %v", err)
	}
	log.Info("database init success")
	db.RegisterModel()
	log.Info("database register model success")

	err = db.InitModel()
	if err != nil {
		log.Panic("failed to init model: %v", err)
	}
	log.Info("model init success")
}

func initServer() {
	config := static.GetCoshubGlobalConfigurations()
	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	router.Setup(r, &config)

	r.Run(fmt.Sprintf(":%d", config.App.Port))
}

func initMiddleware() {
	oss.InitOSS()
}

func Server() {
	initConfig()
	initDB()
	initMiddleware()
	initServer()
}
