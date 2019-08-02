package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func openJSONStudentFile(fileName string) arrayMap {
	if fileName == "" {
		fileName = "students_test.json"
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Could not open %s file: %s.\n", fileName, err)
		os.Exit(1)
	}
	a := arrayMap{}
	if err := json.Unmarshal(data, &a); err != nil {
		log.Fatal(err)
	}
	return a
}
