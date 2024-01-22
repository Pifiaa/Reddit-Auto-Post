# Este Makefile busca automatizar tareas y simplificar la ejecución de ellas.

.PHONY: help

# Nombre del ejecutable
BINARY_NAME=Reddit-Auto-Post

# Configuración de compilación
GO_BUILD=env GOARCH=amd64
GOOS_DARWIN=darwin
GOOS_LINUX=linux
GOOS_WINDOWS=windows

# Rutas
CMD_DIR=./cmd/app
MIGRATION_DIR=./internal/migration

# Archivo principal
MAIN_FILE=main.go


# Comandos de compilación
# Este comando se utiliza para construir el ejecutable para una plataforma específica.
BUILD_CMD=$(GO_BUILD) go build -o $(BINARY_NAME)-$(1) $(MAIN_FILE)

build: 
	build-darwin build-linux build-windows

build-darwin:
	$(call BUILD_CMD,$(GOOS_DARWIN))

build-linux:
	$(call BUILD_CMD,$(GOOS_LINUX))

build-windows:
	$(call BUILD_CMD,$(GOOS_WINDOWS))


# Comandos de ejecución
# Este comando compila y ejecuta el programa en un solo paso. En lugar de generar un archivo ejecutable por separado.
run: 
	cd $(CMD_DIR) && go run $(MAIN_FILE)


# Comandos de limpieza
# Este comando se utiliza para eliminar archivos y directorios generados durante la compilación y construcción del proyecto.
clean:
	go clean
	rm -f $(BINARY_NAME)-*


# Comandos de migración
# Este comando crea las tablas en la base de datos utilizando migraciones.
migrate-up:
	cd $(MIGRATION_DIR) && goose mysql "root:@/test?parseTime=true" up


# Este comando elimina todas las tablas de la base de datos utilizando migraciones.
migrate-down:
	cd $(MIGRATION_DIR) && goose mysql "root:@/test?parseTime=true" reset


# Este comando crea un nuevo archivo de migración para una tabla específica.
# Uso: make create TABLE=[nombre_de_la_tabla]
create:
	cd $(MIGRATION_DIR) && goose create $(TABLE) sql


# Comandos de ayuda
# Muestra todo los comandos disponibles

help:
	@echo "Uso: make [comando]"
	@echo ""
	@echo "Comandos disponibles:"
	@echo ""
	@echo "   Compilacion:"
	@echo "   --------------------------"
	@echo "   build                  	:   Compilar un programa Go y generar un ejecutable para todas las plataformas."
	@echo "   build-windows          	:   Compilar un programa Go y generar un ejecutable para Windows."
	@echo "   build-linux            	:   Compilar un programa Go y generar un ejecutable para Linux."
	@echo "   build-darwin           	:   Compilar un programa Go y generar un ejecutable para macOS."
	@echo ""
	@echo "   Ejecucion:"
	@echo "   --------------------------"
	@echo "   run                    	:   Compila y ejecuta el programa en un solo paso. En lugar de generar un archivo ejecutable por separado."
	@echo ""
	@echo "   Limpieza:"
	@echo "   --------------------------"
	@echo "   clean                  	:   Elimina archivos y directorios generados durante la compilacion y construcción del proyecto."
	@echo ""
	@echo "   Migracion de Base de Datos:"
	@echo "   --------------------------"
	@echo "   create TABLE=[tabla]   	:   Crea el archivo de migraciones para la tabla especificada."
	@echo "   migrate-up             	:   Crea las tablas en la base de datos."
	@echo "   migrate-down           	:   Elimina las tablas de la base de datos."