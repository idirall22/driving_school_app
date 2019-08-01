package service

import "testing"

func TestMain(t *testing.T) {
	t.Run("createDatabaseDir", testCreateDatabaseDir)
	t.Run("createDatabaseFile", testCreateDatabaseFile)

	t.Run("createStudent", testCreateStudent)
	t.Run("createExamList", testCreateExamList)

	t.Run("getExamList", testGetExamList)
	t.Run("getStudent", testGetStudent)

	t.Run("getExamLists", testGetExamLists)
	t.Run("getStudents", testGetStudents)

	t.Run("updateExamList", testUpdateExamList)
	t.Run("updateStudents", testUpdateStudent)

	t.Run("deleteExamList", testDeleteExamList)
	t.Run("deleteStudents", testDeleteStudent)

	t.Run("cleanedDatabase", testCleanedDatabase)
	t.Run("clean", testCleanDatabaseTestFilesAndDirectory)
}
