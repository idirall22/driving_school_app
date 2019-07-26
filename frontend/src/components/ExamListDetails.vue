<template>

 <div v-if="loaded" id="examListDetails">
   {{errorUpdateExamList}}
   <b-alert
     v-model="examListUpdated"
     :variant="alertVariant"
     dismissible
     @dismissed="closeAlert()">
     {{alertMessage}}
   </b-alert>

   <Header initTitle="Liste d'examen"></Header>

   <div class="mb-3 d-flex justify-content-between">
      <div class="">
        <button v-if="!examList.examListInfos.archived"
          class="btn btn-warning"
          @click="deleteExamList"
          type="button">Suprimer la liste
        </button>
        <button v-if="!examList.examListInfos.archived"
          class="ml-2 btn btn-success"
          @click="archiveExamList()"
          type="button">Archivé
        </button>
        <button v-else
          class="ml-2 btn btn-warning"
          @click="archiveExamList()"
          type="button">Modifier
        </button>
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
       label="Cherché des étudiants:" label-for="examinerName">
       <b-form-input
         id="examinerName"
         v-model="examList.examListInfos.examiner"
         type="text"
         placeholder="Nom de l'éxaminateur"
         :readonly="examList.examListInfos.archived">
       </b-form-input>
     </b-form-group>
     <!-- End examiner Name -->
     <!-- Exam date -->
       <b-form-group id="examDate"
         label="Date d'examen:" label-for="examDate">
         <div v-if="!examList.examListInfos.archived">
           <vue-bootstrap-datetimepicker
             v-model="examList.examListInfos.date_exam"
             id="datepicker"
             :config="options">
           </vue-bootstrap-datetimepicker>
         </div>
         <div v-else>
           {{examList.examListInfos.date_exam}}
         </div>
       </b-form-group>
     <!-- end exam Date -->
     <!-- Search students -->
     <b-form-group id="searchStudents"
       label="Cherché des étudiants:" label-for="searchStudents">
       <b-form-input
         id="searchStudents"
         v-model="studentLastName"
         type="text"
         placeholder="Nom de l'étudiant"
         :readonly="examList.examListInfos.archived">
       </b-form-input>
     </b-form-group>
     <b-button
       :disabled="examList.examListInfos.archived"
       @click="searchStudent()"
       variant="primary">Cherché
     </b-button>
     <!-- end search students -->
     <div v-if="found">
       <div class="col mt-3 p-0">
         <ul class="list-group p-0 m-0">
           <li v-for="(student, index) in studentsFound" :key="student.studentInfos.id"
             class="list-group-item d-flex justify-content-between align-items-center">
             <p class="p-0 m-0">
               {{student.studentInfos.last_name_fr | capitalize}}
               {{student.studentInfos.first_name_fr | capitalize}}
             </p>
             <button class="btn btn-primary ml-3"
               @click="addStudent(student.studentInfos.id, index)"
               type="button"
               name="button">Add
             </button>
           </li>
         </ul>
       </div>
     </div>
     <br>
     <br>
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
             <th v-if="!examList.archived" class="align-middle text-center">Retirer</th>
           </tr>
         </thead>
         <tbody>
           <tr v-for="(exam, index) in examList.examListInfos.students_exams"
              :key="exam.student.studentInfos.id">

             <td class="align-middle text-center">{{index+1}}</td>
             <td class="align-middle text-center">
               {{exam.student.studentInfos.file_number}}
             </td>
             <td class="align-middle text-center">
               {{exam.student.studentInfos.last_name_fr | capitalize}}
               {{exam.student.studentInfos.first_name_fr | capitalize}}
             </td>
             <td class="align-middle text-center">
               {{exam.student.studentInfos.birthday}}
             </td>
             <td class="align-middle text-center">B</td>
             <td class="align-middle text-center">{{exam.student.getStudentNextExamName()}}</td>
             <td class="align-middle text-center">
               <input v-model="exam.status" type="checkbox">
             </td>
             <td class="align-middle text-center">
               <button @click.prevent="removeStudent(index)"
                 type="button" class="form-control close"
                 data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </td>
           </tr>
           <tr v-for="(student, index) in students" :key="student.studentInfos.id">
             <td class="align-middle text-center">{{index+1}}</td>
             <td class="align-middle text-center">
               {{student.studentInfos.file_number}}
             </td>
             <td class="align-middle text-center">
               {{student.studentInfos.last_name_fr | capitalize}}
               {{student.studentInfos.first_name_fr | capitalize}}
             </td>
             <td class="align-middle text-center">
               {{student.studentInfos.birthday}}
             </td>
             <td class="align-middle text-center">B</td>
             <td class="align-middle text-center">
               {{student.studentInfos.next_exam}}
             </td>
             <td class="align-middle text-center">
               <input type="checkbox"> </td>
             <td class="align-middle text-center">
               <button @click.prevent="removeStudentAdded(student.studentInfos.id)"
                 type="button" class="form-control close"
                 data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </td>
           </tr>
         </tbody>
       </table>
       <b-button
         :disabled="examList.examListInfos.archived"
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
import Student from './service/student.js';
import ExamList from './service/examList.js';
import Header from './parts/Header';
// import moment from 'moment'

