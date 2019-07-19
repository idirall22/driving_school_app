<template>
  <div v-if="loaded" id="examListDetails">
    <b-alert
      v-model="examListUpdated"
      :variant="alertVariant"
      dismissible
      @dismissed="closeAlert()">
      {{alertMessage}}
    </b-alert>


    <div class="row">
      <Header initTitle="Liste d'examen"></Header>
    </div>
    <div class="mb-3 d-flex justify-content-start">
      <button v-if="!examList.archived"
        class="btn btn-danger"
        @click="deleteExamList"
        type="button">Suprimer la liste
      </button>
      <button v-if="!examList.archived"
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

    <b-form @submit="onSubmit">
      <!-- Examiner Name -->
      <b-form-group id="examinerName"
        label="Cherché des étudiants:" label-for="examinerName">
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
              v-model="examDate"
              id="datepicker"
              :config="options">
            </vue-bootstrap-datetimepicker>
          </div>
          <div v-else>
            {{examDate | moment2}}
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
          :readonly="examList.archived">
        </b-form-input>
      </b-form-group>
      <b-button
        :disabled="examList.archived"
        @click="searchStudent()"
        variant="primary">Cherché
      </b-button>
      <!-- end search students -->

      <div v-if="found">
        <div class="col-4 mt-3 p-0">
          <ul class="list-group p-0 m-0">
            <li v-for="student in studentsFound" :key="student.id"
              class="list-group-item d-flex justify-content-between align-items-center">
              <p class="p-0 m-0">
                {{student.last_name | capitalize}}
                {{student.first_name | capitalize}}
              </p>
              <button class="btn btn-primary m-0"
                @click="addStudent(student)"
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
            <tr v-for="(exam, index) in examList.students_exams" :key="exam.id">
              <td class="align-middle text-center">{{index+1}}</td>
              <td class="align-middle text-center">{{exam.student.file_number}}</td>
              <td class="align-middle text-center">{{exam.student.first_name | capitalize}} {{exam.student.last_name | capitalize}}</td>
              <td class="align-middle text-center">{{exam.student.birthday | moment2}}</td>
              <td class="align-middle text-center">B</td>
              <td class="align-middle text-center">{{exam.exam}}</td>
              <td class="align-middle text-center">
                <input v-model="exam.status" type="checkbox">
              </td>

              <td class="align-middle text-center">
                <button @click.prevent="removeStudent(exam)"
                  type="button" class="form-control close"
                  data-dismiss="alert" aria-label="Close">
                 <span aria-hidden="true">&times;</span>
               </button>
             </td>
            </tr>

            <tr v-for="(student, index) in students" :key="student.id">
              <td class="align-middle text-center">{{index+1}}</td>
              <td class="align-middle text-center">{{student.file_number}}</td>
              <td class="align-middle text-center">{{student.first_name | capitalize}} {{student.last_name | capitalize}}</td>
              <td class="align-middle text-center">{{student.birthday | moment2}}</td>
              <td class="align-middle text-center">B</td>
              <td class="align-middle text-center">{{student.next_exam}}</td>
              <td class="align-middle text-center">
                <input type="checkbox"> </td>
              <td class="align-middle text-center">
                <button @click.prevent="removeStudentAdded(student.id)"
                  type="button" class="form-control close"
                  data-dismiss="alert" aria-label="Close">
                 <span aria-hidden="true">&times;</span>
               </button>
             </td>
            </tr>

          </tbody>
        </table>
        <b-button
          :disabled="examList.archived"
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
import Header from './parts/Header';
import moment from 'moment'

