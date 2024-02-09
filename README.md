# Transaction API

## Project Overwiew

This project was developed as part of the technical challenge for the Traive Engineering interview. The goal is to build a RESTful API in Go (Golang), allowing the management of user transactions. The API will enable the creation and listing of transactions with features like pagination and filtering.
The project has been time-boxed to a maximum of 4 hours in order to showcase the candidate's ability to deliver an application within a limited timeframe.

## Design Assumptions

- **Framework**: The project uses the Gin web framework for its lightweight nature and efficient performance in building RESTful APIs.
- **Database**: PostgreSQL is chosen for its reliability and feature-rich support for transactional data management.
- **Transactions**: Each transaction includes details such as ID, origin, user ID, amount, operation type (credit/debit), and creation timestamp. Additional attributes have not been considered at this stage.

## Technical Challenge Requirements

- Develop an API to manage transactions with capabilities to create and list transactions.
- Ensure the API supports pagination and filtering for listing transactions.
- Aim for a production-ready application with considerations for project structure, code organization, and technology choices.
- Incorporate error handling and robust application behavior analysis strategies.
- Utilize a consistent naming convention and maintain code readability.

## Project Structure

The project follows a standard Go project layout with separate directories for the application, database, and Docker configurations. The main components of the project structure are as follows:

- **`/api`**: Contains the API handlers and routes for managing transactions.
- **`/config`**: Includes the application configuration settings and environment variables.
- **`/repository`**: Contains the database schema and migration scripts for setting up the PostgreSQL database.
- **`/models`**: Includes the data models and database access layer for managing transactions.
- **`/utils`**: Contains utility functions and helper methods used across the application.
- **`main.go`**: The entry point of the application that initializes the server and routes.
- **`Dockerfile`**: The Docker configuration file for building the application image.
- **`docker-compose.yml`**: The Docker Compose configuration for setting up the PostgreSQL database.
- ** `postgresql`**: The directory for the init script for the PostgreSQL database.
- **`go.mod` and `go.sum`**: The Go module files for managing dependencies.
- **`Makefile`**: The Makefile for defining common tasks and commands for the application.
- **`README.md`**: The project documentation and instructions for running the application.

Some References for the project structure:
- [Go Project Layout](https://go.dev/doc/code#Organization)
- [Applied Go - Project Layout](https://appliedgo.com/blog/go-project-layout)

## Prerequisites
Pre-requisites to run the application are as follow: 

- Go (version 1.x)
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

### Stopping the Database

To stop the PostgreSQL database container, run the following command:

```sh
docker-compose down
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

To run the automated tests for the application, execute the following command:

```sh
make test
```

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
To customize, you can modify the value directly in docker-compose.yml or use a .env file with Docker Compose to define DATABASE_URL.

### Interacting with Endpoints

Swagger provides an interactive interface where you can:

- View all available endpoints
- See details about each endpoint, including request and response formats
- Test endpoints directly from the Swagger UI by providing input parameters and sending requests
- View sample responses and error messages

This documentation is useful for developers who want to explore and test the API without having to refer to the codebase or external documentation.



## Future Improvements
## Future Improvements for Production Readiness

### 1. Improve Testing Strategy
- Enhance Unit Tests and introduce Integration Tests to ensure comprehensive test coverage and identify potential issues across the application's components.

### 2. Enhance Filtering Capabilities
- Expand Filtering capabilities to include additional fields, providing users with more flexibility and options when querying transactions.

### 3. Implement Robust Pagination Mechanism
- Introduce a more robust and reliable mechanism for pagination to efficiently handle large datasets and improve the overall performance of the API.

### 4. Support CRUD Operations
- Add support for updating and deleting transactions, enabling users to modify or remove existing records as needed.

### 5. Authentication and Authorization
- Implement user authentication and authorization mechanisms to secure the API endpoints, ensuring that only authorized users can access sensitive data and perform specific actions.

### 6. Enhance Validation and Error Handling
- Improve the validation and error handling mechanisms to provide more detailed and user-friendly responses, helping users understand and resolve issues more effectively.

### 7. Logging and Monitoring
- Integrate logging and monitoring tools to track application behavior and performance, enabling developers to identify and address potential issues in real-time.

### 8. Caching and Performance Optimization
- Implement caching mechanisms and performance optimization strategies to improve the API's responsiveness and reduce latency, enhancing the overall user experience.

### 9. Testing and CI/CD Automation
- Expand the test suite and integrate continuous integration and deployment pipelines to automate the development workflow, ensuring consistent code quality and reliable deployments.

### 10. Container Orchestration
- Explore container orchestration platforms like Kubernetes for deploying and managing the application in a production environment, providing scalability, resilience, and easier management of resources.

These improvements will help make the application more robust, secure, and scalable, making it ready for production use in demanding environments.
