# Transaction API

## Project Overwiew

This project was developed as part of the technical challenge for the Traive Engineering interview. The goal is to build a RESTful API in Go (Golang), allowing the management of user transactions. The API will enable the creation and listing of transactions with features like pagination and filtering.
The project has been time-boxed to a maximum of 5 hours in order to showcase the candidate's ability to deliver an application within a limited timeframe.

## Design Assumptions

- **Framework**: The project uses the Gin web framework for its lightweight nature and efficient performance in building RESTful APIs.
- **Database**: PostgreSQL is chosen for its reliability and feature-rich support for transactional data management.
- **Transactions**: Each transaction includes details such as `ID`, `origin`, `user ID`, `amount`, `transaction type` (credit/debit), and `createdAt` timestamp. Additional attributes have not been considered at this stage.
- **Pagination**: The API supports pagination for listing transactions, allowing users to navigate through large datasets efficiently.
- **Filtering**: The API supports filtering transactions based on `origin` and `transaction type`, providing users with the flexibility to query specific records.

## Technical Challenge Requirements

- Develop an API to manage transactions with capabilities to create and list transactions. Please note that as an ORM of choice, we ended up using [[Bun](https://bun.uptrace.dev/)], a lightweight ORM for Go, which is a wrapper around the `pgx` PostgreSQL driver.
- Ensure the API supports pagination and filtering for listing transactions.
- Aim for a production-ready application in terms of project structure, code organization, technology choices and best practices. 
- Incorporate error handling and robust application behavior analysis strategies.
- Utilise a consistent naming convention and maintain code readability.
- Include comprehensive documentation and instructions for running the application as well as Swagger documentation for the API endpoints.
- Implement a Dockerized setup for the application and database, ensuring easy deployment and management.
- Integrate OpenTelemetry for capturing telemetry data and trace requests, and, not part of this project, generate custom metrics. We just showcased the use of OpenTelemetry to instrument the application's handlers. 

## Project Structure

The project follows a standard Go project layout with separate directories for the application, database, and Docker configurations. The main components of the project structure are as follows:

- **`/api`**: Contains the API handlers and routes for managing transactions.
- **`/config`**: Includes the application configuration settings and environment variables.
- **`/domain`**: Contains the domain model and business logic for managing transactions. 
- **`/repository/model`**: Contains the database model and repository for interacting with the database. It also includes mappings between the domain and database models.
- **`/repository/postgres`**: Contains the PostgreSQL repository for handling database operations.
- **`/repository/filter`**: Handles filtering options 
- **`/service`**: Contains the service layer for handling business logic and data operations.
- **`/support`**: Contains utility functions and helper methods used for testing purposes.
- **`/cmd`**: It includes `main.go` that is the entry point of the application that initializes the server and routes.
- **`Dockerfile`**: The Docker configuration file for building the application image.
- **`docker-compose.yml`**: The Docker Compose configuration for setting up the PostgreSQL database.
- **`/postgresql`**: The directory for the init script for the PostgreSQL database.
- **`/docs`**: The directory for the swagger documentation.
- **`Makefile`**: The Makefile for defining common tasks and commands for the application.
- **`README.md`**: The project documentation and instructions for running the application.

Some References for the project structure:
- [Go Project Layout](https://go.dev/doc/code#Organization)
- [Applied Go - Project Layout](https://appliedgo.com/blog/go-project-layout)

## Prerequisites
Pre-requisites to run the application are as follow: 

- Go (version 1.x) - currently using 1.21
- PostgreSQL
- Docker (for containerized setup)
- Local Database Setup and Testing

In terms of Docker, the following are required:

- [Docker](https://www.docker.com/products/docker-desktop) installed on your machine.
- [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine.

## Setup and Running the Application

### Starting the Database

To start the PostgreSQL database using Docker Compose, run the following command:

```sh
docker-compose up -d postgres
```

   This will initialize and start a PostgreSQL container in detached mode.

Alternatively:

```
make start-db
```

### Stopping the Database

To stop the PostgreSQL database container, run the following command:

```sh
docker-compose down --volumes
```

Alternatively:

```
make stop-db
```

### Building and Running the Application

Build the application binary and start the application using the Makefile:

```sh
make run
```

   This command performs the following actions:
    - Ensures the database container is up.
    - Builds the application binary.
    - Starts the application.

### Accessing the Application

Once the application is running, it will be accessible at `http://localhost:8080`. You can interact with the application's API endpoints through this URL.

## Running Tests

To run the unit tests for the application, execute the following command:

```sh
make test
```

Please note that currently the projects only have unit tests, and no integration tests need implementing.

## Additional Commands

The Makefile includes additional commands for common tasks such as building the application, cleaning up the environment, and running the application in development mode. For example, to clean up the environment and remove the application binary, run the following command:

```sh
make clean
```

To format the Go code using `gofmt`, run the following command:

```sh
make gofumpt
```

## Configuration

The application is configured through environment variables. The primary configuration is the database connection string, managed via the `DATABASE_URL` environment variable.

### Default Configuration

By default, if no `DATABASE_URL` is provided, the application uses:

- **Database URL**: `postgres://postgres:password@localhost:5432/transactions?sslmode=disable`

Alternatively, in `docker-compose.yml`, the database URL is set as an environment variable for the PostgreSQL container.


```yaml
app:
  environment:
    - DATABASE_URL=postgresql://postgres:password@postgres:5432/transactions-app_development?sslmode=disable
```

## Customizing Environment Variables
To customize, you can modify the value directly in `docker-compose.yml` or use a `.env file` with Docker Compose to define DATABASE_URL.

### Interacting with Endpoints

Swagger provides an interactive interface where you can:

- View all available endpoints
- See details about each endpoint, including request and response formats
- Test endpoints directly from the Swagger UI by providing input parameters and sending requests
- View sample responses and error messages

This documentation is useful for developers who want to explore and test the API without having to refer to the codebase or external documentation.

The swagger docs are available at `http://localhost:8080/swagger/index.html`. and they were already generated using the `swag` package, by running the following command in the root of the project: 

```
swag init -g main.go 
```
and saved in the `docs` folder.

Alternatively, you can use the Makefile to generate the swagger docs by running the following command:

```sh
make swagger-gen
```

## Future Improvements for Production Readiness

### 1. Improve Testing Strategy
- Enhance Unit Tests and introduce Integration Tests to ensure comprehensive test coverage and identify potential issues across the application's components.

### 2. Enhance Filtering Capabilities
- Expand Filtering capabilities to include additional fields, providing users with more flexibility and options when querying transactions.

### 3. Implement Robust Pagination Mechanism
- Introduce a more robust and reliable mechanism for pagination to efficiently handle large datasets and improve the overall performance of the API.

### 4. Use Migration Tools
- Integrate database migration tools to manage schema changes and versioning, ensuring consistent and reliable database updates across different environments. For example, using tools like Goose or Golang Migrate.
Currently, the project loads the schema and data using the `init.sql` file in the `postgresql` directory.

### 5. Use Structured Logging
- Integrate a structured logging library to capture detailed application logs, enabling developers to analyze and troubleshoot issues more effectively.

### 6. Support CRUD Operations
- Add support for updating and deleting transactions, enabling users to modify or remove existing records as needed.

### 7. Authentication and Authorization
- Implement user authentication and authorization mechanisms to secure the API endpoints, ensuring that only authorized users can access sensitive data and perform specific actions.

### 8. Enhance Validation and Error Handling
- Improve the validation and error handling mechanisms to provide more detailed and user-friendly responses, helping users understand and resolve issues more effectively.

### 9. Logging and Monitoring
- Integrate logging and monitoring tools to track application behavior and performance, enabling developers to identify and address potential issues in real-time. For example, instrumenting the application with OpenTelemetry to capture telemetry data and trace requests and generate custom metrics.

### 10. Caching and Performance Optimization
- Implement caching mechanisms and performance optimization strategies to improve the API's responsiveness and reduce latency, enhancing the overall user experience.

### 11. Testing and CI/CD Automation
- Integrate continuous integration and deployment pipelines to automate the development workflow, ensuring consistent code quality and reliable deployments.

### 11. Container Orchestration
- Explore container orchestration platforms like Kubernetes for deploying and managing the application in a production environment, providing scalability, resilience, and easier management of resources.

These improvements will help make the application more robust, secure, and scalable, making it ready for production use in demanding environments.
