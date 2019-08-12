package service

import (
	"os"
	"path/filepath"
	"testing"
)

// Test create database directory
func testCreateDatabaseDir(t *testing.T) {
	// When database directory name is valid
	if err := createDatabaseDir("db", 0777); err != nil {
		t.Error(err)
	}
}

// Test create database file
func testCreateDatabaseFile(t *testing.T) {
	// When database directory name and file name are  valid
	if err := createDatabaseFile("db", "db.sqlite3", 0777); err != nil {
		t.Error(err)
	}
}

func testCleanDatabaseTestFilesAndDirectory(t *testing.T) {
	path := filepath.Join(GetRootDir(), "db")
	if err := os.RemoveAll(path); err != nil {
		t.Error(err)
	}
}
