# bootcamp-payment-api

## Project Description

This is a RESTful API developed in Go for managing products. It uses the Gin framework to handle HTTP routes and
PostgreSQL database for data persistence. The API allows CRUD operations (Create, Read, Update, Delete) on products.

## Project Structure

- **`repository/product_repository.go`**: Contains database access logic for product manipulation.
- **`controller/product_controller.go`**: Defines controllers that handle HTTP requests and call use cases.
- **`db/conn.go`**: Configures PostgreSQL database connection.
- **`models`**: Contains data structure definitions used in the project.
- **`usecase`**: Implements application business logic.

## Prerequisites

- Go 1.20 or higher
- PostgreSQL installed and configured
- `go mod` enabled in project

## Database Configuration

1. Create a PostgreSQL database named `bootcamp_db`.
2. Configure access credentials in file `db/conn.go`:
   ```go
   const (
       host     = "localhost"
       port     = 5432
       user     = "user"
       password = "password"
       dbname   = "bootcamp_db"
   )
    ```
3. ## To-Do List
- [ ] Add input validation in controllers.
- [ ] Create .env file for environment variables.
- [ ] Implement product update functionality (PUT endpoint).
- [ ] Implement product deletion functionality (DELETE endpoint).
- [ ] Add unit tests for repositories and controllers.
- [x] Create detailed API documentation using Swagger.
- [ ] Set up CI/CD for the project.