export default {
  components: {
    VueBootstrapDatetimepicker,
    Header
  },

  mounted() {
    this.getExamList();
    this.resetMessagesAlert();
  },
  name: "ExamListDetails",
  data: () => ({
    alertVariant:"",
    alertMessage:"",

    errorDelete: null,
    errorArchive: null,
    examListUpdated: false,
    errorUpdateExamList: null,

    errorGetExamList: null,
    id: 0,
    examList: {},
    studentLastName: "",
    found: false,
    loaded: false,
    examDate: null,
    students: [],
    studentsFound: [],
    options:{
      format:"YYYY-MM-DD",
      useCurrent: false
    },
  }),
  methods:{

    // Rset Alert Message
    resetMessagesAlert:function(){
      this.alertVariant = "success";
      this.alertMessage = "La liste d'éxamen a été modifier";
    },

    // get an exam list from database
    getExamList: function(){
      // Get id from url
      this.id = this.$route.params.id;
      if(this.id != 0){
        // Make request to backend
        window.backend.Service.GetExamList(this.id)
        .then(
          data=>{
            this.examList = data;
            this.examDate = moment(this.examList.date_exam).format();
            if(this.examList.students_exams == null){
              this.examList.students_exams = [];
            }
            this.loaded = true;
          },
          err=>{this.errorGetExamList = err;}
        );
      }
    },

    // Remove a student added to exam list
    removeStudentAdded: function(studentID){
      for (var i = 0; i < this.students.length; i++) {
        if(this.students[i].id == studentID){
          this.students.splice(i, 1)
        }
      }
    },

    // Search a student in database
    searchStudent:function(){
      //Check if student name length is greather then 2
      if(this.studentLastName.length > 2){
        window.backend.Service.GetStudents(
          this.studentLastName, "", 10, 0).then(data=>{
          this.found = true;
          this.studentsFound = data["students"];
        })
      }else{
        this.studentsFound = [];
        this.found = false;
      }
    },

    // Add student
    addStudent:function(student){

      if(this.students.length == 0){
        // Check if students is empty if yes add student directly

        this.students.push(student);
        this.studentFound = [];
        this.found = false;
        this.studentLastName="";

      }else{
        // else check if the student was not already added befor
        // add the student

        for (var i = 0; i < this.students.length; i++) {
          if(student.id !=
            this.students[i].id &&
            i+1 == this.students.length){

            this.students.push(student);
            this.studentFound = [];
            this.found = false;
            this.studentLastName="";
          }
        }
      }
    },

    // Remove a student whos already in students exams
    removeStudent: function(exam){
      for (var i = 0; i < this.examList.students_exams.length; i++) {
        if(this.examList.students_exams[i].id == exam.id){
          this.examList.students_exams.splice(i, 1);
        }
      }
    },
    deleteExamList: function(){
      window.backend.Service.DeleteExamList(this.id).then(
        err=>this.errorDelete = err
      );
      if(this.errorDelete == null){
        this.$router.push({name: "examLists"})
      }
    },
    closeAlert:function(){
      this.examListUpdated = false;
      this.errorUpdateExamList = null;
      this.errorArchive = null;
      this.errorDelete = null;
      this.resetMessagesAlert();
    },
    archiveExamList: function(){
      this.examList.archived = !this.examList.archived;
      window.backend.Service.ArchiveExamList(
        this.examList.id,
        this.examList.archived
      ).then(err=>{this.errorArchive = err});
      if(this.errorArchive != null){
        this.alertVariant = "warning";
        this.alertMessage = "La liste d'éxamen n'a pas été archivé";
      }else{
        if(this.examList.archived){
          this.alertVariant = "success";
          this.alertMessage = "La liste d'éxamen a été archivé";
        }else{
          this.alertVariant = "success";
          this.alertMessage = "La liste d'éxamen a été Désarchiver";
        }
      }
      this.examListUpdated = true;
    },
    updateExamList: function(){
      // Update Exam list
      window.backend.Service.UpdateExamList(
        this.examList.id,
        moment(this.examDate).format(),
        this.examList.examiner,
        this.examList.students_exams,
        this.archived,
        this.students,
      ).then(
        data=>{this.data= data},
        err=>{this.errorUpdateExamList = err},
      );

      if(this.errorUpdateExamList != null){
        this.alertVariant = "warning";
        this.alertMessage = "La liste d'éxamen n'a pas été modifier";
      }
      this.examListUpdated = true;
    },
    onSubmit: function(){
      this.updateExamList();
    }
  }
}
</script>
<style scoped>
</style>