package service

import "testing"

func TestMain(t *testing.T) {
	t.Run("createDatabaseDir", testCreateDatabaseDir)
	t.Run("createDatabaseFile", testCreateDatabaseFile)
	t.Run("getStudent", testGetStudent)
	t.Run("getStudents", testGetStudents)

	t.Run("clean", testCleanDatabaseTestFilesAndDirectory)
}
