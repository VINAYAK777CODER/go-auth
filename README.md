# Go Auth API (Gin + GORM + JWT)

A simple authentication backend built using **Golang**, **Gin**, **GORM**, **JWT**, and **PostgreSQL**.

---

## 📦 Tech Stack

- **Go** – Programming language  
- **Gin** – Web framework  
- **GORM** – ORM for database  
- **PostgreSQL** – Database (Neon.tech)  
- **Bcrypt** – Password hashing  
- **JWT** – User authentication  
- **Godotenv** – Load .env files  
- **CompileDaemon** – Auto reload on file changes  

---

## 📁 Folder Structure

go-auth/
│
├── controllers/
│ └── userController.go
│
├── initializers/
│ ├── connectToDb.go
│ ├── loadEnvVariables.go
│ └── syncDatabase.go
│
├── middleware/
│ └── requireAuth.go
│
├── models/
│ └── usersModel.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── main.go



---

## ⚙️ Installation Guide

### 1️⃣ Clone the Repository

git clone https://github.com/VINAYAK777CODER/go-auth.git
cd go-auth



---

### 2️⃣ Initialize Go Module (if needed)

go mod init github.com/VINAYAK777CODER/go-auth



---

### 3️⃣ Install Dependencies

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv



---

### 4️⃣ Install Auto-Reload Tool (optional)

go install github.com/githubnemo/CompileDaemon@latest

csharp
Copy code

Run with:

CompileDaemon --build="go build -o go-auth.exe" --command="./go-auth.exe"



---

## 🌍 Environment Setup

Create a `.env` file:

PORT=8000
DB_URL=your_postgres_connection_string
SECRET=your_jwt_secret



---

## 🗄 Database Setup (Neon.tech)

1. Create a free PostgreSQL DB at https://neon.tech  
2. Copy your connection string  
3. Add it to `.env` as `DB_URL=`  
4. Run the project (GORM will auto-create tables)

---

## ▶️ Run the Server

go run main.go



Or using auto-reload:

CompileDaemon --command="./go-auth.exe"



---

## 🔐 API Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| POST | `/signup` | Create new user |
| POST | `/login` | Login & get JWT cookie |
| GET | `/validate` | Validate logged-in user |

---

## 📌 Notes

- Passwords are hashed before storing  
- JWT stored in HttpOnly cookie  
- Protected routes use `RequireAuth` middleware  

---

## 👨‍💻 Author

**VINAYAK777CODER**  
GitHub: https://github.com/VINAYAK777CODER

---
