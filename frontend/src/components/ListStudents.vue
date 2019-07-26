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
        <tr v-for="student in students" :key="student.studentInfos.id">
          <td class="align-middle text-center">{{student.studentInfos.file_number}}</td>
          <td class="align-middle text-center">
            {{student.studentInfos.last_name_fr | capitalize}}
          </td>
          <td class="align-middle text-center">
            {{student.studentInfos.first_name_fr | capitalize}}
          </td>
          <td class="align-middle text-center">
            {{student.studentInfos.registred_date}}
          </td>
          <td class="align-middle text-center">{{student.getExamName(student.studentInfos.next_exam)}}</td>
          <td class="align-middle text-center"><router-link class= "btn btn-danger" :to="{ name: 'studentDetails',
          params: { id: student.studentInfos.id}}">infos</router-link></td>
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
import Student from './service/student.js';
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
      window.backend.Service.GetStudents("", "", 10, 0)
      .then(
        data=>{
          this.data = data;
          for (var i = 0; i < data["students"].length; i++) {
            let s = new Student(data["students"][0], null)
            this.students.push(s);
          }
          this.studentsCount = data["count"];
          this.getNumPages();
          this.loaded = true;
        },
      )
      err=>{this.error = err}
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
  }
}
</script>
<style scoped>
</style>