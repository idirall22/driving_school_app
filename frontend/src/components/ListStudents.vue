<template>
<div id="listStudents">
  <div class="mb-5 mt-3">
    <h1>Students List</h1>
  </div>
  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
        <tr>
          <th>N°</th>
          <th>First Name</th>
          <th>Last Name</th>
          <th>Exam</th>
          <th>Date</th>
        </tr>

      </thead>
      <tbody>
        <tr v-for="student in students" :key="student.key">
          <td>{{student.file_number}}</td>
          <td>{{student.first_name}}</td>
          <td>{{student.last_name}}</td>
          <td>exam</td>
          <td>01/01/2019</td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
</template>

<script>
export default {
  name: "listStudents",
  created:function() {
    this.getStudentsList();
  },
  data: () => ({
    students: [],
  }),
  methods:{
    getStudentsList: function(){
      if (this.students.length == 0){
        window.backend.Service.GetStudents(10, 0).then(data=>{
          this.students = data;
        })
      }
    }
  }
}
</script>
<style scoped>
</style>