export default {
  components: {
    VueBootstrapDatetimepicker,
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
    found: false,
    loaded: false,

    id: 0,
    data: {},
    examList: {},
    studentLastName: "",
    students: [],
    studentsFound: [],
    updateButton: false,
    errPDF: null,

    options:{
      format:"DD-MM-YYYY",
      useCurrent: false
    },
  }),
  watch:{
    data(){
      this.examList = this.data;
    }
  },
  methods:{
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
      this.updateButton = false;
    },

    // Remove a student added to exam list
    removeStudentAdded: function(studentID){
      for (var i = 0; i < this.students.length; i++) {
        if(this.students[i].studentInfos.id == studentID){
          this.students.splice(i, 1)
        }
      }
    },

    // Search a student in database
    searchStudent:function(){
      //Check if student name length is greather then 2
      if(this.studentLastName.length > 2){
        window.backend.Service.GetStudents(
        this.studentLastName, "", 10, 0).then(
          data=>{
            this.studentsFound = [];
            for (var i = 0; i < data["students"].length; i++) {
              let student = new Student(data["students"][i], null);
              this.studentsFound.push(student);
            }
            this.found = true;
        })
      }else{
        this.studentsFound = [];
        this.found = false;
      }
    },
    removeStudent: function(index){
      this.examList.examListInfos.students_exams.splice(index, 1);
    },
    checkIfAlreadyAdded: function(studentID){
      for (var i = 0; i < this.students.length; i++) {
        if(this.students[i].studentInfos.id == studentID){
          return true
        }
      }
      return false;
    },
    addStudent: function(studentID, index){
      //check if not student is already in exam list
      if(!this.examList.checkIfStudentInExamList(studentID)[0] &&
        !this.checkIfAlreadyAdded(studentID)){
        this.students.push(this.studentsFound[index]);
      }
      this.studentsFound = [];
      this.found = false;
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

    closeAlert:function(){
      this.examListUpdated = false;
      this.errorUpdateExamList = null;
      this.errorArchive = null;
      this.errorDelete = null;
    },

    archiveExamList: function(){
      // switch betxeen true an false to archive an deachive an
      // exam list
      this.examList.archiveAndDearchive();

      // Make the request to backend
      window.backend.Service.ArchiveExamList(
        this.id,
        this.examList.examListInfos.archived
      ).then(
        err=>{this.errorArchive = err}
      );

      // Chenck if there is an error
      if(this.errorArchive != null){
        this.alertVariant = "warning";

        // Check if we wanted to archive or deachive
        if(this.examList.examListInfos.archived){
          this.alertMessage = this.examList.EXAM_LIST_ERR_ARCHIVED_MESSAGE;
        }else{
          this.alertMessage = this.examList.EXAM_LIST_ERR_DEARCHIVED_MESSAGE;
        }

      }else{
        this.alertVariant = "success";
        if(this.examList.examListInfos.archived){
          this.alertMessage = this.examList.EXAM_LIST_ARCHIVED_MESSAGE;
        }else{
          this.alertMessage = this.examList.EXAM_LIST_DEARCHIVED_MESSAGE;
        }
      }
      this.examListUpdated = true;
    },
    updateExamList: function(){
      // Update Exam list to make it exportable to backend
      this.examList.updateExamList();

      // parse student to object
      let outStudentsList = []
      if(this.students.length > 0){
        for (var i = 0; i < this.students.length; i++) {
          this.students[i].outStudent();
          outStudentsList.push(this.students[i].studentInfos);
        }
      }
      // make request to backend
      window.backend.Service.UpdateExamList(
        this.examList.examListInfos.id,
        this.examList.examListInfos.date_exam,
        this.examList.examListInfos.examiner,
        this.examList.examListInfos.students_exams,
        outStudentsList,
      ).then(
        data=>{
          this.examList = new ExamList(data);
          this.students = [];

        },
        err=>{this.errorUpdateExamList = err},
      );

      if(this.errorUpdateExamList != null){
        this.alertVariant = "warning";
        this.alertMessage = this.examList.EXAM_LIST_ERR_UPDATED_MESSAGE;
      }else{
        this.alertVariant = "success";
        this.alertMessage = this.examList.EXAM_LIST_UPDATED_MESSAGE;
      }
      this.examListUpdated = true;
    },
    exportPDF: function(){
      this.examList.updateExamList();
      window.backend.Service.ExportExamListPDF(this.id)
      .then(err=>this.errPDF = err);
    }
  }
}
</script>

<style scoped>
</style>