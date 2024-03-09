package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(viperConfig *viper.Viper) (*gorm.DB, error) {
	user := viperConfig.GetString("db_user")
	pass := viperConfig.GetString("db_pass")
	host := viperConfig.GetString("db_host")
	port := viperConfig.GetString("db_port")
	dbName := viperConfig.GetString("db_name")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
