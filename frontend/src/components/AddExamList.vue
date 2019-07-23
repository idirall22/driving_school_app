<template>
  <div id="addExamList">

    <div class="d-flex justify-content-end">
      <b-alert
        v-model="examListCreated"
        :variant="alertVariant"
        dismissible
        @dismissed="closeAlert()">
        {{alertMessage}}
      </b-alert>
    </div>
    <!-- title -->
    <Header initTitle="Ajouter une Liste D'Examen"></Header>

    <b-form @submit="onSubmit">

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

      <!-- Search students -->
      <b-form-group id="searchStudents"
        label="Cherché des étudiants:" label-for="searchStudents">
        <b-form-input
          id="searchStudents"
          v-model="studentLastName"
          type="text"
          placeholder="Nom de l'étudiant">
        </b-form-input>
      </b-form-group>
      <b-button @click="searchStudent()" variant="primary">Cherché</b-button>
      <!-- end search students -->

      <div v-if="found">
        <div class="col-4 mt-3 p-0">
          <ul class="list-group p-0 m-0">
            <li v-for="student in studentsFound" :key="student.studentInfos.id"
              class="list-group-item d-flex justify-content-between align-items-center">
              <p class="p-0 m-0">
                {{student.studentInfos.last_name_fr | capitalize}}
                {{student.studentInfos.first_name_fr | capitalize}}
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

      <!-- array students -->
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
            <tr v-for="(student, index) in students" :key="student.key">
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
                {{student.getStudentNextExamName()}}</td>
              <td>
                <button @click.prevent="deleteStudent(index)"
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
import Student from './service/student.js';
import moment from 'moment'

export default {
  name: "addExamList",
  components: {
    VueBootstrapDatetimepicker,
    Header
  },
  data: () => ({
    alertVariant:"success",
    alertMessage:"La liste d'examen a été crée.",
    errorCreateExamList: null,
    examListCreated: false,
    examinerName: "",
    studentLastName: "",
    examDate: null,
    found: false,
    studentsFound:[],
    students: [],
    options:{
      format:"DD-MM-YYYY",
      useCurrent: false
    },
  }),
  methods:{

    // Delete student from exam list
    deleteStudent: function(index){
      this.students.splice(index, 1)
    },
    searchStudent:function(){
      //Check if student name length is greather then 2
      if(this.studentLastName.length > 2){
        window.backend.Service.GetStudents(
          this.studentLastName, "", 10, 0)
          .then(
            data=>{
              // Convert data to student object
              for (var i = 0; i < data["students"].length; i++) {
                let student = new Student(data["students"][i], null);
                this.studentsFound.push(student)
              }
              this.found = true;
        })
      }else{
        this.studentsFound = [];
        this.found = false;
      }
    },
    addStudent:function(student){

      if(this.students.length == 0){
        // Check if students is empty if yes add student directly
        this.students.push(student);
      }else{
        // else check if the student was not already added befor
        // add the student
        for (var i = 0; i < this.students.length; i++) {
          if(student.id != this.students[i].id && i+1 == this.students.length){
            this.students.push(student);
          }
        }
      }
      this.studentFound = [];
      this.found = false;
      this.studentLastName="";
    },
    closeAlert: function(){
      this.errorCreateExamList = null;
      this.examListCreated =  false;
    },
    onSubmit: function(){
      let outStudents = [];
      if(this.students.length > 0){
        for (var i = 0; i < this.students.length; i++) {
          this.students[i].outStudent();
          outStudents.push(this.students[i].studentInfos);
        }
      }
      window.backend.Service.CreateExamList(
        moment(this.examDate, "DD-MM-YYYY").format(),
          this.examinerName, outStudents)
        .then(
          data=>{this.data = data},
          err=>{this.errorCreateExamList= err}
        )
      if(this.errorCreateExamList != null){
        this.alertVariant = "warning";
        this.alertMessage ="La liste d'éxamen n'a pas été crée";
      }else{
        this.alertVariant = "success";
        this.alertMessage = "La liste d'éxamen a été crée";
      }
      this.examListCreated = true;
    }
  },
}
</script>
<style scoped>
</style>

