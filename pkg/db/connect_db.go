package db

import (
	"fmt"

	"github.com/ajalck/spine/pkg/config"
	"github.com/ajalck/spine/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = SyncDB(db); err != nil {
		return nil, err
	}
	return db, nil
}
func SyncDB(db *gorm.DB) error {
	
	return db.AutoMigrate(
		&domain.User{},
	)
}
