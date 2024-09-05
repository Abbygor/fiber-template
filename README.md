# Fiber Template RestAPI

This is a template project example that uses Fiber, GORM, and PostgreSQL to implement a REST API. It follows a modular structure and best practices in Go.

## Project Structure

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
## Installation
### Clone the repository:

```bash
git clone https://github.com/tu_usuario/fiber-template.git
cd fiber-template
```

### Install dependencies:
```bash
go mod tidy
```
### Update the variables in launch.json with the required information:
```bash
"POSTGRES_SERVER": "localhost",
"POSTGRES_DATABASE": "testdb",
"POSTGRES_USER": "postgres",
"POSTGRES_PASSWORD": "root",
"POSTGRES_PORT": "5432"
```
Start the PostgreSQL database and make sure the credentials in the launch.json file are correct.

## Running Locally
1. Go to the **docker** folder
```bash
cd docker/
```
2. Run the container build
```bash
docker-compose up
```

3. This should start the database container, and the health and health/dependencies should work correctly (modify the docker-compose.yaml file as needed, adding or removing services).

4. To run the project:
```bash
go run main.go
```
## Endpoints

- `GET /health` - Gets the application status.
- `GET /health/dependencies` - Gets the status of the project's dependencies (should be added as needed).

## Package Structure

- **.vscode**: Contains the necessary variables to run the project in debug mode.
- **cmd/constants**: Contains the constants used in the project.
- **cmd/httpserver**: Contains the configuration to start the HTTP server and its routes.
- **cmd**: Contains the main file to start the project.
- **docker**: Contains the files needed for local testing (devcontainers).
- **docker/postgres**: Contains the files to start the container with the database.
- **internal/app**: Contains the application logic, including controllers, repositories, and services.
- **internal/config**: Contains the application configuration.
- **internal/container**: Contains the initialization and management of dependencies.
- **internal/models**: Contains the application's data models.
- **pkg/gorm**: Handles the database connection.

## Tecnologías Utilizadas

- [Fiber](https://gofiber.io/) - A web framework inspired by Express.js.
- [GORM](https://gorm.io/) - An ORM for Go.
- [PostgreSQL](https://www.postgresql.org/) - A relational database management system.

## Contribuciones

Contributions are welcome. Please follow these steps:

1. Fork the project.
2. Create a new branch (git checkout -b feature/new-feature).
3. Make your changes.
4. Commit your changes (git commit -am 'Add new feature').
5. Push to the branch (git push origin feature/new-feature).
6. Open a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
