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
		return nil, fmt.Errorf("falló la conexión a la base de datos: %v", err)
	}

	// Configurar y establecer un pool de conexiones si es necesario
	sqlDB, err := connection.DB()
	if err != nil {
		return nil, fmt.Errorf("falló al obtener la instancia de la base de datos: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)  // Ajustar según sea necesario
	sqlDB.SetMaxOpenConns(100) // Ajustar según sea necesario

	db = connection // Asignar la instancia de la base de datos a la variable del paquete

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
