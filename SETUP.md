# Setup Guide

This guide will help you set up the required tools to build and run this Go Docker application.

## Prerequisites

### 1. Install Go (Optional - for local development)

**Windows:**
1. Download Go from https://golang.org/dl/
2. Run the installer and follow the instructions
3. Verify installation: `go version`

**macOS:**
```bash
# Using Homebrew
brew install go

# Or download from https://golang.org/dl/
```

**Linux:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Or download from https://golang.org/dl/
```

### 2. Install Docker (Required)

**Windows:**
1. Download Docker Desktop from https://www.docker.com/products/docker-desktop/
2. Install and start Docker Desktop
3. Verify installation: `docker --version`

**macOS:**
1. Download Docker Desktop from https://www.docker.com/products/docker-desktop/
2. Install and start Docker Desktop
3. Verify installation: `docker --version`

**Linux:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install docker.io
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER

# Log out and back in, then verify
docker --version
```

## Quick Start

Once Docker is installed, you can build and run the application:

### 1. Build the Docker Image

```bash
# Replace 'yourusername' with your Docker Hub username
docker build -t yourusername/go-docker-demo:latest .
```

### 2. Run the Container

```bash
docker run -p 8080:8080 yourusername/go-docker-demo:latest
```

### 3. Test the Application

Open your browser and go to:
- http://localhost:8080 - API documentation
- http://localhost:8080/health - Health check

Or use curl:
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/users
```

## Docker Hub Setup

### 1. Create Docker Hub Account

1. Go to https://hub.docker.com/
2. Sign up for a free account
3. Verify your email address

### 2. Login to Docker Hub

```bash
docker login
# Enter your Docker Hub username and password
```

### 3. Build and Push

```bash
# Build the image (replace 'yourusername' with your actual username)
docker build -t yourusername/go-docker-demo:latest .

# Push to Docker Hub
docker push yourusername/go-docker-demo:latest
```

### 4. Make Image Public

1. Go to your Docker Hub repository
2. Click on "Settings" tab
3. Change visibility to "Public"

## Troubleshooting

### Docker Issues

**"Docker is not running"**
- Start Docker Desktop (Windows/macOS)
- Start Docker service: `sudo systemctl start docker` (Linux)

**"Permission denied"**
- Add your user to docker group: `sudo usermod -aG docker $USER`
- Log out and back in

**"Port already in use"**
- Use a different port: `docker run -p 3000:8080 yourusername/go-docker-demo:latest`

### Go Issues

**"go: command not found"**
- Install Go from https://golang.org/dl/
- Add Go to your PATH environment variable

## Next Steps

1. **Test locally** - Make sure the application works
2. **Push to Docker Hub** - Upload your image
3. **Create GitHub repository** - Upload your code
4. **Update README.md** - Replace 'yourusername' with your actual Docker Hub username
5. **Submit** - Share your GitHub repository link

## Useful Commands

```bash
# Build image
docker build -t yourusername/go-docker-demo:latest .

# Run container
docker run -p 8080:8080 yourusername/go-docker-demo:latest

# Run in background
docker run -d -p 8080:8080 --name go-demo yourusername/go-docker-demo:latest

# View logs
docker logs go-demo

# Stop container
docker stop go-demo

# Remove container
docker rm go-demo

# List images
docker images

# Remove image
docker rmi yourusername/go-docker-demo:latest
```
