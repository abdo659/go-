# Assignment Submission Guide

## 📋 What You Have

This repository contains a complete Go Docker application with:

- ✅ **Go REST API** (`main.go`) - Simple user management API
- ✅ **Dockerfile** - Multi-stage build for optimal image size
- ✅ **Docker configuration** - `.dockerignore`, health checks, security
- ✅ **Documentation** - Comprehensive README.md and setup guides
- ✅ **CI/CD** - GitHub Actions workflow for automated builds
- ✅ **Helper scripts** - Docker commands script for easy management

## 🚀 Steps to Complete Assignment

### 1. Install Required Tools

Follow the `SETUP.md` guide to install:
- Docker Desktop
- Go (optional, for local development)

### 2. Test Locally

```bash
# Build the Docker image
docker build -t yourusername/go-docker-demo:latest .

# Run the container
docker run -p 8080:8080 yourusername/go-docker-demo:latest

# Test the API
curl http://localhost:8080/health
curl http://localhost:8080/api/users
```

### 3. Create Docker Hub Account

1. Go to https://hub.docker.com/
2. Sign up for a free account
3. Verify your email

### 4. Push to Docker Hub

```bash
# Login to Docker Hub
docker login

# Build and push (replace 'yourusername' with your actual username)
docker build -t yourusername/go-docker-demo:latest .
docker push yourusername/go-docker-demo:latest
```

### 5. Make Image Public

1. Go to your Docker Hub repository
2. Click "Settings" tab
3. Change visibility to "Public"

### 6. Create GitHub Repository

1. Create a new repository on GitHub
2. Upload all files from this project
3. Update `README.md` to replace 'yourusername' with your actual Docker Hub username

### 7. Update Documentation

In `README.md`, replace all instances of `yourusername` with your actual Docker Hub username:

```bash
# Find and replace
sed -i 's/yourusername/your-actual-username/g' README.md
```

## 📁 Files to Upload to GitHub

```
go-docker-demo/
├── .github/
│   └── workflows/
│       └── docker.yml
├── main.go
├── go.mod
├── Dockerfile
├── .dockerignore
├── .gitignore
├── docker-commands.sh
├── README.md
├── SETUP.md
└── ASSIGNMENT_SUBMISSION.md
```

## 🔗 Required Links

After completing the steps above, you should have:

1. **GitHub Repository Link**: `https://github.com/yourusername/go-docker-demo`
2. **Docker Hub Image Link**: `https://hub.docker.com/r/yourusername/go-docker-demo`

## ✅ Verification Checklist

- [ ] Go application runs locally with Docker
- [ ] Docker image builds successfully
- [ ] Container runs and responds to health checks
- [ ] Image is pushed to Docker Hub
- [ ] Docker Hub image is publicly accessible
- [ ] All code is uploaded to GitHub
- [ ] README.md contains correct Docker Hub link
- [ ] GitHub repository is public

## 🧪 Testing Commands

```bash
# Test health endpoint
curl http://localhost:8080/health

# Test users endpoint
curl http://localhost:8080/api/users

# Test user creation
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com"}'

# Test user retrieval
curl http://localhost:8080/api/users/1
```

## 📝 Final Submission

Submit the link to your GitHub repository. The repository should contain:

1. Complete Go source code
2. Working Dockerfile
3. README.md with Docker Hub link
4. All supporting files

## 🆘 Troubleshooting

If you encounter issues:

1. **Docker not working**: Follow the setup guide in `SETUP.md`
2. **Build fails**: Check Docker is running and you have internet connection
3. **Push fails**: Verify Docker Hub login and image name
4. **API not responding**: Check port 8080 is available and container is running

## 🎯 Success Criteria

Your submission will be successful if:

- ✅ Docker image builds without errors
- ✅ Container runs and serves the API
- ✅ Image is publicly available on Docker Hub
- ✅ GitHub repository contains all required files
- ✅ README.md has working Docker Hub link
- ✅ API endpoints respond correctly

Good luck! 🚀
