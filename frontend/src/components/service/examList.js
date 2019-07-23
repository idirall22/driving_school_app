import moment from 'moment';
// import Student from './student.js'


const DATE_FORMAT = "DD-MM-YYYY";


export default class ExamList{

  EXAM_LIST_CREATED_MESSAGE = "La liste d'éxamen a été crée"
  EXAM_LIST_ERR_CREATED_MESSAGE = "La liste d'éxamen n'a pas été crée"
  EXAM_LIST_UPDATED_MESSAGE = "La liste d'éxamen a été modifier"
  EXAM_LIST_ERR_UPDATED_MESSAGE = "La liste d'éxamen n'a pas été modifier"
  EXAM_LIST_ARCHIVED_MESSAGE = "La liste d'éxamen a été archivé"
  EXAM_LIST_ERR_ARCHIVED_MESSAGE = "La liste d'éxamen n'a pas été archivé"
  EXAM_LIST_DEARCHIVED_MESSAGE = "La liste d'éxamen a été déarchivé"
  EXAM_LIST_ERR_DEARCHIVED_MESSAGE = "La liste d'éxamen n'a pas été déarchivé"

  constructor(examList) {
    this.examListInfos = examList;

    // Check if there is a key students_exams in examListInfos
    this.createStudentsExamslistIfNotExist();
    this.momentExamDate();
  }
  momentExamDate(){
    if("date_exam" in this.examListInfos){
      let de = this.examListInfos.date_exam;
      this.examListInfos.date_exam = moment(de).format(DATE_FORMAT);
    }
  }
  // return false if student is not in students_exams list
  checkIfStudentInExamList(studentID){
    if(this.examListInfos.students_exams.length > 0){
      for (var i = 0; i < this.examListInfos.students_exams.length; i++) {
        if(this.examListInfos.students_exams[i].id == studentID){
          return [true, i];
        }
      }
    }
    return [false, -1]
  }

  createStudentsExamslistIfNotExist(){
    if(!("students_exams" in this.examListInfos)){
      this.examListInfos.students_exams = [];
    }
  }

  addStudentToStudentsExams(student){
    if(!this.checkIfStudentInExamList(student)[0]){
      this.examListInfos.students_exams.push(student);
    }
  }

  removeStudentFromStudentsExams(studentID){
    let check = this.checkIfStudentInExamList(studentID);
    if(check[0]){
      this.examListInfos.students_exams.splice(check[1], 1);
    }
  }

  updateExamList(){
    if("date_exam" in this.examListInfos){
      let ed = this.examListInfos.date_exam;
      this.examListInfos.date_exam = moment(ed, DATE_FORMAT).format();
    }
  }

  archiveAndDearchive(){
    this.examListInfos.archived = !this.examListInfos.archived;
  }
}