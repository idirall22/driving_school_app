<template>

 <div v-if="loaded" id="examListDetails">
   <b-alert
     v-model="examListUpdated"
     :variant="alertVariant"
     dismissible
     @dismissed="closeAlert()">
     {{alertMessage}}
   </b-alert>

   <Header initTitle="Liste d'examen"></Header>

   <div class="mb-3 d-flex justify-content-between">
      <div class="d-flex justify-content-start">
        <button v-if="!examList.archived"
          class="btn btn-warning"
          @click="deleteExamList"
          type="button">Suprimer la liste
        </button>
        <b-button v-if="!examList.archived" :disabled="checkIfCanBeArchived()"
          class="ml-2 btn btn-success"
          @click="examList.archived=!examList.archived"
          type="button">Archivé
        </b-button>
        <b-button v-if="examList.archived" :disabled="checkIfCanBeArchived()"
          class="btn btn-success"
          @click="examList.archived=!examList.archived"
          type="button">Déarchivé
        </b-button>
      </div>
      <div class="">
        <button
          class="btn btn-danger"
          @click="exportPDF"
          type="button">PDF
        </button>
      </div>
   </div>

   <b-form @submit="updateExamList">
     <!-- Examiner Name -->
     <b-form-group id="examinerName"
       label="Nom et Prénom de L'éxaminateur:" label-for="examinerName">
       <b-form-input
         id="examinerName"
         v-model="examList.examiner"
         type="text"
         placeholder="Nom de l'éxaminateur"
         :readonly="examList.archived">
       </b-form-input>
     </b-form-group>
     <!-- End examiner Name -->

     <!-- Exam date -->
       <b-form-group id="examDate"
         label="Date d'examen:" label-for="examDate">
         <div v-if="!examList.archived">
           <vue-bootstrap-datetimepicker
             v-model="examList.date_exam"
             id="datepicker"
             :config="options">
           </vue-bootstrap-datetimepicker>
         </div>
         <div v-else>
           {{examList.date_exam}}
         </div>
       </b-form-group>
     <!-- end exam Date -->

     <SearchStudent v-if="!examList.archived" class="mb-3"
       ref="search"
       v-on:studentAdded= "createStudentExam($event)"
       >
     </SearchStudent>

     <!-- array students -->
     <div class="table-responsive">
       <table class="table table-striped table-sm">
         <thead>
           <tr>
             <th class="align-middle text-center">Ordre</th>
             <th class="align-middle text-center">Numero Dossier</th>
             <th class="align-middle text-center">Nom et Prénom</th>
             <th class="align-middle text-center">D/Naissance</th>
             <th class="align-middle text-center">Cat</th>
             <th class="align-middle text-center">Examen</th>
             <th class="align-middle text-center">Résultat</th>
             <th
               v-if="!examList.archived"
               class="align-middle text-center">Retirer
             </th>
           </tr>
         </thead>
         <tbody>
           <tr class="exams border-success"
               v-for="(exam, index) in examList.students_exams"
               :key="exam.student_id">

             <td class="align-middle text-center">{{index+1}}</td>
             <td class="align-middle text-center">
               {{exam.student.file_number}}
             </td>
             <td class="align-middle text-center">
               {{exam.student.last_name_fr }}
               {{exam.student.first_name_fr}}
             </td>
             <td class="align-middle text-center">
               {{exam.student.birthday}}
             </td>
             <td class="align-middle text-center">B</td>
             <td class="align-middle text-center">
               {{exam.student.getExamName(exam.exam)}}
             </td>
             <td v-if="examList.archived"
                  class="align-middle text-center"
             >
              {{getStatus(exam.status)}}
             </td>
             <td v-if="!examList.archived"
                  class="align-middle text-center">
               <b-form-checkbox
                  v-model="exam.status"
                ></b-form-checkbox>
              </td>
              <td v-if="examList.archived"></td>
             <td v-if="!examList.archived" class="align-middle text-center">
               <button @click.prevent="removeStudentFromStudentsExams(exam.id, index)"
                 type="button" class="form-control close"
                 data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </td>
           </tr>
         </tbody>
       </table>
       <b-button
         block type="submit"
         variant="primary">
         Modifier la liste d'éxamen
       </b-button>
     </div>
   </b-form>
 </div>
