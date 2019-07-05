<template>
  <div id="examLists">
    <div class="mb-5 mt-3">
      <h1>Exam Lists</h1>
    </div>
    <div class="table-responsive">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th>N°</th>
            <th>Examiner</th>
            <th>Exam Date</th>
            <th>Details</th>
          </tr>

        </thead>
        <tbody>
          <tr v-for="exam in examLists" :key="exam.ID">
            <td>{{exam.ID}}</td>
            <td>{{exam.examiner}}</td>
            <td>{{exam.date_exam}}</td>
            <td><router-link :to="{ name: 'studentDetails'}"
              class="btn btn-danger">infos</router-link></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
export default {
  name: "examLists",
  created: function(){
    this.getExams();
  },
  data: () => ({
    examLists:[],
    test: ""
  }),
  methods:{
    getExams:function(){
      window.backend.Service.GetExamLists(10, 0).then(data =>{
        this.examLists = data;
      })
    }
  }
}
</script>
<style scoped>
</style>