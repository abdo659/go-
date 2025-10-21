# Go Docker Demo API

A simple REST API built with Go and containerized with Docker. This project demonstrates how to dockerize a Go application and deploy it to Docker Hub.

## 🚀 Features

- **RESTful API** with CRUD operations for users
- **Health check endpoint** for monitoring
- **Dockerized** with multi-stage build for optimal image size
- **Security-focused** with non-root user
- **Production-ready** with health checks and proper error handling

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | API documentation and usage examples |
| GET | `/health` | Health check endpoint |
| GET | `/api/users` | Get all users |
| POST | `/api/users` | Create a new user |
| GET | `/api/users/{id}` | Get user by ID |
| DELETE | `/api/users/{id}` | Delete user by ID |

## 🐳 Docker Hub

**Docker Image:** `abdo659/go-docker-demo:latest`

[![Docker Hub](https://img.shields.io/badge/Docker%20Hub-go--docker--demo-blue?style=for-the-badge&logo=docker)](https://hub.docker.com/r/abdo659/go-docker-demo)

### Pull and Run

```bash
# Pull the image
docker pull abdo659/go-docker-demo:latest

# Run the container
docker run -p 8080:8080 abdo659/go-docker-demo:latest
```

## 🛠️ Local Development

### Prerequisites

- Go 1.21 or later
- Docker (optional)

### Running Locally

1. **Clone the repository:**
   ```bash
   git clone <your-repo-url>
   cd go-docker-demo
   ```

2. **Run with Go:**
   ```bash
   go run main.go
   ```

3. **Run with Docker:**
   ```bash
   # Build the image
   docker build -t go-docker-demo .
   
   # Run the container
   docker run -p 8080:8080 go-docker-demo
   ```

### Testing the API

Once the server is running, you can test the endpoints:

```bash
# Health check
curl http://localhost:8080/health

# Get all users
curl http://localhost:8080/api/users

# Create a new user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com"}'

# Get user by ID
curl http://localhost:8080/api/users/1

# Delete user
curl -X DELETE http://localhost:8080/api/users/1
```

## 🏗️ Docker Build Process

This project uses a multi-stage Docker build for optimization:

1. **Build Stage:** Uses `golang:1.21-alpine` to compile the Go application
2. **Final Stage:** Uses `alpine:latest` for a minimal runtime image

### Key Features:

- **Multi-stage build** reduces final image size
- **Non-root user** for security
- **Health checks** for container monitoring
- **Optimized layers** for better caching

## 📦 Building and Pushing to Docker Hub

1. **Build the image:**
   ```bash
   docker build -t abdo659/go-docker-demo:latest .
   ```

2. **Test locally:**
   ```bash
   docker run -p 8080:8080 abdo659/go-docker-demo:latest
   ```

3. **Login to Docker Hub:**
   ```bash
   docker login
   ```

4. **Push to Docker Hub:**
   ```bash
   docker push abdo659/go-docker-demo:latest
   ```

## 🔧 Configuration

The application can be configured using environment variables:

- `PORT`: Server port (default: 8080)

Example:
```bash
docker run -p 8080:3000 -e PORT=3000 abdo659/go-docker-demo:latest
```

## 📊 Response Format

All API responses follow a consistent format:

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {
    // Response data here
  }
}
```

## 🧪 Sample Data

The application comes with sample users pre-loaded:

- John Doe (john@example.com)
- Jane Smith (jane@example.com)

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📞 Support

If you have any questions or issues, please open an issue on GitHub.

---

**Docker Hub Link:** https://hub.docker.com/r/abdo659/go-docker-demo
