package service

import "testing"

func TestMain(t *testing.T) {
	t.Run("createDatabaseDir", testCreateDatabaseDir)
	t.Run("createDatabaseFile", testCreateDatabaseFile)
	t.Run("createStudent", testCreateStudent)
	t.Run("getStudent", testGetStudent)
	t.Run("getStudents", testGetStudents)
	t.Run("updateStudents", testUpdateStudent)
	t.Run("deleteStudents", testDeleteStudent)

	t.Run("clean", testCleanDatabaseTestFilesAndDirectory)
}
