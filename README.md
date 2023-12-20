# Reddit-Auto-Post <img src="https://seeklogo.com/images/R/reddit-logo-23F13F6A6A-seeklogo.com.png" alt="Reddit" height="30" width="30" margin="auto" display="block" >

---

## Descripción
Proyecto creado con el propósito de automatizar publicaciones en Reddit, agilizando la gestión de contenidos con facilidad y eficacia. Simplifica la tarea de programar y realizar publicaciones en la plataforma Reddit. Esta herramienta ofrece una solución eficiente para gestionar y programar publicaciones, mejorando la experiencia de los usuarios en la plataforma y optimizando el flujo de trabajo en la administración de contenido.


> "*La felicidad está en las pequeñas cosas, como un buen par de tetas*."
> Ettolini - 2023

---

## Requerimientos de la aplicación
1. Instalar y configurar git
- [Descargar Git](https://git-scm.com/downloads)
- [Pasos para la configuración](https://git-scm.com/book/en/v2/Getting-Started-First-Time-Git-Setup)

2. Instalar Make
- Opción 1: Descargar directamente de [GNUwin32](https://gnuwin32.sourceforge.net/packages/)
- Opción 2: Instalar [Chocolatey](https://chocolatey.org/install) y ejecutar el comando `choco install make`

3. Instalación de Go (Versión utilizada: 1.21.5)
- [Instrucciones de instalación de Go](https://go.dev/doc/install)
- Verificar la instalación con el comando `go version`  
    
## Inicializar la aplicación localmente
1. Clonar el repositorio.
 ```bash   
    git clone https://github.com/Pifiaa/Reddit-Auto-Post.git
```

2. Ejecute el comando.
```bash   
    go mod download
```
Este es un comando utilizado para descargar los módulos necesarios para construir y ejecutar un programa Go.

---

### Paquetes y Herramientas Instalados
1. [gin-gonic](https://gin-gonic.com/es/): Framework web para Go que facilita la construcción de aplicaciones web y APIs de manera eficiente.

2. [viper](https://github.com/spf13/viper): Librería de configuración flexible que simplifica el manejo de la configuración de la aplicación con soporte para diversos formatos y fuentes.

3. [go-sql-driver](https://github.com/go-sql-driver/mysql): Controlador de base de datos SQL para interactuar con bases de datos SQL mediante el lenguaje de programación Go.

4. [goose](https://github.com/pressly/goose): Herramienta que facilita la gestión de cambios estructurales en bases de datos, asegurando una evolución ordenada.

5. [gorm](https://gorm.io/): ORM (Object-Relational Mapping) para Go que simplifica la interacción con bases de datos relacionales a través de operaciones orientadas a objetos.

---

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

---

## Endpoints
| Nombre                   | Metodo HTTP    | Rutas                       |
|--------------------------|----------------|-----------------------------|
| Obtener token            | POST           | /api/v1/access-token        |
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

<style>
  tr:nth-child(even) {
    background-color: #393d42;
  }
</style>