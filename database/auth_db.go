package database

import (
	"BankTellerAPI/models"
	"encoding/json"
	"log"
	"os"
)

const UserDBPath = "users.json"

// ReadUsers reads all users from the JSON file
func ReadUsers() ([]models.User, error) {
	log.Printf("Reading users from %s", UserDBPath)
	var users []models.User
	file, err := os.Open(UserDBPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing file:", err)
		}
	}()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() == 0 {
		return users, nil
	}

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		log.Println("error decoding file:", UserDBPath)
		return nil, err
	}

	return users, nil
}

// FindUserByUsername finds a user by their username
func FindUserByUsername(username string) (*models.User, error) {
	log.Printf("Finding user by username %s in db", username)
	users, err := ReadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	//if no user is found
	return nil, nil
}

// RegisterUser adds a new user to the JSON database
func RegisterUser(user models.User) error {
	log.Printf("Registering user %s to %s", user.Username, UserDBPath)
	users, err := ReadUsers()
	if err != nil {
		return err
	}
	users = append(users, user)
	return WriteUsers(users)
}

// WriteUsers writes the list of users to the JSON file
func WriteUsers(users []models.User) error {
	log.Printf("Writing users to %s", UserDBPath)
	file, err := os.Create(UserDBPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("error closing file:", err)
		}
	}()

	err = json.NewEncoder(file).Encode(users)
	if err != nil {
		return err
	}

	return nil
}
