#!/bin/bash

# Docker commands for Go Docker Demo
# Make sure to replace 'yourusername' with your actual Docker Hub username

DOCKER_IMAGE="yourusername/go-docker-demo"
TAG="latest"

echo "🐳 Go Docker Demo - Docker Commands"
echo "=================================="

case "$1" in
    "build")
        echo "Building Docker image..."
        docker build -t $DOCKER_IMAGE:$TAG .
        echo "✅ Build complete!"
        ;;
    "run")
        echo "Running Docker container..."
        docker run -p 8080:8080 $DOCKER_IMAGE:$TAG
        ;;
    "test")
        echo "Testing Docker container..."
        docker run -d -p 8080:8080 --name go-demo-test $DOCKER_IMAGE:$TAG
        sleep 5
        echo "Testing health endpoint..."
        curl -f http://localhost:8080/health || echo "❌ Health check failed"
        echo "Testing users endpoint..."
        curl -f http://localhost:8080/api/users || echo "❌ Users endpoint failed"
        docker stop go-demo-test
        docker rm go-demo-test
        echo "✅ Tests complete!"
        ;;
    "push")
        echo "Pushing to Docker Hub..."
        docker push $DOCKER_IMAGE:$TAG
        echo "✅ Push complete!"
        ;;
    "clean")
        echo "Cleaning up Docker resources..."
        docker stop $(docker ps -q --filter ancestor=$DOCKER_IMAGE:$TAG) 2>/dev/null || true
        docker rm $(docker ps -aq --filter ancestor=$DOCKER_IMAGE:$TAG) 2>/dev/null || true
        docker rmi $DOCKER_IMAGE:$TAG 2>/dev/null || true
        echo "✅ Cleanup complete!"
        ;;
    *)
        echo "Usage: $0 {build|run|test|push|clean}"
        echo ""
        echo "Commands:"
        echo "  build  - Build the Docker image"
        echo "  run    - Run the Docker container"
        echo "  test   - Test the Docker container"
        echo "  push   - Push to Docker Hub"
        echo "  clean  - Clean up Docker resources"
        echo ""
        echo "Don't forget to update 'yourusername' in this script!"
        ;;
esac
