package database

import (
	"RedditAutoPost/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var connection *sql.DB

func Connect() (*sql.DB, error) {
	isConfigEnable, err := config.InitializeConfig()

	if !isConfigEnable {
		fmt.Println("Error:", err)
	}

	hostname := config.GetEnv("database.host")
	port := config.GetEnv("database.port")
	username := config.GetEnv("database.username")
	password := config.GetEnv("database.password")
	name := config.GetEnv("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, name)
	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func close() {
	connection.Close()
}
