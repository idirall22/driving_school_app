<template>
  <div id="">
    <div class="row">
      <div v-if="readyLicence" class="col-6">
        <h5 class="d-flex justify-content-center">Permis Obtenu Durant l'année {{selectedYearLicences}}</h5>
        <line-chart :chart-data="datacollectionLicence"></line-chart>
        <b-form-select
          v-model="selectYearLicenceOption"
          :options="options"
          size="sm"
          class="mt-3"
        ></b-form-select>
      </div>

      <div v-if="readyExams" class="col-6">
        <h5 class="d-flex justify-content-center">Examen gangné durant l'année {{selectedYearExams}}</h5>
        <line-chart :chart-data="datacollectionExams"></line-chart>
        <b-form-select
          v-model="selectYearExamOption"
          :options="options"
          size="sm"
          class="mt-3"
        ></b-form-select>
      </div>
    </div>
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
        datacollectionLicence: null,
        datacollectionExams: null,
        errorgetLicenceWonPerYear: null,
        options: [],
        selectedYearLicences: moment().year(),
        selectedYearExams: moment().year(),
        readyLicence: false,
        readyExams: false
      }
    },
    computed:{
      selectYearLicenceOption:{
        get:function(){return this.selectedYearLicences},
        set:function(value){
          this.selectedYearLicences = value;
          this.getLicenceWonPerYear();
        }
      },
      selectYearExamOption:{
        get:function(){return this.selectedYearExams},
        set:function(value){
          this.selectedYearExams = value;
          this.getExamsWonPerYear();
        }
      }
    },
    mounted () {
      this.getYears();
      this.getLicenceWonPerYear();
      this.getExamsWonPerYear();
    },
    methods: {
      getYears:function(){
          let start = 2010;
          let end = moment().year();
          for (var i = 0; i < (end-start+1); i++) {
            this.options[i] = start + i
          }
      },
      getExamsWonPerYear: function(){
        window.backend.Service.GetExamsResults(this.selectedYearExams)
        .then(
          data=>{
            this.data= analyseDataPerMonth(data);
            this.fillDataExams ()
            this.readyExams = true
          },
          err=>{this.errorgetLicenceWonPerYear= err}
        )
      },
      getLicenceWonPerYear: function(){
        window.backend.Service.GetWinLicencePerYear(this.selectedYearLicences)
        .then(
          data=>{
            this.data= analyseDataPerMonth(data);
            this.fillDataLicence ()
            this.readyLicence = true
          },
          err=>{this.errorgetLicenceWonPerYear= err}
        )
      },
      fillDataLicence () {
        this.datacollectionLicence = {
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
      fillDataExams () {
        this.datacollectionExams = {
          labels: ['Jan', 'Fev', 'Mar', 'Avr', 'Mai', 'Jun', 'Jui', 'Aut', 'Sep', 'Oct', 'Nov', 'Dec'],
          datasets: [
            {
              label: 'Examens Gagné',
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