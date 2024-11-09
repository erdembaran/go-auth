# Go Authentication API

A robust authentication API built with Go, Fiber, and MongoDB that provides comprehensive user authentication functionality.

## Features

- 🔐 User Authentication (Register/Login/Logout)
- 🔑 JWT-based Authentication
- 📧 Password Reset via Email
- 👤 User Management
- 🔒 Secure Password Hashing
- 🍪 HTTP-Only Cookie Support
- 📝 MongoDB Integration

## Prerequisites

Before running this project, make sure you have the following installed:

- Go (1.16 or later)
- MongoDB
- SMTP Server (for email functionality)

## Installation

1. Clone the repository

```
bash
git clone https://github.com/erdembaran/go-auth.git
cd go-auth

```

2. Install dependencies

```
bash
go mod download
```

3. Set up environment variables (create a `.env` file)

```
PORT=4000
MONGO_URI=YOUR_MONGO_DB_URI
ENV=development
JWT_SECRET=YOUR_JWT_SECRET
SMTP_HOST=smtp.example.com
SMTP_PORT=smtp_port
SMTP_USER=your-email@example.com
SMTP_PASS=your-email-password
```

## API Endpoints

### Authentication Routes

```
POST /api/v1/auth/register - Register a new user
POST /api/v1/auth/login - Login user
POST /api/v1/auth/logout - Logout user
POST /api/v1/auth/forgot-password - Request password reset
PUT /api/v1/auth/reset-password/:token - Reset password
```

### User Routes (Protected)

```
GET /api/v1/users - Get all users
GET /api/v1/users/:id - Get specific user
```

```
bash
go run main.go
```

2. The API will be available at `http://localhost:3000`

## Security Features

- Password Hashing using bcrypt
- JWT Authentication
- HTTP-Only Cookies
- Secure Password Reset Flow
- Protected Routes with Middleware

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments

- [Fiber](https://github.com/gofiber/fiber)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [JWT Go](https://github.com/golang-jwt/jwt)
