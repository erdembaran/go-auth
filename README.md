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
- 🐳 Docker Support
- 🚀 Production-Ready Deployment

## Prerequisites

Before running this project, make sure you have the following installed:

- Go (1.22 or later)
- MongoDB
- Docker & Docker Compose (for containerized setup)
- SMTP Server (for email functionality)

## Installation

### Local Development

1. Clone the repository

```bash
git clone https://github.com/erdembaran/go-auth.git
cd go-auth
```

2. Install dependencies

```bash
go mod download
```

3. Set up environment variables (create a `.env.local` file)

```env
PORT=4000
MONGO_URI=mongodb://mongodb:27017/go-auth?directConnection=true
ENV=local
JWT_SECRET=your-jwt-secret
SMTP_HOST=smtp.example.com
SMTP_PORT=smtp_port
SMTP_USER=your-email@example.com
SMTP_PASS=your-email-password
```

### Using Docker

1. Make sure Docker and Docker Compose are installed
2. Run the application using Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:4000`

### Production Deployment (Render)

For production deployment on Render:

1. Create a new Web Service on Render
2. Choose "Docker Registry" as deployment type
3. Use the following Docker image: `baranerdem/go-auth:latest`
4. Set the following:
   - Instance Type: Choose based on your needs (Free/Standard)
   - Region: Choose closest to your users
   - Branch: main
5. Add the following environment variables in Render dashboard:
   ```
   PORT=4000
   MONGO_URI=your-mongodb-atlas-uri
   ENV=production
   JWT_SECRET=your-production-jwt-secret
   SMTP_HOST=your-smtp-host
   SMTP_PORT=your-smtp-port
   SMTP_USER=your-smtp-user
   SMTP_PASS=your-smtp-password
   ```
6. Click "Create Web Service"

The service will automatically pull the Docker image and deploy your application.

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

## Development

### Running Locally with Air (Hot Reload)

```bash
air
```

### Running with Docker

```bash
# Start the application and MongoDB
docker-compose up --build

# Stop the application
docker-compose down

# Remove volumes (if needed)
docker-compose down -v
```

## Security Features

- Password Hashing using bcrypt
- JWT Authentication
- HTTP-Only Cookies
- Secure Password Reset Flow
- Protected Routes with Middleware
- Environment-based Configuration
- Secure Production Setup

## Project Structure

```
.
├── config/         # Configuration and environment setup
├── controllers/    # Request handlers
├── database/      # Database connection and models
├── middleware/    # Custom middleware
├── models/        # Data models
├── routes/        # API routes
├── utils/         # Utility functions
├── .env.local     # Local environment variables
├── Dockerfile     # Docker configuration
└── docker-compose.yml  # Docker Compose configuration
```

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
- [Air](https://github.com/cosmtrek/air)

```

```
