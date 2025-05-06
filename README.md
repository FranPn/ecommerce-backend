# Ecommerce Backend (Golang + PostgreSQL + Meilisearch)

This is a modular backend API for an ecommerce platform, built with:
- Golang 1.24.2
- PostgreSQL 17
- GORM for ORM (PostgreSQL)
- Gin web framework
- Meilisearch for full-text product search
- JWT-based authentication (admin/user)
- Docker-ready setup (optional)

## 📦 Features

- ✅ User registration and login with JWT
- ✅ Role-based access (admin vs user)
- ✅ Password hashing (bcrypt)
- ✅ CRUD for products
- ✅ Full-text product search with Meilisearch
- ✅ Protected profile route (`/api/profile`)
- ✅ Admin-only product creation/edit/delete
- ✅ Graceful concurrency handling with goroutines
- ✅ Ready to deploy or extend (e.g. orders, cart, frontend)

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
/pkg/
    /db/                  # Database connection
    /search/              # Meilisearch integration
    /utils/               # JWT utils, password hashing
Makefile                  # Automation commands
Dockerfile                # Build instructions for Go app (optional)
docker-compose.yml        # Local development with Docker (optional)
.env                      # Environment variables
.gitignore                # Git ignore rules
README.md                 # Project documentation
```
---

## 🛠️ How to run locally (without Docker)

1. **Install dependencies**

- [Go 1.24.2](https://golang.org/dl/)
- [PostgreSQL 17](https://www.postgresql.org/)
- [Meilisearch](https://www.meilisearch.com/) (`localhost:7700`)

2. **Create a PostgreSQL database**

```bash
psql -U postgres
CREATE DATABASE ecommerce_db;

```
## 💻 Technologies Used

- Go (Golang) 1.24.2
- Gin (Web framework)
- GORM (ORM)
- PostgreSQL 17
- bcrypt (Password hashing)
- JWT (Token-based authentication)
- Docker (optional)
- Makefile
- Git & GitHub
```
