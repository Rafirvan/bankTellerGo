package database

import (
	"log"
	"os"
)

// EnsureFileExists runs on startup and checks if json file exist or create one if not
func EnsureFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("%s not initialized, Creating file %s", filePath, filePath)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		err2 := file.Close()
		if err2 != nil {
			return err2
		}
	}
	return nil
}
