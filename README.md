# 🛡️ Guardian Auth API (Go + JWT + RBAC)

A high-performance, security-focused Backend API built with **Go** to manage users, roles, and permissions. Designed with a cybersecurity mindset.

### 🚀 Planned Features
* **JWT-Based Authentication:** Secure login with Access and Refresh tokens.
* **Role-Based Access Control (RBAC):** Multi-level permissions (Admin, Moderator, User).
* **Secure Password Hashing:** Using **Argon2id** (the winner of the Password Hashing Competition).
* **Rate Limiting & Security Middlewares:** Protection against Brute-Force and common OWASP threats.
* **Structured Logging:** Tracking access attempts for security audits.

### 🛠️ Tech Stack
* **Language:** Go (Golang)
* **Framework:** Gin Gonic (for high-speed routing)
* **Database:** PostgreSQL (with GORM or sqlx)
* **Caching:** Redis (for session management and blacklisting)
* **Containerization:** Docker & Docker-compose
