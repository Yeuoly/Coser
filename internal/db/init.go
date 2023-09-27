package db

import (
	"fmt"
	"time"

	"github.com/Yeuoly/billboards/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitBillboardsDB(host string, port int, user string, pass string) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=billboards sslmode=disable", host, port, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		return err
	}
	pgsqlDB.SetConnMaxIdleTime(time.Minute * 1)

	billboardsDB = db
	return nil
}

func RegisterModel() {
	billboardsDB.AutoMigrate(&types.User{})
	billboardsDB.AutoMigrate(&types.UserProfile{})
	billboardsDB.AutoMigrate(&types.Password{})
	billboardsDB.AutoMigrate(&types.UserGithubBind{})
}

func InitModel() error {
	return nil
}
