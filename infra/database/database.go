package database

import (
	"fmt"

	"github.com/emaforlin/offr-app-api/config"
	"github.com/emaforlin/offr-app-api/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlRepositoryImpl struct {
	db *gorm.DB
}

func NewDBConn(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Db.User, cfg.Db.Pass, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if cfg.App.Debugmode {
		db.AutoMigrate(
			entities.Account{},
			entities.Profile{},
			entities.Role{},
		)
	}
	return db, nil
}
