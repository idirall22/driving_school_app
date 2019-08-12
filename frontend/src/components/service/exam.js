// import moment from 'moment';

// import Student from './student.js'

export default class Exam{
  constructor(exam, student, dateExam){
    if(exam != null){
      this.date_exam = exam.date_exam;
      this.exam_level = exam.exam_level
      this.status = exam.status;
      this.student = exam.student;
      this.student_id = exam.student_id;
    }else{
      this.date_exam = dateExam;
      this.exam_level = student.exam_level
      this.status = false;
      this.student = student;
      this.student_id = student.id;
    }
  }
  outExam(dateExam){
    if(dateExam != null){
      this.date_exam = dateExam;
    }
    return {
      "date_exam": this.date_exam,
      "exam_level": this.exam_level,
      "status": this.status,
      "student": this.student.outStudent(),
      "student_id": this.student_id,
    }
  }
}

