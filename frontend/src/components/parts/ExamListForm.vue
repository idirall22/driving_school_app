<template>
  <div id="examListForm">
    <!-- block one -->
    <div class="row">
      <!-- examiner name -->
      <div class="col-6">
        <label for="examinerName">Examiner</label>
        <input v-model="changeExaminerName" type="text" class="form-control" id="examinerName"
        placeholder="Search">
      </div>
      <!-- date picker -->
      <div class="col-6">
        <label for="datepicker">Exam date</label>
        <vue-bootstrap-datetimepicker v-model="changeExamDate"
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
      <button @click="searchStudent()"
      class="btn btn-warning">Find</button>
    </div>

    <!-- display students founded -->
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

  </div>
</template>

<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
// import moment from 'moment'

export default {
  name: "examListForm",
  props: {
    'students': Array,
    'examinerName': String,
    'examDate': null
  },
  components: {
    VueBootstrapDatetimepicker
  },
  data: () => ({
    studentLastName: "",
    found: false,
    examDateCopy: null,
    examinerNameCopy: "",
    studentsFound: [],
    options:{
      format:"YYYY-MM-DD",
      useCurrent: false
    },
  }),
  computed:{
    changeExamDate:{
      get: function(){return this.examDateCopy},
      set: function(value){this.examDateCopy = value}
    },
    changeExaminerName:{
      get: function(){return this.examinerNameCopy},
      set: function(value){ this.examinerNameCopy = value}
    }
  },
  mounted(){
    if(this.examinerName){
      this.changeExaminerName = this.examinerName
    }

    this.changeExamDate = this.examDate
  },
  methods: {
    //Search a student in databse
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
    addStudent:function(student){

      if(this.students.length == 0){
        // Check if studebts is empty if yes add student directly

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
    resetForm: function(){
      this.studentFound = [];
      this.students = [];
      this.found = false;
      this.studentLastName="";
    }
  },
}
</script>
<style lang="scss" scoped>
</style>