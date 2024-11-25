package tests

import (
	"BankTellerAPI/database"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTokenToBlacklist(t *testing.T) {
	// Prepare the mock data and clean up before the test
	tmpFile, err := os.CreateTemp("", "jwt_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	// Set the JWTDBPath to the temporary file path
	database.JWTDBPath = tmpFile.Name()

	// Add a token to the blacklist
	token := "mockToken"
	err = database.AddTokenToBlacklist(token)
	assert.NoError(t, err, "Error adding token to blacklist")

	// Verify if the token has been added to the blacklist
	tokens, err := database.ReadBlacklistedTokens()
	assert.NoError(t, err, "Error reading blacklisted tokens")
	assert.Contains(t, tokens, token, "Token was not added to the blacklist")
}

func TestIsTokenBlacklisted(t *testing.T) {
	// Prepare the mock data and clean up before the test
	tmpFile, err := os.CreateTemp("", "jwt_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	// Set the JWTDBPath to the temporary file path
	database.JWTDBPath = tmpFile.Name()

	// Add a token to the blacklist
	token := "mockToken"
	err = database.AddTokenToBlacklist(token)
	assert.NoError(t, err, "Error adding token to blacklist")

	// Check if the token is blacklisted
	isBlacklisted, err := database.IsTokenBlacklisted(token)
	assert.NoError(t, err, "Error checking if token is blacklisted")
	assert.True(t, isBlacklisted, "Token should be blacklisted")

	// Check a token that is not blacklisted
	isBlacklisted, err = database.IsTokenBlacklisted("nonExistentToken")
	assert.NoError(t, err, "Error checking if token is blacklisted")
	assert.False(t, isBlacklisted, "Token should not be blacklisted")
}

func TestReadBlacklistedTokens(t *testing.T) {
	// Prepare the mock data and clean up before the test
	tmpFile, err := os.CreateTemp("", "jwt_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	// Set the JWTDBPath to the temporary file path
	database.JWTDBPath = tmpFile.Name()

	// Write some blacklisted tokens to the file
	mockTokens := []string{"token1", "token2"}
	err = database.WriteBlacklistedTokens(mockTokens)
	assert.NoError(t, err, "Error writing blacklisted tokens")

	// Read blacklisted tokens from the file
	tokens, err := database.ReadBlacklistedTokens()
	assert.NoError(t, err, "Error reading blacklisted tokens")
	assert.Equal(t, len(mockTokens), len(tokens), "Number of tokens does not match")
	assert.ElementsMatch(t, mockTokens, tokens, "Tokens read from file do not match the expected tokens")
}

func TestWriteBlacklistedTokens(t *testing.T) {
	// Prepare the mock data and clean up before the test
	tmpFile, err := os.CreateTemp("", "jwt_*.json")
	assert.NoError(t, err, "Error creating temporary file")
	defer os.Remove(tmpFile.Name())

	// Set the JWTDBPath to the temporary file path
	database.JWTDBPath = tmpFile.Name()

	// Write blacklisted tokens to the file
	mockTokens := []string{"token1", "token2"}
	err = database.WriteBlacklistedTokens(mockTokens)
	assert.NoError(t, err, "Error writing blacklisted tokens")

	// Verify that the tokens were written correctly
	tokens, err := database.ReadBlacklistedTokens()
	assert.NoError(t, err, "Error reading blacklisted tokens")
	assert.ElementsMatch(t, mockTokens, tokens, "Tokens written to file do not match the expected tokens")
}
