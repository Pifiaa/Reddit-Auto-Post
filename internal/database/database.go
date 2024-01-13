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

func DatabaseConnect(config config.Config) (Database, error) {
	password := getPassword(config.Db.Password)
	address := fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port)

	dsn := fmt.Sprintf(
		"%s%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Db.User,
		password,
		address,
		config.Db.DBName,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falló la conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos realizada")

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
