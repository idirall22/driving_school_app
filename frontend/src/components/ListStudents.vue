<template>
<div v-if="loaded" id="listStudents">
  <Header initTitle="Students List"></Header>
  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
        <tr>
          <th class="align-middle text-center">N°</th>
          <th class="align-middle text-center">First Name</th>
          <th class="align-middle text-center">Last Name</th>
          <th class="align-middle text-center">Registred</th>
          <th class="align-middle text-center">Exam</th>
          <th class="align-middle text-center">Details</th>
        </tr>

      </thead>
      <tbody>
        <tr v-for="student in students" :key="student.key">
          <td class="align-middle text-center">{{student.file_number}}</td>
          <td class="align-middle text-center">{{student.first_name | capitalize}}</td>
          <td class="align-middle text-center">{{student.last_name | capitalize}}</td>
          <td class="align-middle text-center">{{student.registred_date | moment2}}</td>
          <td class="align-middle text-center">{{student.next_exam}}</td>
          <td class="align-middle text-center"><router-link class= "btn btn-danger" :to="{ name: 'studentDetails',
          params: { id: student.id}}">infos</router-link></td>
        </tr>
      </tbody>
    </table>
    <div class="d-flex justify-content-center m-5">
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
import moment from 'moment';

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
          this.studentsCount = data["count"],
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
  }
}
</script>
<style scoped>
</style>