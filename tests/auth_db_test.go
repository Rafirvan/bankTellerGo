package tests

import (
	"BankTellerAPI/database"
	"BankTellerAPI/models"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUsers(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "users_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	mockUsers := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
	}

	err = json.NewEncoder(tmpFile).Encode(mockUsers)
	assert.NoError(t, err, "Error writing to temporary file")

	tmpFile.Seek(0, 0)

	database.UserDBPath = tmpFile.Name()

	users, err := database.ReadUsers()
	assert.NoError(t, err, "Error reading users")
	assert.Equal(t, len(mockUsers), len(users), "Expected number of users does not match")

	// Verify each user
	for i, user := range users {
		assert.Equal(t, mockUsers[i].Username, user.Username, "User at index %d does not match username", i)
		assert.Equal(t, mockUsers[i].Password, user.Password, "User at index %d does not match password", i)
	}
}

func TestFindUserByUsername(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "users_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	mockUsers := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
	}

	err = json.NewEncoder(tmpFile).Encode(mockUsers)
	assert.NoError(t, err, "Error writing to temporary file")

	tmpFile.Seek(0, 0)

	database.UserDBPath = tmpFile.Name()

	username := "user1"
	user, err := database.FindUserByUsername(username)
	assert.NoError(t, err, "Error finding user")
	assert.NotNil(t, user, "Expected user with username %s, but found none", username)
	assert.Equal(t, username, user.Username, "Expected username %s, got %s", username, user.Username)
}

func TestRegisterUser(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "users_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	mockUser := models.User{Username: "newUser", Password: "newPassword"}
	database.UserDBPath = tmpFile.Name()

	err = database.RegisterUser(mockUser)
	assert.NoError(t, err, "Error registering user")

	users, err := database.ReadUsers()
	assert.NoError(t, err, "Error reading users")
	assert.Contains(t, users, mockUser, "User was not found in the list after registration")
}

func TestWriteUsers(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "users_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	mockUsers := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
	}

	database.UserDBPath = tmpFile.Name()

	err = database.WriteUsers(mockUsers)
	assert.NoError(t, err, "Error writing users to file")

	users, err := database.ReadUsers()
	assert.NoError(t, err, "Error reading users after writing")
	assert.Equal(t, len(mockUsers), len(users), "Number of users after writing does not match")
}
