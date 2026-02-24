# 🚀 Go Minimal Backend Project

## 📌 Overview

`go-minimal` is a lightweight backend project built using **Go (Golang)** following clean architecture principles.

The project is structured using:

- Handler Layer (HTTP Layer)
- Service Layer (Business Logic Layer)
- Repository Layer (Database Access Layer)
- PostgreSQL Database Integration
- Dependency Injection Pattern
- Middleware Logging System

This project demonstrates modern backend engineering practices using Go.

---

## 🏗 Project Architecture
```
cmd/
├── server/
├── worker/
internal/
├── config/
├── handler/
├── middleware/
├── model/
├── repository/
└── service/
```

### Layer Flow

```
HTTP Request
↓
Middleware
↓
Handler
↓
Service
↓
Repository
↓
PostgreSQL Database
```


---

## ⚡ Technology Stack

- Go (Golang)
- PostgreSQL
- pgx Driver
- Air (Hot Reload Development Tool)
- Docker (Optional Database Setup)

---

## 📦 Prerequisites

Make sure you have installed:

- Go (version 1.23+ recommended)
- Docker (if using PostgreSQL via container)
- Git

---

## 🔧 Environment Setup

Create a `.env` file in project root:

```

DATABASE_URL=postgres://admin:admin123@localhost:5432/go_db
PORT=8082

````

---

## 🐳 Database Setup (Docker PostgreSQL - Recommended)

### Stop system PostgreSQL if running

```bash
sudo systemctl stop postgresql
````

---

### Run PostgreSQL Container

```bash
docker run -d \
--name go-postgres \
-e POSTGRES_USER=admin \
-e POSTGRES_PASSWORD=admin123 \
-e POSTGRES_DB=go_db \
-p 5432:5432 \
postgres:16
```

---

## 🗄 Database Schema

Run inside PostgreSQL:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NOT NULL
);
```

---

## 🚀 Installation & Running Project

### Clone Repository

```bash
git clone https://github.com/SarojGhalanCubit/go_learning.git
cd go-minimal
```

---

### Install Dependencies

```
go mod tidy
```

---

### Run Project (Development Mode)

If using Air hot reload:

```bash
air
```

---

### Run Manually

```bash
go run cmd/server/main.go
```

---

## 📌 API Endpoints

### Get Users

```
GET /users
```

Response:

```json
[
  {
    "id": 1,
    "name": "Saroj",
    "age": 25
  }
]
```

---

### Create User

```
POST /users/create
```

Request Body:

```json
{
  "name": "John",
  "age": 30
}
```

---

## 🔥 Development Tools Used

* Air — Hot reload development
* pgx — PostgreSQL driver
* Middleware logging system

---

## ⚠️ Important Notes

* Do not run system PostgreSQL and Docker PostgreSQL simultaneously.
* Ensure port `5432` is free before starting database container.
* Always use `.env` for sensitive configuration.

---

## 🤝 Contribution

Feel free to fork and improve the project.

Pull requests are welcome.

---

## ⭐ Author
Saroj Ghalan


