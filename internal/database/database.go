package database

import (
	"RedditAutoPost/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	GetDb() *gorm.DB
	Close()
}

type connectionDatabase struct {
	Db *gorm.DB
}

func DatabaseConnect(cfg *config.Config) (Database, error) {
	password := getPassword(cfg.Db.Password)
	address := fmt.Sprintf("%s:%s", cfg.Db.Host, cfg.Db.Port)

	dsn := fmt.Sprintf(
		"%s%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Db.User,
		password,
		address,
		cfg.Db.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falló la conexión a la base de datos: %v", err)
	}

	return &connectionDatabase{Db: db}, nil
}

func getPassword(password string) string {
	if password != "" {
		return ":" + password
	}
	return ""
}

func (c *connectionDatabase) GetDb() *gorm.DB {
	return c.Db
}

func (c *connectionDatabase) Close() {
	sqlDB, err := c.Db.DB()
	if err != nil {
		fmt.Printf("Error al obtener la conexión a la base de datos: %v\n", err)
		return
	}
	sqlDB.Close()
}
