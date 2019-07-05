<template>
  <div id="studentDetails">
    <div class="mb-5 mt-3">
      <h1>{{student.last_name}} {{student.first_name}}</h1>
    </div>

    <div class="row">
      <div class="contentLeft col-9 ">
      </div>
      <div class="col" v-for="s in student.exams" :key="student.ID">
        <p>{{s.exam}}</p>
      </div>
    </div>


  </div>
</template>
<script>
export default {
  name: "studentDetails",
  created: function(){
    this.getStudent();
  },
  data: () => ({
    student:{},
    id: null
  }),
  methods:{
      getStudent: function(){
        this.id = this.$route.params.id;
        window.backend.Service.GetStudent(this.id).then(data=>{
          this.student = data;
        })
      }
  }
}
</script>
<style scoped>

.contentLeft{
  border-right: 1px solid #333;
}
</style>