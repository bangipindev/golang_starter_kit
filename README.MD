# GPT Auth Service

Production-ready REST API menggunakan:

- Golang
- Fiber v2
- MySQL
- JWT Authentication
- Clean Architecture
- Goose Migration
- Docker Support

---

# 📦 Project Structure
```
	cmd/
	└─server/
	│ └── main.go             # Entry point aplikasi (main.go)
	│
	├─config/
	│ └── config.go				# Konfigurasi aplikasi (env, database, dll)
	│
	├──internal/
	│ └──auth/
	│ 	└── token/				# JWT, hashing, auth service
	│
	├── container/
	│ └── container.go			# Dependency injection container
	│
	├── delivery/				# Layer HTTP (handler, routes, dto)
	│ └── http/
	│ ├── dto/
	│ ├── handler/
	│ └── routes.go
	│
	├── domain/					# Entity & interface (business core contract)
	│ └── user.go
	│
	├── infrastructure/			# Implementasi repository (DB, external service)
	│ ├── repository/
	│ ├── database.go
	│ └── migrate.go
	│
	├── middleware/				# HTTP middleware (Auth, logging, dll)
	│ └── jwt_middleware.go
	│
	├── pkg/					# Shared utilities (response, helper, dll)
	│ └── response/
	│
	├── usecase/				# Business logic
	  └── auth.go


	migrations/					# Database migration files
	└── 001_create_user.sql
	├── .air.toml           	# Konfigurasi hot reload (Air)
	├── .env                	# Environment variables
	├── .env.example        	# Contoh konfigurasi env
	├── docker-compose.yml  	# Docker compose config
	├── Dockerfile         	 	# Docker image config
	├── go.mod              	# Go module definition
	├── go.sum              	# Dependency checksum
	└── README.md           	# Dokumentasi project

```



---

# 🧠 Architecture Overview
```
	Project ini menggunakan Clean Architecture:

	HTTP (Fiber)
	↓
	Handler (Delivery Layer)
	↓
	Usecase (Business Logic)
	↓
	Domain (Interface + Entity)
	↓
	Infrastructure (DB, Repository, Migration)
```


## 📌 Layer Explanation

### 1️⃣ delivery/http
Berisi:
- DTO (request/response struct)
- Handler (controller)
- Route definition

Tidak boleh berisi business logic.

---

### 2️⃣ domain
Berisi:
- Entity (User struct)
- Interface repository

Pure business model.
Tidak boleh import Fiber / MySQL.

---

### 3️⃣ usecase ( Bisa juga disebut Layer Service)
Berisi:
- Business logic
- Auth logic
- JWT creation
- Password hashing

Tidak tahu HTTP.
Hanya tahu interface repository.

---

### 4️⃣ infrastructure
Berisi:
- Database connection
- Repository implementation ( Penulisan Query Ada Disini)
- Migration runner

Layer ini implementasi teknis.

---

### 5️⃣ middleware
JWT middleware untuk proteksi endpoint.

---

### 6️⃣ pkg/response
Standard API response wrapper agar konsisten.

---

# 🚀 Features

- Register
- Login
- JWT Authentication
- Protected Route
- Auto Migration
- Standard JSON Response
- Environment Based Config
- APP_MODE (dev / prod)

---

# ⚙️ Environment Setup

Buat file `.env`:

```
	DB_HOST=localhost
	DB_PORT=3306
	DB_USER=root
	DB_PASS=password
	DB_NAME=gpt_db

	JWT_SECRET=supersecretkey

	APP_PORT=8080
	APP_MODE=dev

```

### INSTALL AIR 
```
go install github.com/air-verse/air@latest

```
---

# ▶️ Run Without Docker

1️⃣ Install dependencies

```
go mod tidy

```

Copy env.example to .env

```
 Copy env.example to .env

```


2️⃣ Pastikan MySQL jalan

3️⃣ Jalankan app

Manual Run

```
go run cmd/server/main.go

```

Hot Reload / Automatic Run

```
air

```

---

# 🐳 Run With Docker

```
docker-compose up --build

```


---

# 🗄 Migration

Migration menggunakan Goose.

File migration ada di:

```
migrations/001_create_user.sql

```


Auto run migration saat startup:

```
infrastructure.RunMigration(db)

```


Manual run:

```
goose -dir migrations mysql "user:pass@tcp(localhost:3306)/gpt_db" up

```



---

# 🔐 API Endpoints

## Register

POST `/api/auth/register`

```
	{
		"name": "Ervin",
		"email": "ervin@mail.com",
		"password": "123456"
	}
```


---

## Login

POST `/api/auth/login`

Response:

```
	{
		"success": true,
		"data": {
			"token": "JWT_TOKEN"
		}
	}
```


---

## Profile (Protected)

GET `/api/profile`

Header:


```
	Authorization: Bearer JWT_TOKEN

```


---

# 🏗 Clean Architecture Rules

✔ Domain tidak boleh import Fiber
✔ Usecase tidak boleh import database
✔ Handler tidak boleh akses database langsung
✔ Repository implementasi ada di infrastructure
✔ Semua dependency di inject dari main.go

---
### FULL FLOW VISUAL
```
	Client
	↓
	Fiber Router
	↓
	Handler (delivery/http)
	↓
	Usecase (business logic)
	↓
	Domain (interface)
	↓
	Repository (infrastructure)
	↓
	Database (MySQL)
	↑
	Repository return result
	↑
	Usecase process
	↑
	Handler format response
	↑
	Client menerima JSON

```
---

### 🔐 FLOW PROTECTED ROUTE (JWT)
```
	Client
	↓
	JWT Middleware
	↓ (valid token?)
	Handler.Profile
	↓
	Usecase.GetProfile
	↓
	Repository.FindByID
	↓
	Database
	↑
	Return user data
	↑
	Response JSON


```

Jika token invalid:

```
	Middleware stop
	Return 401

```
---

### 📊 Versi Diagram Arsitektur Clean (Conceptual)

```
	┌────────────────────┐
	│      DELIVERY      │
	│  (Fiber, Handler)  │
	└─────────┬──────────┘
			  ↓
	┌────────────────────┐
	│      USECASE       │
	│  (Auth Service)    │
	└─────────┬──────────┘
			  ↓
	┌────────────────────┐
	│       DOMAIN       │
	│ Entity + Interface │
	└─────────┬──────────┘
			  ↓
	┌────────────────────┐
	│   INFRASTRUCTURE   │
	│   DB + Repository  │
	└────────────────────┘

```
---

### CREATE NEW FILE MIGRATIONS
```
 migrate create -ext sql -dir migrations -seq create_new_file

```
# 📈 Production Ready Improvements

- Add refresh token
- Add structured logging (zap)
- Add request ID middleware
- Add validation (go-playground/validator)
- Add rate limiter
- Add role based access
- Add Swagger docs
- Add unit test
- Add CI/CD pipeline

---

# 👨‍💻 Author

Ervin Santoso
Backend Developer


