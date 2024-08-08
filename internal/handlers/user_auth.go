package handlers

import (
	"database/sql"
	"docker_go/internal/database"
	"docker_go/internal/models"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// user_auth.go will handle registering and logging in users

// to follow best practices, function naming should be in camelCase
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error validating struct: %s", err.Error()), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password: ", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Assign default role
	//user.Role = "user"

	_, err = database.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User created successfully")
}

// LoginUser will handle logging in users
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	row := database.DB.QueryRow("SELECT id, username, email, password, created_at FROM users WHERE email = ?", creds.Email)
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "User does not exist", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User logged in successfully")
}
