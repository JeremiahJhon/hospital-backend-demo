Hospital Middleware API

A secure Hospital Middleware system built with Go (Gin + GORM + PostgreSQL) that allows hospital staff to authenticate and search patient information through an external Hospital Information System (HIS) API.

📌 Overview

This project implements:

🔐 Staff authentication (JWT-based)

🏥 Multi-hospital data isolation

👨‍⚕️ Staff management

🧑 Patient search functionality

🔄 Integration with external HIS API

🗄️ PostgreSQL database

🧱 Clean architecture (handler → service → repository)

The system acts as a middleware layer between client applications and external hospital systems.

🚀 Tech Stack

Language: Go

Framework: Gin

ORM: GORM

Database: PostgreSQL

Authentication: JWT

Password Hashing: bcrypt

Environment Management: godotenv

Containerization: Docker (optional)

🏗️ Architecture
internal/
  config/        # Database, environment setup
  models/        # Database models
  repository/    # Data access layer
  service/       # Business logic
  handler/       # HTTP handlers
  middleware/    # JWT authentication
  client/        # External HIS API integration

Request flow:

Client Request
   ↓
JWT Middleware
   ↓
Handler
   ↓
Service
   ↓
Repository
   ↓
Database

If patient not found locally:

Service
   ↓
External HIS API
   ↓
Save to DB
   ↓
Return response
🔐 Authentication

Authentication uses JWT.

Protected endpoints require:

Authorization: Bearer <token>

JWT contains:

staff_id

hospital_id

🗄️ Database Models
Hospital

id (UUID)

name

created_at

Staff

id (UUID)

username (unique)

password_hash

hospital_id (FK)

created_at

Patient

id (UUID)

first_name_th

middle_name_th

last_name_th

first_name_en

middle_name_en

last_name_en

date_of_birth

patient_hn

national_id (indexed)

passport_id (indexed)

phone_number

email

gender

hospital_id (FK)

created_at

📡 API Endpoints

Base URL:

http://localhost:8080/api
👨‍⚕️ Staff
Create Staff

POST /staff/create

{
  "username": "john",
  "password": "securepassword",
  "hospital_id": "uuid"
}
Login

POST /staff/login

{
  "username": "john",
  "password": "securepassword",
  "hospital_id": "uuid"
}

Response:

{
  "token": "jwt_token_here"
}
🧑 Patient
Search Patient (Requires Login)

GET /patient/search

Query Parameters (optional):

national_id

passport_id

patient_hn

first_name_en

last_name_en

Example:

GET /patient/search?national_id=1234567890123
Authorization: Bearer <token>

Behavior:

Search local database

If not found → call external HIS API

Save result

Return patient data

🔒 Security Rules

Passwords are hashed using bcrypt

JWT required for protected routes

Staff can only access patients within the same hospital

Hospital isolation enforced using hospital_id filtering

⚙️ Setup Instructions
1️⃣ Clone Repository
git clone https://github.com/yourusername/hospital-middleware.git
cd hospital-middleware
2️⃣ Configure Environment

Create .env file:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hospital
JWT_SECRET=your_secret_key
HIS_BASE_URL=https://hospital-a.api.co.th
3️⃣ Run PostgreSQL (Docker)
docker-compose up -d
4️⃣ Run Application
go mod tidy
go run main.go

Server runs on:

http://localhost:8080
🧪 Testing

You can test using:

Postman

curl

Swagger (if implemented)

🧠 Design Decisions

Clean separation of concerns

Multi-tenant architecture via hospital_id

External API integration abstracted via client layer

Secure password storage

Proper database indexing for performance

📌 Future Improvements

Pagination support

Rate limiting

Structured logging

Unit and integration tests

OpenAPI (Swagger) documentation

Docker production setup

Role-based access control

📄 License

This project was created as part of a technical assessment.
