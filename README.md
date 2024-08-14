# Task Management

A task management web application built with Go, using the Gin framework for HTTP routing and middleware, JWT for authentication, and secure password hashing. This project follows a clean, modular architecture to ensure maintainability and scalability.

## Features

- **Task Management**: Create, read, update, and delete tasks.
- **User Authentication**: Register and log in users using JWT.
- **Secure Password Handling**: Passwords are hashed and compared securely.
- **Unit Testing**: Comprehensive tests for each component and feature.

## Folder Structure

```plaintext
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   ├── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
├── Usecases/
│   ├── task_usecases.go
│   └── user_usecases.go
└── tests/
    ├── task_controller_test.go         // Unit tests for task controller
    ├── user_controller_test.go         // Unit tests for user controller
    ├── jwt_service_test.go             // Unit tests for JWT service
    ├── password_service_test.go        // Unit tests for password service
    ├── task_usecases_test.go            // Unit tests for task use cases
    └── user_usecases_test.go            // Unit tests for user use cases
```

## Getting Started

### Installation

1. **Clone the Repository:**

   ```sh
   git clone [https://github.com/Simret101/Clean_Architecute_for_TaskManagement]
   cd task-manager
   ```

2. **Install Dependencies:**

   ```sh
   go mod tidy
   ```

3. **Run the Application:**

   ```sh
   go run Delivery/main.go
   ```

   The application will start on `http://localhost:8080`.

### Running Tests

To run tests for the entire project:

```sh
go test ./... -v
```

### Endpoints

- **POST /register**: Register a new user.
- **POST /login**: Log in a user and receive a JWT token.
- **GET /tasks**: Retrieve all tasks.
- **POST /tasks**: Create a new task.
- **GET /tasks/{id}**: Retrieve a task by ID.
- **PUT /tasks/{id}**: Update a task by ID.
- **DELETE /tasks/{id}**: Delete a task by ID.

## Contact

For any questions or issues, please open an issue or contact us at [semret.b74@gmail.com]
