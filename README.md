## 📚 Bookstore Microservices (Go + Echo + PostgreSQL)

This project demonstrates a basic microservices architecture using:

- **Go (Golang)**
- **Echo web framework**
- **GORM (ORM)**
- **PostgreSQL**
- **JWT authentication**
- **Docker + Docker Compose**

---

### 🔧 Services

### 🧑 `user-service`

Handles:
- User registration
- User login with JWT token
- JWT-protected user data access

### 📖 `book-service`

Handles:
- Book creation (authenticated users only)
- Book listing

---

### 🚀 Getting Started

### ✅ Prerequisites

- Docker + Docker Compose
- Go 1.21+ (for local dev)

---

### 🛠 Build & Run

```bash
make dev
