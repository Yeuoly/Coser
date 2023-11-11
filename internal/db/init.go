package db

import (
	"fmt"
	"time"

	"github.com/Yeuoly/coshub/internal/types"
	"github.com/Yeuoly/coshub/internal/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitCoshubDB(host string, port int, user string, pass string) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=coshub sslmode=disable", host, port, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// install postgim
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;").Error
	if err != nil {
		return err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		return err
	}
	pgsqlDB.SetConnMaxIdleTime(time.Minute * 1)

	coshubDB = db
	return nil
}

func RegisterModel() {
	// billboardsDB.AutoMigrate(&types.User{})
	// billboardsDB.AutoMigrate(&types.UserProfile{})
	// billboardsDB.AutoMigrate(&types.Password{})
	// billboardsDB.AutoMigrate(&types.UserGithubBind{})
	err := coshubDB.AutoMigrate(&types.File{})
	if err != nil {
		log.Error("Failed to migrate file model: %s", err.Error())
	}
	err = coshubDB.AutoMigrate(&types.Place{})
	if err != nil {
		log.Error("Failed to migrate place model: %s", err.Error())
	}
	err = coshubDB.AutoMigrate(&types.Gallery{})
	if err != nil {
		log.Error("Failed to migrate gallery model: %s", err.Error())
	}
	err = coshubDB.AutoMigrate(&types.Image{})
	if err != nil {
		log.Error("Failed to migrate image model: %s", err.Error())
	}
	err = coshubDB.AutoMigrate(&types.Tag{})
	if err != nil {
		log.Error("Failed to migrate tag model: %s", err.Error())
	}
}

func InitModel() error {
	return nil
}
