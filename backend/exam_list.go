package service

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

var timeFormat = time.RFC3339

// CreateExamList create an exam list
func (s *Service) CreateExamList(date, examiner string,
	studentsList []interface{}) (*ExamList, error) {

	// Marshal an unmarshal the list interface into a list of students
	students := []*Student{}

	// Check if length > 0
	if len(studentsList) > 0 {
		data, err := json.Marshal(studentsList)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(data, &students); err != nil {
			return nil, err
		}
	}

	// Chek if there is a name for examiner
	if examiner == "" {
		examiner = "No Name"
	}

	// Parse exam date
	dateTime, err := time.Parse(timeFormat, date)
	if err != nil {
		return nil, err
	}

	// Create an exam instance
	examList := &ExamList{
		DateExam: dateTime,
		Examiner: examiner,
		Archived: false,
	}

	if err := MainService.db.Create(&examList).Error; err != nil {
		return nil, err
	}

	if len(students) > 0 {
		query := `INSERT INTO exams (
			exam,
			date_exam,
			status,
			student_id,
			exam_list_id
			)Values(
				?,?,?,?,?
				);`

		tx := MainService.db.Begin()
		for i := 0; i < len(students); i++ {
			if err := tx.Exec(query,
				students[i].NextExam,
				examList.DateExam,
				false,
				students[i].ID,
				examList.ID,
			).Error; err != nil {
				return nil, err
			}
		}
		tx.Commit()
	}

	return examList, nil
}

// GetExamList get an exam list by id
func (s *Service) GetExamList(id uint) (*ExamList, error) {

	examList := &ExamList{}
	tx := MainService.db.Begin()

	if err := tx.Find(&examList, "id=?", id).
		Related(&examList.StudentsExams, "StudentsExams").
		Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(examList.StudentsExams); i++ {
		if err := tx.Find(&examList.StudentsExams[i].Student, "id=?",
			examList.StudentsExams[i].StudentID).Error; err != nil {
			return nil, err
		}
	}
	tx.Commit()
	return examList, nil
}

// GetExamLists get a list of exam list
func (s *Service) GetExamLists(limit, offset uint) (*ExamListsOut, error) {
	examListsOut := &ExamListsOut{}

	if err := MainService.db.
		Model(&ExamList{}).
		Order("date_exam desc").
		Count(&examListsOut.Count).
		Limit(limit).Offset(offset).
		Find(&examListsOut.ExamLists).
		Error; err != nil {
		return nil, err
	}

	return examListsOut, nil
}

// UpdateExamList update an exam list
func (s *Service) UpdateExamList(examListID uint, date, examiner string,
	exams []interface{}, studentsList []interface{}) (*ExamList, error) {

	// Parse exam date
	dateParsed, errParse := time.Parse(timeFormat, date)
	if errParse != nil {
		return nil, errParse
	}

	newStudentsExams := []*Exam{}
	tx := MainService.db.Begin()

	// Check if there are new students added to pass exam
	if len(studentsList) > 0 {
		students := []*Student{}
		data, err := json.Marshal(studentsList)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &students); err != nil {
			return nil, err
		}
		for i := 0; i < len(students); i++ {
			exam := &Exam{
				Exam:       students[i].NextExam,
				DateExam:   dateParsed,
				StudentID:  students[i].ID,
				Student:    *students[i],
				ExamListID: examListID,
			}
			if err := tx.Create(&exam).Error; err != nil {
				return nil, err
			}
			newStudentsExams = append(newStudentsExams, exam)
		}
	}

	// Check if the examiner name is empty
	if examiner == "" {
		examiner = "No name"
	}

	// Create an examList object
	examList := &ExamList{
		ID:       examListID,
		Examiner: examiner,
		DateExam: dateParsed,
	}

	// Check if there are students in exams list
	if len(exams) > 0 {
		data, err := json.Marshal(exams)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &examList.StudentsExams); err != nil {
			return nil, err
		}
	}

	// Append the two exams slices
	examList.StudentsExams = append(examList.StudentsExams, newStudentsExams...)

	for i := 0; i < len(examList.StudentsExams); i++ {
		student := &Student{}
		if err := tx.
			Find(&student, "id=?", examList.StudentsExams[i].StudentID).
			UpdateColumn("next_exam", examList.StudentsExams[i].Exam+1).
			Error; err != nil {
			return nil, err
		}
	}
	// Delete All Exams Before add new ones
	examsToDelete := []Exam{}
	if err := tx.Find(&examsToDelete, "exam_list_id=?", examList.ID).
		Unscoped().Delete(&examsToDelete).Error; err != nil {
		return nil, err
	}

	// Update exam list an replace associated exams
	if err := tx.Save(&examList).
		Association("StudentsExams").
		Replace(examList.StudentsExams).
		Error; err != nil {
		return nil, err
	}

	tx.Commit()

	return examList, nil
}

