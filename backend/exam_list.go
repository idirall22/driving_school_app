package service

import (
	"encoding/json"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

// CreateExamList create an exam list
func (s *Service) CreateExamList(examListMap map[string]interface{}) (uint, error) {

	// Get examList model from json
	examList, err := getExamListModelFromMap(examListMap)
	if err != nil {
		return 0, err
	}

	// Begin tx
	tx := MainService.db.Begin()

	if err := tx.Create(&examList).
		Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return examList.ID, nil
}

// GetExamList get an exam list by id
func (s *Service) GetExamList(id uint) (*ExamList, error) {

	examList := &ExamList{}
	tx := MainService.db.Begin()

	if err := tx.Find(&examList, "id=?", id).
		Related(&examList.StudentsExams, "StudentsExams").
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := 0; i < len(examList.StudentsExams); i++ {
		if err := tx.Find(&examList.StudentsExams[i].Student, "id=?",
			examList.StudentsExams[i].StudentID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return examList, nil
}

// GetExamLists get a list of exam list
func (s *Service) GetExamLists(limit, offset uint) (*ExamListsOut, error) {
	examListsOut := &ExamListsOut{}

	tx := MainService.db.Begin()
	if err := tx.
		Model(&ExamList{}).
		Order("date_exam desc").
		Count(&examListsOut.Count).
		Limit(limit).Offset(offset).
		Find(&examListsOut.ExamLists).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return examListsOut, nil
}

// UpdateExamList update an exam list
func (s *Service) UpdateExamList(examListMap map[string]interface{},
	examListToDelete []interface{}) (*ExamList, error) {

	examList, err := getExamListModelFromMap(examListMap)
	if err != nil {
		return nil, err
	}
	tx := MainService.db.Begin()

	// Create exam model to by id provaided to delete them
	if examListToDelete != nil {
		oldExams := []*Exam{}
		if len(examListToDelete) > 0 {
			for i := 0; i < len(examListToDelete); i++ {
				exam := Exam{ID: uint(examListToDelete[i].(uint))}
				oldExams = append(oldExams, &exam)
			}
			if err := tx.Find(&oldExams, "exam_list_id=?", examList.ID).
				Unscoped().
				Delete(&examList.StudentsExams).
				Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	// Update examList and associated exams
	if errS := tx.
		Save(&examList).
		Association("StudentsExams").
		Replace(examList.StudentsExams).
		Error; err != nil {
		tx.Rollback()
		return nil, errS
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
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
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
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
