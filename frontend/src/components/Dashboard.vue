<template>
  <div id="">
    <h5 class="d-flex justify-content-center">{{selectedTable}} Durant l'année {{selectedYear}}</h5>
    <b-form-select
      v-model="selectYearOptions"
      :options="optionsYears"
      class="mt-3 mb-3"
    ></b-form-select>

    <b-form-select
      v-model="selectTableOptions"
      :options="optionsData"
      class="mt-3 mb-3"
    ></b-form-select>

    <line-chart :chart-data="dataCollection"></line-chart>
  </div>
</template>

<script>
  import LineChart from './scripts/LineChart.js'
  import {analyseDataPerMonth} from './service/analytic.js'
  import moment from 'moment';
  export default {
    components: {
      LineChart
    },
    data () {
      return {
        dataCollection: null,
        errorGetAnalytics: null,
        optionsYears: [],
        optionsData: [
          "Inscription",
          "Examen gagné",
          "Permis Obtenu",
        ],
        selectedYear: moment().year(),
        selectedTable: "Permis Obtenu",

        table: "students",
        column: "win_licence_date",

        ready: false,
      }
    },
    computed:{
      selectYearOptions:{
        get:function(){return this.selectedYear},
        set:function(value){
          this.selectedYear = value;
          this.getAnalytics();
        }
      },
      selectTableOptions:{
        get:function(){
          return this.selectedTable
        },
        set:function(value){
          this.selectedTable = value;
          switch (this.selectedTable) {
            case "Inscription":
              this.table = "students";
              this.column = "registred_date";
              break;
            case "Examen Gagné":
              this.table = "exams";
              this.column = "date_exam";
              break;
            case "Permis Obtenu":
              this.table = "students";
              this.column = "win_licence_date";
              break;
            default:
              this.table = "students";
              this.column = "win_licence_date";
          }
          this.getAnalytics();
        },
      }
    },
    created() {
      this.getAnalytics();
      this.getYears();
    },
    methods: {
      getYears:function(){
          let start = 2010;
          let end = moment().year();
          for (var i = 0; i < (end-start+1); i++) {
            this.optionsYears[i] = start + i
          }
      },
      getAnalytics: function(){
        window.backend.Service.Analytics(
          this.table, this.column, this.selectedYear
        )
        .then(
          data=>{
            this.data= analyseDataPerMonth(data);
            this.fillData();
            this.ready = true;
          },
          err=>{this.errorGetAnalytics= err}
        )
      },
      fillData () {
        this.dataCollection = {
          labels: ['Jan', 'Fev', 'Mar', 'Avr', 'Mai', 'Jun', 'Jui', 'Aut', 'Sep', 'Oct', 'Nov', 'Dec'],
          datasets: [
            {
              label: 'Permis Obtenu',
              backgroundColor: 'rgba(66,139,202,128)',
              data: this.data
            }
          ]
        }
      },
    }
  }
</script>
<style scoped>
</style>