package service

import (
	"os"
	"path/filepath"
)

// Create a database directory if not exist.
func createDatabaseDir(databaseDir string, permission os.FileMode) error {

	_, err := os.Stat(databaseDir)

	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(databaseDir, permission); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

// Create a database file if not exist
func createDatabaseFile(databaseDir, databaseName string,
	permission os.FileMode) error {

	path := filepath.Join(databaseDir, databaseName)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, permission)

	defer file.Close()
	if err != nil {
		return err
	}
	return nil
}
