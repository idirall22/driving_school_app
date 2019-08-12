package service

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
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

	// Check if last examList is archived before create a new one
	var examListLast = &ExamList{}
	if err := tx.Last(&examListLast).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, err
		}
	}
	if examListLast.ID != 0 {
		if examListLast.Archived != true {
			return 0, errors.New("The last exam list need to be archived before create a new one")
		}
	}

	for _, exam := range examList.StudentsExams {
		exam.Student = nil
	}

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
		if examList.StudentsExams[i].Student == nil {
			examList.StudentsExams[i].Student = &Student{}
		}
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
func (s *Service) UpdateExamList(examListMap map[string]interface{}) (*ExamList, error) {

	examList, err := getExamListModelFromMap(examListMap)
	if err != nil {
		return nil, err
	}

	tx := MainService.db.Begin()

	ids := []uint{}
	for _, exam := range examList.StudentsExams {
		exam.Student = nil
		ids = append(ids, exam.ID)
	}

	if examList.Archived {
		for _, exam := range examList.StudentsExams {
			student := &Student{}

			tx.Find(&student, "id=?", exam.StudentID)
			// setStudentExamInfos(student, exam)
			if student.LastExamDate != nil {
				if student.LastExamDate.Unix() == exam.DateExam.Unix() {
					if *student.LastExamStatus {
						student.ExamLevel--
						if student.ExamLevel == 3 {
							student.WinLicenceDate = nil
						}
					}
				}
			}
			student.LastExamDate = &exam.DateExam
			student.LastExamStatus = &exam.Status
			if exam.Status {
				student.ExamLevel++
				if student.ExamLevel >= 4 {
					student.WinLicenceDate = &exam.DateExam
				}
			}
			tx.Save(&student)
		}
	}

	// if examList.DateExam.Unix() >= time.Now().Unix() {
	if err := tx.Unscoped().
		Where("exam_list_id = ? AND id NOT IN (?)", examList.ID, ids).
		Delete(&Exam{}).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// }

	if err := tx.Save(&examList).
		// Association("StudentsExams").
		// Replace(examList.StudentsExams).
		Error; err != nil {
		tx.Rollback()
		return nil, err
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
	pdf.CellFormat(100, 10, "Date d'examen "+examList.DateExam.Format("02-01-2006"), "", 0, "R", false, 0, "")
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

	lvl1 := 0
	lvl2 := 0
	lvl3 := 0

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
		examName := "code"
		for _, exam := range examList.StudentsExams {
			switch exam.ExamLevel {
			case 1:
				lvl1++
				examName = "code"
				break
			case 2:
				lvl2++
				examName = "Manoeuvre"
				break
			case 3:
				lvl3++
				examName = "Circuit"
				break
			}
		}
		pdf.CellFormat(50, 10, tr(mln.FR)+" "+tr(mfn.FR), "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 10, "B", "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, examName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, tr(""), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.SetFont("Arial", "B", 14)

	pdf.Ln(-1)

	pdf.Cell(50, 10, tr("Moniteur: Néant"))
	pdf.Cellf(50, 10, tr("Candidats convoqués: %d"), len(examList.StudentsExams))
	pdf.Ln(-1)

	pdf.Cellf(50, 10, "Code: %d", lvl1)
	pdf.Cell(50, 10, "Code:")
	pdf.Ln(-1)

	pdf.Cellf(50, 10, "Manoeuvre: %d", lvl2)
	pdf.Cell(50, 10, "Manoeuvre:")
	pdf.Ln(-1)

	pdf.Cellf(50, 10, "Circulation: %d", lvl3)
	pdf.Cell(50, 10, "Circulation:")
	pdf.Ln(-1)

	pdf.CellFormat(0, 10, "Signature de l'inspecteur", "", 0, "R", false, 0, "")
	pdf.Ln(-1)

	yy, mm, dd := examList.DateExam.Date()
	yys := strconv.Itoa(yy)
	dds := strconv.Itoa(dd)
	nameFile := strings.Join([]string{dds, mm.String(), yys}, "-")
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	p := path.Join(dir, "Desktop", nameFile+".pdf")
	return pdf.OutputFileAndClose(p)
}
