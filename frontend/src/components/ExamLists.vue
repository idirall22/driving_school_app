<template>
  <div id="examLists">
    <Header initTitle="Les Listes D'éxamen"></Header>
    <div class="table-responsive">
      <table class="table table-striped table-sm">
        <thead>
          <tr>
            <th class="align-middle text-center">N°</th>
            <th class="align-middle text-center">L'éxaminateur</th>
            <th class="align-middle text-center">Date d'éxamen</th>
            <th class="align-middle text-center">Détails</th>
          </tr>

        </thead>
        <tbody>
          <tr v-for="exam in examLists" :key="exam.id">
            <td class="align-middle text-center">{{exam.id}}</td>
            <td class="align-middle text-center">{{exam.examiner}}</td>
            <td class="align-middle text-center">{{exam.date_exam | moment2}}</td>
            <td class="align-middle text-center">
              <router-link :to="{ name: 'examListDetails',
                params: { id: exam.id}}"
                class="btn btn-outline-primary">Détails</router-link>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div v-if="pages>1" class="d-flex justify-content-center m-5">
      <v-pagination
        v-model="currentPage"
        :page-count="pages"
        :classes="bootstrapPaginationClasses"
        :labels="paginationAnchorTexts">
      </v-pagination>
    </div>
  </div>
</template>
<script>
import vPagination from 'vue-plain-pagination'
import Header from './parts/Header'
export default {
  name: "examLists",
  components: {
    vPagination,
    Header
  },
  created: function(){
    this.getExams();
  },
  data: () => ({
    errorGetData: null,
    url: "",
    examLists:[],
    test: "",
    examlistsCount: 0,
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
        first: '<<',
        prev: '<',
        next: '>',
        last: '>>'
    }
  }),
  watch:{
    currentPage: function(){
      this.getExams();
    }
  },
  methods:{
    getExams:function(){
      window.backend.Service.GetExamLists(
        this.limitPerPage,
        (this.currentPage-1) * this.limitPerPage
      )
      .then(data =>{
        this.examLists = [];
        this.examLists = data["examLists"];
        this.examlistsCount = data["count"];
        this.getNumPages();
      },
      err=> {this.errorGetData= err});
    },
    getNumPages: function(){
      this.pages = Math.floor(this.examlistsCount/this.limitPerPage);
      let remainder = this.examlistsCount%this.limitPerPage;
      if(remainder>0){
        this.pages++;
      }
    },
  }
}
</script>
<style scoped>
</style>