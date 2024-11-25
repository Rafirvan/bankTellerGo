package database

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var JWTDBPath = "jwt.json"

// AddTokenToBlacklist prevents a JWT from being used after logout
func AddTokenToBlacklist(token string) error {
	log.Println("Adding token to blacklist")
	tokens, err := ReadBlacklistedTokens()
	if err != nil {
		return err
	}

	tokens = append(tokens, token)

	return WriteBlacklistedTokens(tokens)
}

// IsTokenBlacklisted returns true if the token is blacklisted
func IsTokenBlacklisted(token string) (bool, error) {
	log.Println("Checking if token is blacklisted")
	file, err := os.Open(JWTDBPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	blacklistedTokens, err := ReadBlacklistedTokens()

	for _, blacklistedToken := range blacklistedTokens {
		if blacklistedToken == token {
			return true, nil
		}
	}

	return false, nil
}

// ReadBlacklistedTokens reads blacklisted tokens from the file
func ReadBlacklistedTokens() ([]string, error) {
	log.Println("Reading blacklisted tokens")
	var blacklistedTokens []string

	file, err := os.Open(JWTDBPath)
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
		return blacklistedTokens, nil
	}

	err = json.NewDecoder(file).Decode(&blacklistedTokens)
	if err != nil {
		log.Println("error decoding file:", JWTDBPath)
		return nil, err
	}

	return blacklistedTokens, nil
}

func WriteBlacklistedTokens(tokens []string) error {
	log.Println("Writing blacklisted tokens to db")
	file, err := os.Create(JWTDBPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tokens)
	if err != nil {
		return err
	}

	return nil
}
