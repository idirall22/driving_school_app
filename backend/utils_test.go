package service

import (
	"os"
	"testing"
)

// Test create database directory
func testCreateDatabaseDir(t *testing.T) {

	// When database directory name is valid
	if err := createDatabaseDir("db", 0777); err != nil {
		t.Error(err)
	}

	// When database directory name is not valid
	if err := createDatabaseDir("/ db /", 0777); err == nil {
		t.Error(err)
	}
}

// Test create database file
func testCreateDatabaseFile(t *testing.T) {

	// When database directory name and file name are  valid
	if err := createDatabaseFile("db", "db.sqlite3", 0777); err != nil {
		t.Error(err)
	}

	// When database directory name or file name is not valid
	if err := createDatabaseFile("/ db /", "db.sqlite3", 0777); err == nil {
		t.Error(err)
	}
}

func testCleanDatabaseTestFilesAndDirectory(t *testing.T) {
	if err := os.RemoveAll("db"); err != nil {
		t.Error(err)
	}
}
