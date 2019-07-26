import moment from 'moment';

const DATE_FORMAT = "DD-MM-YYYY";

export default class Student{
  constructor(studentObject, exams){
    this.id = studentObject["id"]
    this.created_at  = studentObject["created_at"]
    this.updated_at  = studentObject["updated_at"]
    this.file_number = studentObject["file_number"]

    this.first_name = studentObject["first_name"]
    this.last_name = studentObject["last_name"]
    this.maiden_name = studentObject["maiden_name"]

    this.first_name_fr = "";
    this.last_name_fr = "";
    this.maiden_name_fr =  "";

    this.first_name_ar = "";
    this.last_name_ar = "";
    this.maiden_name_ar = "";

    this.phone_number = studentObject["phone_number"]
    this.job = studentObject["job"]
    this.birthday = studentObject["birthday"]
    this.gender = studentObject["gender"]
    this.country = studentObject["country"]
    this.city = studentObject["city"]
    this.department = studentObject["department"]
    this.address_street = studentObject["address_street"]
    this.registred_date = studentObject["registred_date"]
    this.image = studentObject["image"]
    this.next_exam = studentObject["next_exam"]

    if("last_exam_date" in studentObject){
      this.last_exam_date = studentObject["last_exam_date"]
    }

    this.archived = studentObject["archived"]

    this.exams = exams;
    this.parseStudent();
  }

  parseFirstName(){
    try {
      let ln = JSON.parse(this.first_name);
      this.first_name_fr = ln.fr;
      this.first_name_ar = ln.ar;
    } catch (e) {
      this.first_name_fr = "none";
      this.first_name_ar = "none";
    }
  }

  parseLastName(){
    try {
      let ln = JSON.parse(this.last_name);
      this.last_name_fr = ln.fr;
      this.last_name_ar = ln.ar;
    } catch (e) {
      this.last_name_fr = "none";
      this.last_name_ar = "none";
    }
  }

  parseMaidenName(){
    try {
      let ln = JSON.parse(this.maiden_name);
      this.maiden_name_fr = ln.fr;
      this.maiden_name_ar = ln.ar;
    } catch (e) {
      this.maiden_name_fr = "none";
      this.maiden_name_ar = "none";
    }
  }

  parseTimeInfos(){
    this.birthday =
    moment(this.birthday).format(DATE_FORMAT);

    this.registred_date =
    moment(this.registred_date).format(DATE_FORMAT);
  }

  parseStudent(){
    this.parseFirstName();
    this.parseLastName();
    this.parseMaidenName();
    this.parseTimeInfos();
  }
  // Delete all parsed languages key
  outStudent(){
    this.first_name = "";
    this.first_name = JSON.stringify(
      {
        "fr": this.first_name_fr,
        "ar": this.first_name_ar,
      }
    )

    this.last_name = "";
    this.last_name = JSON.stringify(
      {
        "fr": this.last_name_fr,
        "ar": this.last_name_ar
      }
    )

    this.maiden_name = "";
    this.maiden_name = JSON.stringify(
      {
        "fr": this.maiden_name_fr,
        "ar": this.maiden_name_ar
      }
    )

    if("registred_date" in this){
      let rd = this.registred_date;
      this.registred_date = moment(rd, DATE_FORMAT).format();
    }

    if("birthday" in this){
      let br = this.birthday;
      this.birthday = moment(br, DATE_FORMAT).format();
    }
    return {
      "id": this.id,
      "file_number": this.file_number,
      "first_name": this.first_name,
      "last_name": this.last_name,
      "maiden_name": this.maiden_name,
      "phone_number": this.phone_number,
      "job": this.job,
      "birthday": this.birthday,
      "gender": this.gender,
      "country": this.country,
      "city": this.city,
      "department": this.department,
      "address_street": this.address_street,
      "registred_date": this.registred_date,
      "next_exam": this.next_exam,
      "archived": this.archived,
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
