package database

import (
	"RedditAutoPost/config"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *sql.DB

func Connect() (*gorm.DB, error) {

	addr := fmt.Sprintf("%s:%s", config.GetEnv("database.host"), config.GetEnv("database.port"))
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)%s/?charset=utf8&parseTime=True&loc=Local",
		config.GetEnv("database.username"),
		config.GetEnv("database.password"),
		addr,
		config.GetEnv("database.name"),
	)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func close() {
	connection.Close()
}
