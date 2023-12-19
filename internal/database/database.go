package database

import (
	"RedditAutoPost/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // Cambia el nombre de la variable a algo más descriptivo

// Connect inicializa una conexión a la base de datos y devuelve una instancia de Gorm DB.
func Connect() (*gorm.DB, error) {
	user := config.GetEnv("database.username")
	password := config.GetEnv("database.password")
	database := config.GetEnv("database.name")
	addr := fmt.Sprintf("%s:%s", config.GetEnv("database.host"), config.GetEnv("database.port"))

	var dsn string
	dsn = fmt.Sprintf("%s%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		func() string {
			if password != "" {
				return ":" + password
			}
			return ""
		}(),
		addr,
		database,
	)

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falló la conexión a la base de datos: %v", err)
	}

	db = connection // Asignar la instancia de la base de datos a la variable del paquete
	fmt.Println("Conexión a la base de datos establecida")

	return connection, nil
}

// Close cierra la conexión a la base de datos.
func Close() {
	if db != nil {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
