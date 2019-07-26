import moment from 'moment';

const DATE_FORMAT = "DD-MM-YYYY";

export default class Student{
  constructor(student, exams){
    this.studentInfos = student;
    this.exams = exams;
    this.parseStudent();
  }

  parseFirstName(){
    if("first_name" in this.studentInfos){
      try {
        let ln = JSON.parse(this.studentInfos.first_name);
        this.studentInfos.first_name_fr = ln.fr;
        this.studentInfos.first_name_ar = ln.ar;
      } catch (e) {
        this.studentInfos.first_name_fr = "none";
        this.studentInfos.first_name_ar = "none";
      }
    }
  }

  parseLastName(){
    if("last_name" in this.studentInfos){
      try {
        let ln = JSON.parse(this.studentInfos.last_name);
        this.studentInfos.last_name_fr = ln.fr;
        this.studentInfos.last_name_ar = ln.ar;
      } catch (e) {
        this.studentInfos.last_name_fr = "none";
        this.studentInfos.last_name_ar = "none";
      }
    }
  }

  parseMaidenName(){
    if("maiden_name" in this.studentInfos){
      try {
        let ln = JSON.parse(this.studentInfos.maiden_name);
        this.studentInfos.maiden_name_fr = ln.fr;
        this.studentInfos.maiden_name_ar = ln.ar;
      } catch (e) {
        this.studentInfos.maiden_name_fr = "none";
        this.studentInfos.maiden_name_ar = "none";
      }
    }
  }

  parseStudentTimeInfos(){
    if("birthday" in this.studentInfos){
      this.studentInfos.birthday =
      moment(this.studentInfos.birthday).format(DATE_FORMAT);
    }
    if("registred_date" in this.studentInfos){
      this.studentInfos.registred_date =
      moment(this.studentInfos.registred_date).format(DATE_FORMAT);
    }
  }

  parseStudent(){
    this.parseFirstName();
    this.parseLastName();
    this.parseMaidenName();
    this.parseStudentTimeInfos();
  }
  // Delete all parsed languages key
  outStudent(){
    if("first_name_fr" in this.studentInfos){
      this.studentInfos.first_name = "";
      this.studentInfos.first_name = JSON.stringify(
        {
          "fr": this.studentInfos.first_name_fr,
          "ar": this.studentInfos.first_name_ar
        }
      )
    }
    if("last_name_fr" in this.studentInfos){
      this.studentInfos.last_name = "";
      this.studentInfos.last_name = JSON.stringify(
        {
          "fr": this.studentInfos.last_name_fr,
          "ar": this.studentInfos.last_name_ar
        }
      )
    }
    if("maiden_name_fr" in this.studentInfos){
      this.studentInfos.maiden_name = "";
      this.studentInfos.maiden_name = JSON.stringify(
        {
          "fr": this.studentInfos.maiden_name_fr,
          "ar": this.studentInfos.maiden_name_ar
        }
      )
    }
    if("registred_date" in this.studentInfos){
      let rd = this.studentInfos.registred_date;
      this.studentInfos.registred_date = moment(rd, DATE_FORMAT).format();
    }
    if("birthday" in this.studentInfos){
      let br = this.studentInfos.birthday;
      this.studentInfos.birthday = moment(br, DATE_FORMAT).format();
    }
  }

  // Translate to french language
  getStudentNextExamName(){
    switch (this.studentInfos.next_exam) {
      case 1:
        return "Code";
      case 2:
        return "Créneau";
      case 3:
        return "Circuit";
    }
  }
  getExamName(index){
    switch (index) {
      case 1:
        return "Code";
      case 2:
        return "Créneau";
      case 3:
        return "Circuit";
      case 4:
        return "Fini";
    }
  }
}