// DeleteExamList update an exam list
func (s *Service) DeleteExamList(id uint) error {
	examList := ExamList{ID: id}

	tx := MainService.db.Begin()
	tx.Find(&examList, "id=?", id).Related(&examList.StudentsExams).
		Unscoped().Delete(&examList).
		Unscoped().Delete(&examList.StudentsExams)
	err := tx.Commit().Error
	return err
}

// ArchiveExamList archive or an archive examList
func (s *Service) ArchiveExamList(id uint, archived bool) error {
	return MainService.db.Model(&ExamList{}).Where("id=?", id).
		UpdateColumn("archived", archived).Error
}

// ExportExamListPDF allows to export an exam list as pdf
func (s *Service) ExportExamListPDF(id uint) error {
	examList, err := MainService.GetExamList(id)
	if err != nil {
		return err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 14)
	pdf.AddPage()
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.CellFormat(200, 10, tr("République Algérienne Démocratique et Populaire"),
		"", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Cell(100, 10, "MINISTER DES TRANSPORTS")
	pdf.Ln(-1)
	pdf.Cell(100, 10, "DIRECTION DES TRANSPORTS")
	pdf.Ln(-1)
	pdf.Cell(100, 10, "DE LA WILAYA D'ORAN")
	pdf.Ln(-1)
	pdf.CellFormat(200, 10, "Liste D'EXAMEN", "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(200, 10, "DU PERMIS DE CONDUITE", "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(80, 10, "Centre d'examen Oran: USTO", "", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, "Date d'examen 13.06.2019", "", 0, "R", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(80, 10, tr("Nome et prénom de l'examinateur: ............................."),
		"", 0, "L", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(-1)

	pdf.CellFormat(20, 8, "N", "TLR", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "N", "TLR", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Date de", "TLR", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "Nom et Prenom", "TLR", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8, "Cat", "TLR", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Nature de", "TLR", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Resultats", "TLR", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(20, 8, "Ordre", "LRB", 0, "C", false, 0, "")
	pdf.CellFormat(20, 8, "Dossier", "LRB", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Naissance", "LRB", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "", "LRB", 0, "L", false, 0, "")
	pdf.CellFormat(15, 8, "", "LRB", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "L'examen", "LRB", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "", "LRB", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 13)
	// cellsWidth := []float64{20, 20, 30, 50, 15, 30, 30}
	// align := []string{"C", "C", "L", "C", "C", "C", "C"}
	// students := [][]string{
	// 	{"01", "11236", "20-09-1992", "Nom et Prénom", "B", "code", "AJ"},
	// 	{"02", "11659", "20-09-1992", "Nom et Prénom", "B", "code", "AJ"},
	// 	{"03", "12365", "20-09-1992", "Nom et Prénom", "B", "code", "AJ"},
	// 	{"04", "23158", "20-09-1992", "Nom et Prénom", "B", "code", "AJ"},
	// 	{"05", "84785", "20-09-1992", "Nom et Prénom", "B", "code", "AJ"},
	// 	{"06", "87865", "20-09-1992", "Nom et Prénom", "B", "code", "AJ"},
	// }
	for i, exam := range examList.StudentsExams {
		pdf.CellFormat(20, 10, strconv.Itoa(i), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, tr(exam.Student.FileNumber), "1", 0, "C", false, 0, "")
		tm := exam.Student.BirthDay.Format("02-01-2006")
		pdf.CellFormat(30, 10, tm, "1", 0, "C", false, 0, "")

		mln := &MultiLanguageField{}
		mfn := &MultiLanguageField{}

		if err := json.Unmarshal([]byte(exam.Student.LastName),
			&mln); err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(exam.Student.FirstName),
			&mfn); err != nil {
			return err
		}

		pdf.CellFormat(50, 10, tr(mln.FR)+" "+tr(mfn.FR), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 10, "B", "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, "code", "1", 0, "C", false, 0, "")
		status := "ajourné"
		if exam.Status == true {
			status = "réussie"
		}
		pdf.CellFormat(30, 10, tr(status), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.SetFont("Arial", "B", 14)

	pdf.Ln(-1)

	pdf.Cell(50, 10, tr("Moniteur: Néant"))
	pdf.Cellf(50, 10, tr("Candidats convoqués: %d"), len(examList.StudentsExams))
	pdf.Ln(-1)

	pdf.Cellf(50, 10, "Code: %d", len(examList.StudentsExams))
	pdf.Cell(50, 10, "Code:")
	pdf.Ln(-1)

	pdf.Cellf(50, 10, "Manoeuvre: %d", len(examList.StudentsExams))
	pdf.Cell(50, 10, "Manoeuvre:")
	pdf.Ln(-1)

	pdf.Cellf(50, 10, "Circulation: %d", len(examList.StudentsExams))
	pdf.Cell(50, 10, "Circulation:")
	pdf.Ln(-1)

	pdf.CellFormat(0, 10, "Signature de l'inspecteur", "", 0, "R", false, 0, "")
	pdf.Ln(-1)

	return pdf.OutputFileAndClose("exam.pdf")
}
