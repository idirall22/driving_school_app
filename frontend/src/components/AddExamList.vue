<template>
  <div id="addExamList">

    <b-alert
      v-model="examListCreated"
      :variant="alertVariant"
      dismissible
      @dismissed="closeAlert()">
      {{alertMessage}}
    </b-alert>
    <!-- title -->
    <Header initTitle="Ajouter une Liste D'Examen"></Header>

    <b-form @submit="createExamList">

      <!-- Exam date -->
      <b-form-group id="examDate"
        label="Date d'examen:" label-for="examDate">
        <vue-bootstrap-datetimepicker
          v-model="examDate"
          id="datepicker"
          :config="options">
        </vue-bootstrap-datetimepicker>
      </b-form-group>
      <!-- end exam Date -->

      <SearchStudent
        ref="search"
        v-on:studentAdded= "createStudentExam($event)">
      </SearchStudent>


      <div class="table-responsive mt-5">
        <table class="table table-striped table-sm">
          <thead>
            <tr>
              <th class="align-middle text-center">Ordre</th>
              <th class="align-middle text-center">Num Dossier</th>
              <th class="align-middle text-center">Nom et Prénom</th>
              <th class="align-middle text-center">Date De Naissance</th>
              <th class="align-middle text-center">Catégory</th>
              <th class="align-middle text-center">Examen</th>
              <th class="align-middle text-center"></th>
            </tr>

          </thead>
          <tbody>
            <tr v-for="(exam, index) in studentsExams" :key="exam.student_id">
              <td class="align-middle text-center">{{index+1}}</td>
              <td class="align-middle text-center">
                {{exam.student.file_number}}
              </td>
              <td class="align-middle text-center">
                {{exam.student.last_name_fr | capitalize}}
                {{exam.student.first_name_fr | capitalize}}
              </td>
              <td class="align-middle text-center">
                {{exam.student.birthday}}
              </td>
              <td class="align-middle text-center">B</td>
              <td class="align-middle text-center">
                {{exam.student.getExamName(exam.student.next_exam)}}</td>
              <td>
                <button @click.prevent="removeStudentExam(index)"
                  type="button" class="form-control close"
                  data-dismiss="alert" aria-label="Close">
                 <span aria-hidden="true">&times;</span>
               </button>
             </td>
            </tr>
          </tbody>
        </table>
        <b-button block type="submit" variant="primary">Ajouter la liste d'éxamen</b-button>
      </div>
    </b-form>
  </div>
</template>
<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import Header from './parts/Header';
import SearchStudent from './parts/SearchStudent';
import Exam from './service/exam.js';
import {ExamListMessages} from './service/messages.js'

import moment from 'moment'
export default {
  name: "addExamList",
  components: {
    VueBootstrapDatetimepicker,
    SearchStudent,
    Header
  },
  data: () => ({
    alertVariant:"",
    alertMessage:"",
    errorCreateExamList: null,
    examDate: null,
    examListCreated: false,

    examinerName: "",
    studentsExams: [],

    options:{
      format:"DD-MM-YYYY",
      useCurrent: false
    },
  }),
  methods:{
    closeAlert: function(){
      this.errorCreateExamList = null;
      this.examListCreated =  false;
    },
    createStudentExam: function(student){
      if(student != null){
        let exam = new Exam(student, null)
        this.studentsExams.push(exam)
      }
    },
    removeStudentExam: function(index){
      this.studentsExams.splice(index, 1)
    },
    createExamList: function(){
      let outStudentsExams = []

      let dateExamOut = moment(this.examDate, "DD-MM-YYYY").format();
      if(this.studentsExams.length > 0){
        for (var i = 0; i < this.studentsExams.length; i++) {
          outStudentsExams.push(this.studentsExams[i].
            outExam(dateExamOut));
        }
      }
      let examList = {
        "date_exam": dateExamOut,
        "examiner": this.examinerName,
        "students_exams": outStudentsExams,
        "archived": false,
      }
      window.backend.Service.CreateExamList(examList)
      .then(
        ()=>{},
        err=>{this.errorCreateExamList= err}
      )
      let messages = new ExamListMessages();
      if(this.errorCreateExamList != null){
        this.alertVariant = "warning";
        // this.alertMessage ="La liste d'éxamen n'a pas été crée";
        this.alertMessage = messages.EXAM_LIST_ERR_CREATED_MESSAGE;
      }else{
        this.alertVariant = "success";
        // this.alertMessage = "La liste d'éxamen a été crée";
        this.alertMessage = messages.EXAM_LIST_CREATED_MESSAGE;
      }
      this.examListCreated = true;
      this.examDate = null;
      this.studentsExams = [];
    },
  },
}
</script>
<style scoped>
</style>

