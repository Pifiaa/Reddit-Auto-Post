package database

import (
	"RedditAutoPost/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func Connect() (*sql.DB, error) {

	dsn := mysql.Config{
		AllowNativePasswords: true,
		Net:                  "tcp",
		DBName:               config.GetEnv("database.name"),
		User:                 config.GetEnv("database.username"),
		Passwd:               config.GetEnv("database.password"),
		Addr:                 fmt.Sprintf("%s:%s", config.GetEnv("database.host"), config.GetEnv("database.port")),
	}

	connection, err := sql.Open("mysql", dsn.FormatDSN())

	if err != nil {
		return nil, err
	}

	pingErr := connection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	return connection, nil
}

func close() {
	connection.Close()
}
