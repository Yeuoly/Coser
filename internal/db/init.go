package db

import (
	"fmt"
	"time"

	"github.com/Yeuoly/billboards/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitBocchiDB(host string, port int, user string, pass string) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=bocchi sslmode=disable", host, port, user, pass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		return err
	}
	pgsqlDB.SetConnMaxIdleTime(time.Minute * 1)

	bocchiDB = db
	return nil
}

func RegisterModel() {
	bocchiDB.AutoMigrate(&types.User{})
	bocchiDB.AutoMigrate(&types.UserProfile{})
	bocchiDB.AutoMigrate(&types.Password{})
	bocchiDB.AutoMigrate(&types.UserGithubBind{})
}

func InitModel() error {
	return nil
}
