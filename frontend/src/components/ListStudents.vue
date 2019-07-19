<template>
<div v-if="loaded" id="listStudents">
  <Header initTitle="Liste d'étudiants"></Header>
  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
        <tr>
          <th class="align-middle text-center">N°</th>
          <th class="align-middle text-center">Nom</th>
          <th class="align-middle text-center">Prénom</th>
          <th class="align-middle text-center">Date d'Inscription</th>
          <th class="align-middle text-center">Examen</th>
          <th class="align-middle text-center">Détails</th>
        </tr>

      </thead>
      <tbody>
        <tr v-for="student in students" :key="student.key">
          <td class="align-middle text-center">{{student.file_number}}</td>
          <td class="align-middle text-center">{{student.last_name | capitalize}}</td>
          <td class="align-middle text-center">{{student.first_name | capitalize}}</td>
          <td class="align-middle text-center">{{student.registred_date | moment2}}</td>
          <td class="align-middle text-center">{{getStudentNextExamName(student)}}</td>
          <td class="align-middle text-center"><router-link class= "btn btn-danger" :to="{ name: 'studentDetails',
          params: { id: student.id}}">infos</router-link></td>
        </tr>
      </tbody>
    </table>
    <div v-if="pages>1" class="d-flex justify-content-center m-5">
      <v-pagination
        v-model="currentPage"
        :page-count="pages"
        :classes="bootstrapPaginationClasses"
        :labels="paginationAnchorTexts">
      </v-pagination>
    </div>
  </div>
</div>
</template>

<script>
import vPagination from 'vue-plain-pagination'
import Header from './parts/Header'
// import moment from 'moment';

export default {
  name: "listStudents",
  components: {
    vPagination,
    Header
  },
  mounted:function() {
    this.getStudentsList();
  },
  data: () => ({
    students: [],
    loaded: false,
    error: null,
    studentsCount: 0,

    limitPerPage: 10,
    pages: 1,
    currentPage: 1,
    data:null,
    bootstrapPaginationClasses: {
      ul: 'pagination',
      li: 'page-item',
      liActive: 'active',
      liDisable: 'disabled',
      button: 'page-link'
    },
    paginationAnchorTexts: {
        first: 'First',
        prev: 'Previous',
        next: 'Next',
        last: 'Last'
    }
  }),
  methods:{
    getStudentsList: function(){
      // if (this.students.length == 0){
      window.backend.Service.GetStudents("", "", 10, 0)
      .then(
        data=>{
          this.students = data["students"];
          this.studentsCount = data["count"];

          for (var i = 0; i < this.students.length; i++) {
            let fn = JSON.parse(this.students[i].first_name);
            this.students[i].first_name = fn.fr;
            let ln = JSON.parse(this.students[i].last_name);
            this.students[i].last_name = ln.fr;
          }
          this.loaded = true;
          this.getNumPages();
        },
        err=>{this.error = err}
      )
    },
    getNumPages: function(){
      if(this.studentsCount > this.limitPerPage){
        this.pages = Math.floor(this.studentsCount/this.limitPerPage);
        let remainder = this.examlistsCount%this.limitPerPage;
        if(remainder>0){
          this.pages++;
        }
      }
    },
    getStudentNextExamName: function(student){
      switch (student.next_exam) {
        case "Highway code":
          return "Code";
        case "Niche":
          return "Créneau";
        case "Circuit":
          return "Circuit";
      }
    },
  }
}
</script>
<style scoped>
</style>