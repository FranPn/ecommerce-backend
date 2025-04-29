# Ecommerce Backend (Golang + PostgreSQL)

This is a simple backend for an ecommerce system built with:
- Golang (1.24.2)
- Gin web framework
- GORM for ORM (PostgreSQL)
- JWT authentication
- Docker ready (optional)

## 📦 Features
- User registration (`/register`)
- User login (`/login`)
- JWT-based authentication
- Protected routes (`/api/profile`)
- Passwords are hashed securely with bcrypt
- Ready to run locally or with Docker Compose

---

## 🚀 Project Structure
```
/cmd/server/             # Application entrypoint
/internal/
    /controllers/         # HTTP handlers
    /models/              # Database models
    /services/            # Business logic
    /middleware/          # JWT authentication middleware
    /utils/               # Utility functions (JWT, hashing)
    /routes/              # API routing
/pkg/db/                  # Database connection
Makefile                  # Automation commands
Dockerfile                # Build instructions for Go app (optional)
docker-compose.yml        # Local development with Docker (optional)
.env                      # Environment variables
.gitignore                # Git ignore rules
README.md                 # Project documentation
```
---

## 🛠️ How to run locally (without Docker)

1. **Install Go 1.24.2+ and PostgreSQL 17+**

2. **Create a PostgreSQL database**

```bash
psql -U postgres
CREATE DATABASE ecommerce_db;
