# Financfy-Server

Welcome to the **Financfy-Server**, the backend server for managing accounts

## Table of Contents

- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [API Routes](#api-routes)
- [Server Details](#server-details)

---

## Tech Stack

- **Go**: Backend server
- **Go-chi**: HTTP router
- **PostgreSQL**: Database
- **Structure**: Hexagonal Architecture(DDD)

### Go Packages Used

- **Go-chi**: `github.com/go-chi/chi/v5`
- **godotenv**: `github.com/joho/godotenv`
- **pgx**: `github.com/jackc/pgx/v5`
- **uuid**: `github.com/google/uuid`

---

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/adibaruet/financfy-backend.git
   cd financfy-backend
   ```

2. **Run the server**:
   ```bash
   docker-compose up --build -d
   ```

---

## API Routes

### Auth Routes

- **POST /api/v1/login**
    - Description: Log in a user with email and password.
    - Request Body:
        ```json
        {
            "email": "user@example.com",
            "password": "password123"
        }
        ```
  - Response Body:
    ```json
    {
      "user_id": "723a5e27-483e-453d-bab2-3f00eafb6688"
    }
    ```        

- **POST /api/v1/account**

  - Description: Register a new user. Must Include .ac.bd in Email
  - Request Body:
    ```json
    {
      "name": "John Doe",
      "email": "user@student.ruet.ac.bd",
      "password_hash": "password123",
    }
    ```
  - Response Body:
    ```json
    {
      "user_id": "723a5e27-483e-453d-bab2-3f00eafb6688"
    }
    ```
---

## Server Details

- **Host**: `localhost`
- **Port**: `8080`
- **Base URL**: `http://localhost:8080`

---
