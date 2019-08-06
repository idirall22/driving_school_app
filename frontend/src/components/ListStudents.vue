<template>
<div v-if="loaded" id="listStudents">
  <Header initTitle="Liste d'Étudiants"></Header>

  <!-- student search form -->
  <b-form @submit="searchStudent">
    <b-form-group label="Récherché Un Étudiant:">
      <b-radio v-model="radioSelectedSearchMethod" name="some-radios" value="lastName">Nom de Famille</b-radio>
      <b-radio v-model="radioSelectedSearchMethod" name="some-radios" value="phoneNumber">Numero de Téléphone</b-radio>
    </b-form-group>

    <!-- search -->
    <b-form-group
      id="search-group"
      label="Recherché:"
      label-for="search">

      <b-form-input
        v-model="searchValue"
        id="search"
        type="text"
        placeholder="Recherche"
      ></b-form-input>
    </b-form-group>
    <!-- end search -->

    <b-button
      type="submit"
      variant="primary"
      class="mb-3">
      Recherché
    </b-button>
  </b-form>

  <div class="table-responsive">
    <table class="table table-striped table-sm">
      <thead>
        <tr>
          <th class="align-middle text-center">N°</th>
          <th class="align-middle text-center">Nom</th>
          <th class="align-middle text-center">Prénom</th>
          <th class="align-middle text-center">Date d'Inscription</th>
          <th class="align-middle text-center">Détails</th>
        </tr>

      </thead>
      <tbody v-if="!searchMode">
        <tr  v-for="student in students" :key="student.id">
          <td class="align-middle text-center">{{student.file_number}}</td>
          <td class="align-middle text-center">
            {{student.last_name_fr | capitalize}}
          </td>
          <td class="align-middle text-center">
            {{student.first_name_fr | capitalize}}
          </td>
          <td class="align-middle text-center">
            {{student.registred_date}}
          </td>
          <td class="align-middle text-center"><router-link class= "btn btn-danger" :to="{ name: 'studentDetails',
          params: { id: student.id}}">infos</router-link></td>
        </tr>
      </tbody>

      <tbody v-if="searchMode && student !== null">
        <tr>
          <td class="align-middle text-center">{{student.file_number}}</td>
          <td class="align-middle text-center">
            {{student.last_name_fr | capitalize}}
          </td>
          <td class="align-middle text-center">
            {{student.first_name_fr | capitalize}}
          </td>
          <td class="align-middle text-center">
            {{student.registred_date}}
          </td>
          <td class="align-middle text-center"><router-link class= "btn btn-danger" :to="{ name: 'studentDetails',
          params: { id: student.id}}">infos</router-link></td>
        </tr>
      </tbody>
    </table>
    <div v-if="searchMode && student === null"  class="d-flex justify-content-center">
      <p>Il n'y pas d'étudiant</p>
    </div>
    <div v-if="pages > 1" class="d-flex justify-content-center m-5 overflow-auto">
      <b-pagination
        v-model="changePage"
        :total-rows="studentsCount"
        :per-page="limitPerPage"
      ></b-pagination>
    </div>
  </div>
</div>
</template>

<script>
import Student from './service/student.js';
import Header from './parts/Header'
// import moment from 'moment';

export default {
  name: "listStudents",
  components: {
    Header
  },
  mounted:function() {
    this.getStudentsList(this.limitPerPage, 0);
  },
  computed:{
    changePage:{
      get: function(){return this.currentPage},
      set: function(val){
        this.currentPage = val;
        this.getStudentsList(this.limitPerPage, (val - 1) * this.limitPerPage)
      },
    }
  },
  data: () => ({
    radioSelectedSearchMethod: "lastName",
    searchValue: "",
    student: null,
    searchMode: false,
    errorSearch: null,

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
  watch:{
    searchValue: function(val){
      if(val.length < 2 && this.searchMode){
        this.searchMode = false;
        this.student = null;
      }
    }
  },
  methods:{
    getStudentsList: function(limit, offset){
      this.students = [];
      window.backend.Service.GetStudents("", "", limit, offset)
      .then(
        data=>{
          this.data = data;
          for (var i = 0; i < data["students"].length; i++) {
            let s = new Student(data["students"][i], null)
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
        let remainder = this.studentsCount%this.limitPerPage;
        if(remainder>0){
          this.pages++;
        }
      }
    },
    searchStudent: function(evt){
      evt.preventDefault()
      let ln = "";
      let fn = "";

      if(this.radioSelectedSearchMethod == "lastName"){
        if(this.searchValue.length > 2){
          ln = this.searchValue;
          fn = "";
        }
      }else if(this.radioSelectedSearchMethod == "phoneNumber"){
        if(this.searchValue.length == 10){
          ln = "";
          fn = this.searchValue;
        }
      }
      if(ln != "" || fn != ""){
        window.backend.Service.GetStudent(0, ln, fn)
        .then(
          data=>{
            this.student = new Student(data["student"], null);
            this.searchMode = true;
          },
          err=>{
            this.errorSearch = err;
            this.searchMode = true;
          }
        );
      }
    }
  }
}
</script>
<style scoped>
</style>