# CRM-Monolith-Service

## 📋 Overview

This project is a **monolithic web service** written in **Golang 1.22**, designed with a **layered architecture** to ensure clear separation of concerns and maintainability. The project uses the following technologies and libraries:

- **Server**: Built on top of the `net/http` package for both serving HTTP requests and routing.
- **Dependency Injection**: Managed using [Fx](https://github.com/uber-go/fx), a powerful dependency injection framework for Go.
- **Database**: Interacts with a PostgreSQL database via [GORM](https://gorm.io/index.html) as the ORM (Object-Relational Mapping) library.
- **Migrations**: Database schema migrations are handled using [Goose](https://github.com/pressly/goose).
- **Authorization**: User authorization is implemented using JWT (JSON Web Tokens).
- **Configuration**: Application configuration is managed with [Viper](https://github.com/spf13/viper).

## 🏗️ Project Structure

The project follows a **layered architecture**, divided into the following components:

- **Handlers**: Handle HTTP requests and responses. They are the entry point for incoming requests and communicate with services to perform business logic.
- **Services**: Contain the core business logic of the application. Services interact with repositories to fetch or persist data and perform operations that may span multiple repositories.
- **Repositories**: Responsible for direct interaction with the database. They perform CRUD (Create, Read, Update, Delete) operations using GORM.

## 🚀 Setup and Installation

### Prerequisites

- Go 1.22
- Docker and Docker Compose
- PostgreSQL