</template>

<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
// import Student from './service/student.js';
import ExamList from './service/examList.js';
import Exam from './service/exam.js';
import {ExamListMessages} from './service/messages.js';
import Header from './parts/Header';
import SearchStudent from './parts/SearchStudent';
import moment from 'moment';

export default {
  components: {
    VueBootstrapDatetimepicker,
    SearchStudent,
    Header
  },

  mounted() {
    this.getExamList();
  },
  name: "ExamListDetails",
  data: () => ({
    alertVariant:"",
    alertMessage:"",

    errorDelete: null,
    errorArchive: null,
    errorUpdateExamList: null,
    errorGetExamList: null,

    examListUpdated: false,
    loaded: false,

    id: 0,
    examList: {},
    errPDF: null,
    studentsExamsRemovedIds: [],
    options:{
      format:"DD-MM-YYYY",
      useCurrent: false
    },
  }),
  methods:{

    closeAlert:function(){
      this.examListUpdated = false;
      this.errorUpdateExamList = null;
      this.errorArchive = null;
      this.errorDelete = null;
    },
    createStudentExam: function(student){
      let exam = new Exam(student);
      this.examList.addStudentToStudentsExams(exam);
    },
    // get an exam list from database
    getExamList: function(){
      // Get id from url
      this.id = parseInt(this.$route.params.id, 10);

      // Make request to backend
      window.backend.Service.GetExamList(this.id)
      .then(
        data=>{
          this.examList = new ExamList(data);
          this.loaded = true;
        },
        err=>{
          this.errorGetExamList = err;
        }
      );
    },
    removeStudentFromStudentsExams: function(studentID, index){
      if(studentID != null){
        this.studentsExamsRemovedIds.push(studentID)
      }
      this.examList.students_exams.splice(index, 1);
    },
    deleteExamList: function(){
      window.backend.Service.DeleteExamList(this.id)
      .then(
        err=>this.errorDelete = err
      );

      // Redirect to exams lists
      if(this.errorDelete == null){
        this.$router.push({name: "examLists"})
      }
    },

    updateExamList: function(){
      // Update Exam list to make it exportable to backend
      let outExamList = this.examList.outExamList();

      // make request to backend
      window.backend.Service.UpdateExamList(
        outExamList,
        this.studentsExamsRemovedIds,
      ).then(
        (data)=>{
          if("students_exams" in data) {
            for (var i = 0; i < data.students_exams.length; i++) {
              let student = this.examList.students_exams[i].student;
              data.students_exams[i].student = student;
              // data.students_exams[i].student.parseStudent();
              this.examList.students_exams[i] = data.students_exams[i];
            }
          }
          this.data = data;
        },
        err=>{
          this.errorUpdateExamList = err;
        },
      );

      let message = new ExamListMessages()
      if(this.errorUpdateExamList != null){
        this.alertVariant = "warning";
        this.alertMessage = message.EXAM_LIST_ERR_UPDATED_MESSAGE;
      }else{
        this.alertVariant = "success";
        this.alertMessage = message.EXAM_LIST_UPDATED_MESSAGE;
      }
      this.examListUpdated = true;
      this.updated = true;
    },
    exportPDF: function(){
      window.backend.Service.ExportExamListPDF(this.id)
      .then(err=>this.errPDF = err);
    },
    checkIfCanBeArchived: function(){
      if(moment(this.examList.date_exam, "DD-MM-YYYY").format() >
        moment().format()){
        return false
      }
      return true
    },
    getStatus(status){
      if(status){
        return "Réussi"
      }
      return "Recaler"
    }
  }
}
</script>

<style scoped>
.exams{
  border-left: 8px solid red;
}
.students{
  border-left: 8px solid green;
}
</style>