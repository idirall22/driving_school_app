<template>
  <div id="addExamList">

    <div class="mb-5 mt-3">
      <h1>Add Exam List</h1>
    </div>

    <div class="row">
        <div class="col-9">
        <input type="text" class="form-control" id="student"
              placeholder="Search student">
      </div>
      <div class="col">
        <button class="btn btn-danger">Add</button>
      </div>
      <div class="col-4 mt-3">
        <input type="text" class="form-control" id="student" placeholder="Exam date">
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
    </div>


  </div>
</template>
<script>
export default {
  name: "addExamList",
  data: () => ({

  })
}
</script>
<style scoped>
</style>