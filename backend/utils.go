package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GetRootDir return string root os path
func GetRootDir() string {
	r, _ := os.UserHomeDir()
	return r
}

// Create a database directory if not exist.
func createDatabaseDir(databaseDir string, permission os.FileMode) error {
	p := filepath.Join(GetRootDir(), databaseDir)
	_, err := os.Stat(p)
	fmt.Println(os.IsNotExist(err))
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(p, permission); err != nil {
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

	path := filepath.Join(GetRootDir(), databaseDir, databaseName)
	os.Create(path)
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
