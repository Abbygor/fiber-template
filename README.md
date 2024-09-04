# Fiber Template API

Este es un proyecto template de ejemplo que utiliza Fiber, GORM y PostgreSQL para implementar un API REST. Se sigue una estructura modular y buenas prácticas en Go.

## Estructura del Proyecto

```plaintext
fiber-template/
│
├── .vscode/
│   └── launch.json
├── cmd/
│   ├── constants/
│   │   └── constants.go
|   ├── httpserver
│   │   ├── http_server.go
│   |   └── routes.go
│   └── main.go
├── docker/
│   ├── postgres/
│   │   ├── Dockerfile
│   |   └── initDB.sql
│   └── docker-compose.yaml
├── internal/
│   ├── app/
│   │   ├── health/
│   │   │   ├── controller.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   ├── config/
│   │   └── config.go
│   ├── container/
│   │   └── container.go
│   └── models/
│       └── health.go
├── pkg/
│   └── gorm/
│       └── gorm.go
├── go.mod
└── go.sum
```
## Instalación
### Clona el repositorio:

```bash
git clone https://github.com/tu_usuario/fiber-template.git
cd fiber-template
```

### Instala las dependencias:
```bash
go mod tidy
```
### Actualiza las variables del launch.json con la informacion necesaria:
```bash
"POSTGRES_SERVER": "localhost",
"POSTGRES_DATABASE": "testdb",
"POSTGRES_USER": "postgres",
"POSTGRES_PASSWORD": "root",
"POSTGRES_PORT": "5432"
```
Inicia la base de datos PostgreSQL y asegúrate de que las credenciales en el archivo launch.json son correctas.

## Ejecuciónde forma local
1. Ir a la carpeta **docker**
```bash
cd docker/
```
2. Ejecutar la construcion de los contenedores
```bash
docker-compose up
```
3. Con esto deberia levantar el contenedor de la base de datos y el **health** y **health/dependencies** deberian funcionar correctamente (modificar el archivo docker-compose.yaml segun sea requerido, agregando o quitando servicios).

4. Para ejecutar el proyecto:
```bash
go run main.go
```
## Endpoints

- `GET /health` - Obtiene el estado de la aplicación.
- `GET /health/dependencies` - Obtiene el estado de las dependencias que ocupa el proyecto (se deben ir agregando).

## Estructura de los Paquetes

- **.vscode**: Archivo con las variables necesarias para ejecutar el proyecto en modo debug.
- **cmd/constants**: Contiene el archivo de constantes que se ocupan en el proyecto.
- **cmd/httpserver**: Contiene la configuración para levantar el server http y sus respectivas rotas.
- **cmd**: Contiene el archivo de main para arrancar el proyecto.
- **docker**: Contiene los archivo necesarios para realizar pruebas en local (devcontainers).
- **docker/postgres**: Contiene los archivos para iniciar el contenedor con la BD, 
- **internal/app**: Contiene la lógica de la aplicación, incluyendo controladores, repositorios y servicios.
- **internal/config**: Contiene la configuración de la aplicación.
- - **internal/container**: Contiene la inicialización y gestión de dependencias.
- **internal/models**: Contiene los modelos de datos específicos de la aplicación.
- **pkg/gorm**: Maneja la conexión a la base de datos.

## Tecnologías Utilizadas

- [Fiber](https://gofiber.io/) - Un framework web inspirado en Express.js.
- [GORM](https://gorm.io/) - Un ORM para Go.
- [PostgreSQL](https://www.postgresql.org/) - Un sistema de gestión de bases de datos relacional.

## Contribuciones

Las contribuciones son bienvenidas. Por favor, sigue los siguientes pasos:

1. Haz un fork del proyecto.
2. Crea una nueva rama (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios.
4. Haz un commit de tus cambios (`git commit -am 'Agrega nueva funcionalidad'`).
5. Haz push a la rama (`git push origin feature/nueva-funcionalidad`).
6. Abre un Pull Request.

## Licencia

Este proyecto está licenciado bajo la [MIT License](LICENSE).
