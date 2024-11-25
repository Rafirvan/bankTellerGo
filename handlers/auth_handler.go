package handlers

import (
	"BankTellerAPI/database"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"BankTellerAPI/models"
	"BankTellerAPI/utils"
)

// AddUserHandler runs on startup to create a user, not through an endpoint
func AddUserHandler(username string, password string) error {
	log.Println("firing AddUserHandler")
	foundUser, _ := database.FindUserByUsername(username)
	if foundUser != nil {
		return nil
	}
	newId := uuid.NewString()
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	cpassword := string(bcryptPassword)

	user := models.User{
		Id:       newId,
		Username: username,
		Password: cpassword,
	}

	err = database.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}

// LoginHandler handles user login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("firing LoginHandler")
	var userRequest models.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := database.FindUserByUsername(userRequest.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	compareErr := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(userRequest.Password))

	if compareErr != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(foundUser.Id)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		return
	}
}

// LogoutHandler handles logout requests
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("firing LogoutHandler")
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Blacklist the token
	err := database.AddTokenToBlacklist(string(tokenString))
	if err != nil {
		http.Error(w, "Error blacklisting token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
