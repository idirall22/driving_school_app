<template>
  <div id="examLists">
    <div class="mb-5 mt-3">
      <h1>Exam Lists</h1>
    </div>
    <div class="table-responsive">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th class="align-middle text-center">N°</th>
            <th class="align-middle text-center">Examiner</th>
            <th class="align-middle text-center">Exam Date</th>
            <th class="align-middle text-center">Details</th>
          </tr>

        </thead>
        <tbody>
          <tr v-for="exam in examLists" :key="exam.id">
            <td class="align-middle text-center">{{exam.id}}</td>
            <td class="align-middle text-center">{{exam.examiner}}</td>
            <td class="align-middle text-center">{{exam.date_exam | moment}}</td>
            <td class="align-middle text-center"><router-link :to="{ name: 'examListDetails',
                params: { id: exam.id}}"
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
    test: "",
    examlistsCount: 0,
  }),

  methods:{
    getExams:function(){
      window.backend.Service.GetExamLists(10, 0).then(data =>{
        this.examLists = data["examLists"];
        this.examlistsCount = data["count"];
      })
    },
  }
}
</script>
<style scoped>
</style>