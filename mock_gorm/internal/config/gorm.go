package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseOpener interface {
	Open(dsn string) (*gorm.DB, error)
}

type GormOpener struct{}

func (g *GormOpener) Open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func NewDatabase(viperConfig Viper, opener DatabaseOpener) (*gorm.DB, error) {
	user := viperConfig.GetString("db_user")
	pass := viperConfig.GetString("db_pass")
	host := viperConfig.GetString("db_host")
	port := viperConfig.GetString("db_port")
	dbName := viperConfig.GetString("db_name")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)
	return opener.Open(dsn)
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
