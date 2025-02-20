# Go REST API with PostgreSQL - Book Management API

Demo test for  REST API built with Go, featuring user authentication, book management, and PostgreSQL database integration.

## 🚀 Features

- RESTful API endpoints
- PostgreSQL database with GORM
- JWT Authentication
- Swagger Documentation
- Docker containerization
- Automated database migrations
- Environment-based configuration

## 📋 Prerequisites

- Go 1.22+
- Docker and Docker Compose
- PostgreSQL 16
- Git

## 🛠 Quick Start

1. Clone the repository:
```bash
git clone git@github.com:ianemv/go-lang.git

2. cp .env.example .env
3. docker-compose --build to build the docker image
4. docker-compose up -d to start the docker container

The API will be available at http://localhost:8080

Swagger documentation is automatically generated and available at: http://localhost:8080/swagger/index.html

Key Endpoints
POST /api/auth/register - User registration
POST /api/auth/login - User authentication
GET /api/books - List all books
POST /api/books - Create new book
PUT /api/books/{id} - Update book
DELETE /api/books/{id} - Delete book