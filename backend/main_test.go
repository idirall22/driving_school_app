package service

import "testing"

func TestMain(t *testing.T) {
	t.Run("createDatabaseDir", testCreateDatabaseDir)
	t.Run("createDatabaseFile", testCreateDatabaseFile)

	// t.Run("createStudent", testCreateStudent)
	t.Run("createStudentMap", testCreateStudentMap)

	t.Run("getStudent", testGetStudent)
	t.Run("getStudentByName", testGetStudentByName)

	t.Run("getStudents", testGetStudents)
	t.Run("updateStudents", testUpdateStudent)

	t.Run("createExamList", testCreateExamList)
	t.Run("getExamList", testGetExamList)
	t.Run("getExamLists", testGetExamLists)
	t.Run("updateExamList", testUpdateExamList)

	t.Run("deleteStudents", testDeleteStudent)
	t.Run("deleteExamList", testDeleteExamList)

	t.Run("clean", testCleanDatabaseTestFilesAndDirectory)
}
