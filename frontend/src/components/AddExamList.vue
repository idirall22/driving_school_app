<template>
  <div id="addExamList">
    <!-- title -->
    <div class="row">
      <div class="col mb-5 mt-3">
        <h1>Add Exam List</h1>
      </div>
      <div v-if="examListCreated"
        class="col mb-5 mt-3 mr-3 alert alert-success" role="alert">
        <button @click.prevent="closeAlert()" type="button" class="close"
            data-dismiss="alert" aria-label="Close">
           <span aria-hidden="true">&times;</span>
         </button>
        <h4 class="alert-heading">Exam List Created!</h4>
        <hr>
        <p>The exam list was added Successfuly</p>
      </div>
      <div v-if="!examListCreated && errorCreateExamList != null" class="col mb-5 mt-3 mr-3 alert alert-warning" role="alert">
        <button @click.prevent="closeAlert()" type="button" class="close"
            data-dismiss="alert" aria-label="Close">
           <span aria-hidden="true">&times;</span>
         </button>
        <h4 class="alert-heading">Could not create a new exam list!</h4>
        <hr>
        <p>The Exam List was not created</p>
      </div>
    </div>

    <!-- block one -->
    <div class="row">
      <!-- examiner name -->
      <div class="col-6">
        <label for="examinerName">Examiner</label>
        <input v-model="examinerName" type="text" class="form-control" id="examinerName"
        placeholder="Search">
      </div>

      <!-- date picker -->
      <div class="col-6">
        <label for="datepicker">Exam date</label>
        <vue-bootstrap-datetimepicker v-model="examDate"
        id="datepicker" :config="options"></vue-bootstrap-datetimepicker>
      </div>

    </div>

    <!-- search srudents field -->
    <div class="row mt-3">

      <div class="col-6">
        <label for="students">Search Students</label>
        <input v-model="studentLastName" type="text" class="form-control" id="students"
        placeholder="Search">
      </div>

    </div>

    <!-- find button -->
    <div class="mt-3">
      <button @click="searchStudent" class="btn btn-warning">Find</button>
    </div>

    <!-- display students founded -->
    <div v-if="found">
      <div class="col-4 mt-3 p-0">
        <ul class="list-group p-0 m-0">
          <li v-for="student in studentsFound" :key="student.id"
            class="list-group-item d-flex justify-content-between align-items-center">
            <p class="p-0 m-0">{{student.last_name}} {{student.first_name}}</p>
            <button class="btn btn-primary m-0" @click="addStudent(student)"
            type="button" name="button">Add</button>
          </li>
        </ul>
      </div>
    </div>

    <!-- array students -->
    <div class="table-responsive mt-5">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th>N°</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Next Exam</th>
            <th>Details</th>
            <th></th>
          </tr>

        </thead>
        <tbody>
          <tr v-for="(student, index) in students" :key="student.key">
            <td>{{student.id}}</td>
            <td>{{student.first_name}}</td>
            <td>{{student.last_name}}</td>
            <td>{{student.next_exam}}</td>
            <td><a class="btn btn-warning" href="#">Infos</a></td>
            <td><a class="btn btn-danger" @click="deleteStudent(index)">X</a></td>
          </tr>
        </tbody>
      </table>
      <button @click.prevent="addExamList()"
              class="btn btn-primary btn-lg btn-block"
              type="submit">Add Exam List</button>
    </div>


  </div>
</template>
<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import moment from 'moment'

export default {
  name: "addExamList",

  components: {
    VueBootstrapDatetimepicker
  },

  data: () => ({
    options:{
      format:"YYYY-MM-DD",
      useCurrent: false
    },
    errorCreateExamList: null,
    examListCreated: false,
    data: null,
    studentLastName: "",
    examinerName: "Examiner Name",
    examDate: null,
    found: false,
    students:[],
    studentsFound:[]
  }),

  methods:{
    searchStudent:function(){
      if(this.studentLastName.length > 2){
        window.backend.Service.GetStudents(
          this.studentLastName, "", 10, 0).then(data=>{
          this.found = true;
          this.studentsFound = data;
        })
      }else{
        this.studentsFound = [];
        this.found = false;
      }
    },
    addStudent:function(student){
      if(this.students.length == 0){
        this.students.push(student);
        this.studentFound = [];
        this.found = false;
        this.studentLastName="";
      }else{
        for (var i = 0; i < this.students.length; i++) {
          if(student.id != this.students[i].id &&
            i+1 == this.students.length){

            this.students.push(student);
            this.studentFound = [];
            this.found = false;
            this.studentLastName="";
          }
        }
      }
    },
    addExamList: function(){
      window.backend.Service.CreateExamList(
        moment(this.examDate).format(), this.examinerName, this.students).then(
          data=>{this.data = data},
          err=>{this.errorCreateExamList= err}
        )
      if(this.errorCreateStudent != null){
        this.examListCreated = false;
      }else{
        this.examListCreated = true;
        this.students = [];
        this.studentFound = [];
        this.studentLastName="";
      }
    },
    closeAlert:function(){
      this.examListCreated = false;
      this.errorCreateStudent = null;
    },
    deleteStudent: function(studentID){
      this.students.splice(studentID, 1)
    }
  }
}
</script>
<style scoped>
</style>
