# Este Makefile busca automatizar tareas y simplificar la ejecución de ellas.


# Nombre del ejecutable
BINARY_NAME=Reddit-Auto-Post


#Este comando se utiliza para compilar un programa Go y generar un ejecutable.
build:
	env GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	env GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	env GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go


#Este comando compila y ejecuta el programa en un solo paso. En lugar de generar un archivo ejecutable por separado.
run: build
	./${BINARY_NAME}


# Este comando se utiliza para eliminar archivos y directorios generados durante la compilación y construcción del proyecto.
clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows