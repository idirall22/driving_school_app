<template>
  <div id="addExamList">
    <div class="mb-5 mt-3">
      <h1>Add Exam List</h1>
    </div>

    <div class="row">
      <div class="col-9">
        <input v-model="name" type="text" class="form-control" id="student"
              placeholder="Search student">
      </div>

      <div class="col">
        <button @click="searchStudent" class="btn btn-warning">Find</button>
      </div>

      <div class="col-4 mt-3">
        <datetime placeholder="Exam date"  v-model="date"></datetime>
        {{date}}
      </div>
    </div>

    <div v-if="found">
      <div class="col-4 mt-3 p-0">
        <ul class="list-group p-0 m-0">
          <li v-for="student in studentsFound"  :key="student.ID"
            class="list-group-item d-flex justify-content-between align-items-center">
            <p class="p-0 m-0">{{student.last_name}} {{student.first_name}}</p>
            <button class="btn btn-primary m-0" @click="addStudent(student)" type="button" name="button">Add</button>
          </li>
        </ul>
      </div>
    </div>

    <div class="table-responsive mt-5">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th>N°</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Exam</th>
            <th>Date</th>
            <th>Details</th>
          </tr>

        </thead>
        <tbody>
          <tr v-for="student in students" :key="student.key">
            <td>{{student.file_number}}</td>
            <td>{{student.first_name}}</td>
            <td>{{student.last_name}}</td>
            <td>exam</td>
            <td>01/01/2019</td>
            <td><a class="btn btn-warning" href="#">Infos</a></td>
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
export default {
  name: "addExamList",
  data: () => ({
    name: "",
    examinerName: "No Name",
    date: null,
    found: false,
    students:[],
    studentsFound:[]
  }),
  methods:{
    searchStudent:function(){
      if(this.name.length > 2){
        window.backend.Service.GetStudentsByName(10, 0, this.name).then(data=>{
          this.found = true;
          this.studentsFound = data;
        })
      }else{
        this.studentsFound = [];
        this.found = false;
      }
    },
    addStudent:function(student){
      this.students.push(student);
      this.studentFound = [];
      this.found = false;
      this.name="";
    },
    addExamList: function(){
      window.backend.Service.CreateExamListMap(
        this.date, this.examinerName, this.students)
        
      this.students = [];
      this.studentFound = [];
      this.name="";
    }
  }
}
</script>
<style scoped>
datetime{
  padding: 0px;
  margin: 0px;
}
</style>