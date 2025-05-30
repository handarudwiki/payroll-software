# Payroll Software

A comprehensive payroll management system built with Go, designed to handle employee salary calculations, deductions, and payroll processing efficiently.

## Table of Contents

- [Payroll Software](#payroll-software)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Architecture](#architecture)
    - [Why This Architecture?](#why-this-architecture)
    - [Architecture Layers](#architecture-layers)
  - [Project Structure](#project-structure)
    - [Directory Explanations](#directory-explanations)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Database Setup](#database-setup)
    - [Running Migrations](#running-migrations)
    - [Database Seeding](#database-seeding)
  - [Usage](#usage)
    - [Running the Application](#running-the-application)
    - [Available Make Commands](#available-make-commands)
  - [Testing](#testing)
  - [Database Schema](#database-schema)
  - [Contributing](#contributing)
  - [License](#license)
  - [Support](#support)

## Features

- Employee management
- Salary calculation and processing
- Tax and deduction management
- Payroll report generation
- User authentication and authorization
- RESTful API endpoints
- Database migration support
- Comprehensive testing suite

## Architecture

This project follows a **Clean Architecture** pattern with clear separation of concerns. The architecture choice is based on several key principles:

### Why This Architecture?

1. **Familiarity and Experience**: This architecture pattern has been proven effective in previous projects, ensuring faster development and maintenance.

2. **Separation of Concerns**: Each layer has a specific responsibility, making the codebase more maintainable and testable.

3. **Scalability**: The modular structure allows easy addition of new features without affecting existing functionality.

4. **Testability**: Clear boundaries between layers make unit testing and integration testing straightforward.

5. **Dependency Inversion**: Business logic doesn't depend on external frameworks or databases, making the system more flexible.

### Architecture Layers

```
┌─────────────────────────────────────────┐
│                Controllers              │  ← HTTP Handlers
├─────────────────────────────────────────┤
│                Services                 │  ← Business Logic
├─────────────────────────────────────────┤
│              Repositories               │  ← Data Access Layer
├─────────────────────────────────────────┤
│                Database                 │  ← Data Storage
└─────────────────────────────────────────┘
```

## Project Structure

```
payroll-software/
├── cmd/                    # Application entry points
│   ├── main.go            # Main application
│   └── seed/              # Database seeding
│       └── seed.go
├── internal/              # Private application code
│   ├── controllers/       # HTTP request handlers
│   ├── models/           # Data models and entities
│   ├── dto/              # Data Transfer Objects
│   ├── services/         # Business logic layer
│   ├── repositories/     # Data access layer
│   ├── utils/            # Utility functions
│   └── response/         # API response structures
├── database/
│   └── migrations/       # Database migration files
├── tests/
│   └── e2e/             # End-to-end tests
├── .env                 # Environment variables
├── .env.example         # Environment variables template
├── Makefile            # Build and deployment commands
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksums
└── README.md           # Project documentation
```

### Directory Explanations

- **cmd/**: Contains the main applications for this project. The main.go file is the entry point of the application, while seed/ contains database seeding utilities.

- **internal/**: Houses the private application code that shouldn't be imported by other applications.
  - **controllers/**: HTTP handlers that process incoming requests and return responses
  - **models/**: Domain entities and data structures
  - **dto/**: Data Transfer Objects for API communication
  - **services/**: Business logic and core application functionality
  - **repositories/**: Data access layer that interacts with the database
  - **utils/**: Shared utility functions and helpers
  - **response/**: Standardized API response structures

## Prerequisites

- Go 1.19 or higher
- PostgreSQL 12 or higher
- Goose (for database migrations)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd payroll-software
```

2. Install dependencies:
```bash
go mod download
```

3. Install Goose for database migrations:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Configuration

1. Copy the environment variables template:
```bash
cp .env.example .env
```

2. Update the `.env` file with your database configuration:
```env
DATABASE_URL=postgres://username:password@localhost:5432/payroll_db?sslmode=disable
DATABASE_URL_TEST=postgres://username:password@localhost:5432/payroll_test_db?sslmode=disable
```

## Database Setup

### Running Migrations

The project uses Goose for database migrations. Use the following Makefile commands:

```bash
# Run all pending migrations
make up

# Run migrations for test database
make up_test

# Rollback the last migration
make down

# Create a new migration file
make new name=create_employees_table
```

### Database Seeding

Populate the database with initial data:

```bash
make seed
```

## Usage

### Running the Application

Start the development server:

```bash
make run
```

The application will start on the configured port (default: 8080).

### Available Make Commands

- `make run` - Start the application
- `make test` - Run end-to-end tests
- `make up` - Run database migrations
- `make up_test` - Run migrations for test database
- `make down` - Rollback last migration
- `make new name=<migration_name>` - Create new migration
- `make seed` - Seed database with initial data


## Testing

Run the test suite:

```bash
make test
```

The project includes:
- Unit tests for individual components
- Integration tests for API endpoints
- End-to-end tests for complete workflows

## Database Schema

The application uses PostgreSQL with the following main entities:
- Users (authentication)
- Employees (employee information)
- Payroll (salary calculations)
- Deductions (tax and other deductions)
- Reports (payroll reports)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions, please open an issue in the GitHub repository.