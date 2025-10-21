package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// User represents a user in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// In-memory storage for demo purposes
var users []User
var nextID = 1

func main() {
	// Initialize with some sample data
	users = []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", CreatedAt: time.Now()},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", CreatedAt: time.Now()},
	}
	nextID = 3

	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/users/", userHandler)

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Health check: http://localhost:%s/health", port)
	log.Printf("API docs: http://localhost:%s/", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Docker Demo API</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .endpoint { background: #f5f5f5; padding: 10px; margin: 10px 0; border-radius: 5px; }
        .method { font-weight: bold; color: #0066cc; }
    </style>
</head>
<body>
    <h1>🐳 Go Docker Demo API</h1>
    <p>Welcome to the Go Docker Demo API! This is a simple REST API built with Go and containerized with Docker.</p>
    
    <h2>Available Endpoints:</h2>
    <div class="endpoint">
        <span class="method">GET</span> <code>/health</code> - Health check endpoint
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <code>/api/users</code> - Get all users
    </div>
    <div class="endpoint">
        <span class="method">POST</span> <code>/api/users</code> - Create a new user
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <code>/api/users/{id}</code> - Get user by ID
    </div>
    <div class="endpoint">
        <span class="method">DELETE</span> <code>/api/users/{id}</code> - Delete user by ID
    </div>
    
    <h2>Example Usage:</h2>
    <pre>
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
    </pre>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Success: true,
		Message: "Service is healthy",
		Data: map[string]interface{}{
			"timestamp": time.Now().UTC(),
			"status":    "ok",
		},
	}
	json.NewEncoder(w).Encode(response)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case "GET":
		response := Response{
			Success: true,
			Message: "Users retrieved successfully",
			Data:    users,
		}
		json.NewEncoder(w).Encode(response)
		
	case "POST":
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		
		if newUser.Name == "" || newUser.Email == "" {
			response := Response{
				Success: false,
				Message: "Name and email are required",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		
		newUser.ID = nextID
		newUser.CreatedAt = time.Now()
		nextID++
		users = append(users, newUser)
		
		response := Response{
			Success: true,
			Message: "User created successfully",
			Data:    newUser,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extract user ID from URL path
	idStr := r.URL.Path[len("/api/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Invalid user ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Find user by ID
	var user *User
	var userIndex int = -1
	for i, u := range users {
		if u.ID == id {
			user = &u
			userIndex = i
			break
		}
	}
	
	if user == nil {
		response := Response{
			Success: false,
			Message: "User not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	switch r.Method {
	case "GET":
		response := Response{
			Success: true,
			Message: "User retrieved successfully",
			Data:    user,
		}
		json.NewEncoder(w).Encode(response)
		
	case "DELETE":
		users = append(users[:userIndex], users[userIndex+1:]...)
		response := Response{
			Success: true,
			Message: "User deleted successfully",
		}
		json.NewEncoder(w).Encode(response)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
