import moment from 'moment';
import Student from './student.js'


const DATE_FORMAT = "DD-MM-YYYY";


export default class ExamList{
  constructor(examListObject) {
    this.id = examListObject["id"]
    this.created_at = examListObject["created_at"]
    this.updated_at = examListObject["updated_at"]
    this.date_exam = examListObject["date_exam"]
    this.examiner = examListObject["examiner"]
    this.archived = examListObject["archived"]

    if(!("students_exams" in examListObject)){
      this.students_exams = [];
    }else{
      this.students_exams = examListObject["students_exams"]
    }
    this.convertToStudentObject();
    this.momentExamDate();
  }


  momentExamDate(){
      let de = this.date_exam;
      this.date_exam = moment(de).format(DATE_FORMAT);
  }

  convertToStudentObject(){
    if(this.students_exams.length > 0){
      for (var i = 0; i < this.students_exams.length; i++) {
        let student = new Student(this.students_exams[i].student, null);
        this.students_exams[i].student = student;
      }
    }
  }
  // return false if student is not in students_exams list
  checkIfStudentInExamList(studentID){
    if(this.students_exams.length > 0){
      for (var i = 0; i < this.students_exams.length; i++) {
        if(this.students_exams[i].id == studentID){
          return [true, i];
        }
      }
    }
    return [false, -1]
  }

  // Add student to students_exams list
  addStudentToStudentsExams(student){
    if(!this.checkIfStudentInExamList(student)[0]){
      this.students_exams.push(student);
    }
  }

  // Remove student from students_exams list
  removeStudentFromStudentsExams(studentID){
    let check = this.checkIfStudentInExamList(studentID);
    if(check[0]){
      this.students_exams.splice(check[1], 1);
    }
  }

  outExamList(){
    let ed = this.date_exam;
    this.date_exam = moment(ed, DATE_FORMAT).format();

    for (var i = 0; i < this.students_exams.length; i++) {
      this.students_exams[i].student =
      this.students_exams[i].student.outStudent();
    }
    return {
      "id": this.id,
      "date_exam": this.date_exam,
      "examiner": this.examiner,
      "archived": this.archived,
      "students_exams": this.students_exams,
    }
  }

  // Archive an dearchive an examList
  archiveAndDearchive(){
    this.archived = !this.archived;
  }
}