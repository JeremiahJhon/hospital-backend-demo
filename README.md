🏥 Hospital Middleware API

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

---RUN DOCKER---
docker-compose up -d

📄 License
This project was created as part of a technical assessment.
