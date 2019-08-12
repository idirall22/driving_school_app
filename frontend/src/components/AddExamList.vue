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
        label="Date d'examen:" label-for="datepicker">
        <vue-bootstrap-datetimepicker
          v-model="examList.date_exam"
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
            <tr v-for="(exam, index) in examList.students_exams" :key="exam.student_id">
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
                {{exam.student.getExamName(exam.exam_level)}}</td>
              <td>
                <button @click.prevent="removeStudentExam(index)"
                  type="button" class="form-control close align-middle text-center"
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
import ExamList from './service/examList.js';
import {ExamListMessages} from './service/messages.js'

// import moment from 'moment'
export default {
  name: "addExamList",
  components: {
    VueBootstrapDatetimepicker,
    SearchStudent,
    Header
  },
  created:function(){
    this.examList = new ExamList(null);
  },
  data: () => ({
    alertVariant:"",
    alertMessage:"",
    errorCreateExamList: null,
    examListCreated: false,

    examList: null,

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
        let exam = new Exam(null, student, null);
        this.examList.addStudentToStudentsExams(exam);
      }
    },
    removeStudentExam: function(index){
      this.examList.students_exams.splice(index, 1)
    },
    displayMessage: function(msg, variant){
      this.alertVariant = variant;
      this.alertMessage = msg;
      this.examListCreated = true;
    },
    validateForm: function(){
      if(this.examList.date_exam == null){
        this.displayMessage("La date d'examen n'a pas été choisie", "warning")
        return true
      }
      if(this.examList.students_exams.length == 0){
        this.displayMessage("La liste d'étudiant est vide", "warning")
        return true
      }
      return false
    },
    createExamList: function(){
      this.closeAlert();
      if(this.validateForm()){
        return
      }
      window.backend.Service.CreateExamList(this.examList.outExamList())
      .then(
        (id)=>{
          this.$router.push({name: "examListDetails", params:{"id": id}})
        },
        err=>{
          this.errorCreateExamList= err
          let messages = new ExamListMessages();
          this.displayMessage(messages.EXAM_LIST_ERR_CREATED_MESSAGE, "warning")
        }
      )
    },
  },
}
</script>
<style scoped>
</style>

