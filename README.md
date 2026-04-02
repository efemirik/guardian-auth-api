# 🛡️ Guardian Auth API

A high-performance, security-focused Backend API built with **Go** to manage users, roles, and permissions. Designed with a cybersecurity mindset and Clean Architecture principles.

## ✨ Current Features (Implemented)
* **JWT-Based Authentication:** Stateless, secure token generation and validation.
* **Password Hashing:** Secure credential storage using `bcrypt`.
* **Custom Security Middleware:** Route protection preventing unauthorized access.
* **PostgreSQL & GORM:** Reliable data persistence with auto-migration.
* **Dockerized Infrastructure:** Fully containerized database for isolated local development.

## 🚀 Roadmap (Planned Features)
* **Advanced Hashing:** Migration from `bcrypt` to **Argon2id** (the winner of the Password Hashing Competition).
* **Token Management:** Implementation of Refresh Tokens.
* **Advanced RBAC:** Multi-level permissions (Admin, Moderator, User).
* **Caching & Blacklisting:** Redis integration for session management and token blacklisting.
* **Rate Limiting:** Protection against Brute-Force and common OWASP threats.

## 🛠️ Tech Stack
* **Language:** Go (Golang)
* **Framework:** Gin Web Framework
* **Database:** PostgreSQL
* **ORM:** GORM
* **Infrastructure:** Docker & Docker Compose

## 📦 Quick Start

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/efemirik/guardian-auth-api.git](https://github.com/efemirik/guardian-auth-api.git)
   cd guardian-auth-api