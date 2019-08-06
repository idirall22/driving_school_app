// import moment from 'moment';

// import Student from './student.js'

export default class Exam{
  constructor(student, dateExam){
    this.exam = student.next_exam;
    this.date_exam = dateExam;
    this.status = false;
    this.student = student;
    this.student_id = student.id;
  }
  outExam(dateExam){
    return {
      "exam": this.exam,
      "date_exam": dateExam,
      "status": this.status,
      "student": this.student.outStudent(),
      "student_id": this.student_id,
    }
  }
}

