# Reddit-Auto-Post <img src="https://seeklogo.com/images/R/reddit-logo-23F13F6A6A-seeklogo.com.png" alt="Reddit" height="30" width="30" margin="auto" display="block" >


## Descripción
Proyecto creado con el proposito de automatizar publicaciones en Reddit, agilizando la gestión de contenidos con facilidad y eficacia.

> "*La felicidad está en las pequeñas cosas, como un buen par de tetas*."
> Ettolini - 2023


## Requerimientos de la aplicación
1. Instalar y configurar git 
* https://git-scm.com/downloads
* [Pasos para la configuración](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup)
  
2. Instalar Make
      * Opción 1:
        Descargar directamente de:
        https://gnuwin32.sourceforge.net/packages/     
      * Opción 2:
        - Instalar Chocolatey
            https://chocolatey.org/install 
        - Ejecutar el comando:
            ```shell
            choco install make
            ``` 

3. Instalación de Go version - 1.21.4
* https://go.dev/doc/install

4. Una vez instalado, compruébe con el siguiente comando en linea de comandos.
```bash   
    go version
``` 
    
## Inicializar la aplicación localmente
1. Clonar el repositorio.
 ```bash   
    git clone https://github.com/Pifiaa/Reddit-Auto-Post.git
```

1. Ejecute el comando.
```bash   
    go mod download
```
Este es un comando utilizado para descargar los módulos necesarios para construir y ejecutar un programa Go.

### Paquetes y herramientas instalados
* [gin-gonic](https://gin-gonic.com/es/): Gin-gonic es un framework web para Go. Ayuda a construir aplicaciones web y APIs de manera más fácil en el lenguaje de programación Go.
* [viper](https://github.com/spf13/viper): Viper es una librería de configuración. Permite manejar la configuración de tu aplicación de forma sencilla y flexible, con soporte para diferentes formatos de archivos y fuentes de configuración.
* [go-sql-driver](https://github.com/go-sql-driver/mysql): Go-sql-driver es un controlador de base de datos. Este controlador se utiliza para interactuar con bases de datos SQL.
* [goose](https://github.com/pressly/goose) - Goose es una Herramienta para gestionar cambios estructurales en bases de datos de manera ordenada.


## Estructura de directorios
```bash
REDDIT-AUTO-POST
├── api
│   ├── handler
│   │   ├── post.go
│   │   └── token.go
│   ├── routes
│   │   └── router.go
│   └── app.go
│
├── cmd
│   └── app 
│       └── main.go
│
├── config
│   └── config.go
│
├── internal
│   └── database
│       ├── migration
│       │   ├── 20231211134531_subreddits.sql
│       │   ├── 20231211134759_posts.sql
│       │   └── 20231212124636_redditcredentials.sql
│       └── database.go
│
├── config.yml
├── go.mod
├── go.sum
├── Makefile
└── README.md

```

## Endpoints
| Nombre                   | Metodo HTTP    | Rutas                       |
|--------------------------|----------------|-----------------------------|
| Obtener token            | POST           | /api/v1/access-token        |
| Refrescar token          | PUT            | /api/v1/refresh-token       |
| Nuevo post               | POST           | /api/v1/create-post         |
| Ver post                 | GET            | /api/v1/posts               |
| Nuevo subreddit          | POST           | /api/v1/create-subreddit    |
| Ver subreddit            | GET            | /api/v1/subreddits          |
| Eliminar subreddit       | DELETE         | /api/v1/delete-subreddit    |
| Subir archivo            | POST           | /api/v1/upload-file         |
| Agregar credenciales     | POST           | /api/v1/create-credentials  |
| Ver credenciales         | GET            | /api/v1/credentials         |
| Actualizar credenciales  | PUT            | /api/v1/update-credentials  |
| Eliminar credenciales    | DELETE         | /api/v1/delete-credentials  |